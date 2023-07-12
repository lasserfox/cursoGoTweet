package routers

import (
	"bytes"
	"context"
	"cursoGoTweet/bd"
	"cursoGoTweet/models"
	"encoding/base64"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"io"
	"mime"
	"mime/multipart"
	"strings"
)

type readSeeker struct {
	io.Reader
}

func (rs *readSeeker) Seek(offset int64, whence int) (int64, error) {
	return 0, nil
}

func UploadImage(ctx context.Context, uploadType string, request events.APIGatewayProxyRequest, claim models.Claim) models.RespApi {
	var r models.RespApi
	r.Status = 400
	IDUsuario := claim.ID.Hex()
	fmt.Println("Entrando en UploadImage")
	var filename string
	var usuario models.Usuario

	bucket := aws.String(ctx.Value(models.Key("bucketName")).(string))
	switch uploadType {
	case "A":
		filename = "avatars/" + IDUsuario + ".jpg"
		usuario.Avatar = filename
	case "B":
		filename = "banners/" + IDUsuario + ".jpg"
		usuario.Banner = filename
	}

	mediaType, params, err := mime.ParseMediaType(request.Headers["Content-Type"])
	if err != nil {
		r.Status = 500
		r.Message = err.Error()
		return r
	}
	if strings.HasPrefix(mediaType, "multipart/") {
		body, err := base64.StdEncoding.DecodeString(request.Body)
		if err != nil {
			r.Status = 500
			r.Message = err.Error()
			return r
		}
		mr := multipart.NewReader(bytes.NewReader(body), params["boundary"])
		p, err := mr.NextPart()
		if err != nil && err != io.EOF {
			r.Status = 500
			r.Message = err.Error()
			return r
		}

		if err != io.EOF {
			if p.FileName() != "" {
				buf := bytes.NewBuffer(nil)
				if _, err := io.Copy(buf, p); err != nil {
					r.Status = 500
					r.Message = err.Error()
					return r
				}
				sess, err := session.NewSession(&aws.Config{
					Region: aws.String("us-east-1")})

				if err != nil {
					r.Status = 500
					r.Message = err.Error()
					return r
				}
				uploader := s3manager.NewUploader(sess)
				_, err = uploader.Upload(&s3manager.UploadInput{
					Bucket: bucket,
					Key:    aws.String(filename),
					Body:   &readSeeker{buf}})

				if err != nil {
					r.Status = 500
					r.Message = err.Error()
					return r
				}
			}
		}

		var status bool
		status, err = bd.ModificoRegistro(usuario, IDUsuario)
		if err != nil || !status {
			r.Status = 400
			r.Message = "Error al modificar el registro de usuario"
			return r
		}

	} else {
		r.Status = 400
		r.Message = "Debe enviar un Content-Type tipò 'multipart/'"
		return r
	}

	r.Status = 200
	r.Message = "Image upload OK"
	return r
}
