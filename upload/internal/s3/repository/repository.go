package repository

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"gitlab.com/Spide_IT/spide_it/upload/config"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type Repository struct {
	config *config.Config
	s3 *session.Session
}

func NewRepository(config *config.Config, s3 *session.Session)*Repository{
	return &Repository{
		config:config,
		s3:s3,
	}
}

func (u *Repository) AddFileToS3(s *session.Session, fileDir string) error {
	file, err := os.Open(fileDir)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)

	_ , err =s3.New(s).PutObject(&s3.PutObjectInput{
		ACL:                       aws.String(u.config.S3.S3ACL),
		Body:                      bytes.NewReader(buffer),
		Bucket:                    aws.String(u.config.S3.S3Bucket),
		ContentDisposition:        aws.String("attachment"),
		ContentLength:              aws.Int64(size),
		ContentType:               aws.String(http.DetectContentType(buffer)),
		Key:                       aws.String(fileDir),
		ServerSideEncryption:      aws.String("AES256"),

	})
	return err
}

func (u *Repository) AddFileToS3Stream(s *session.Session, data string) error {
	reader := strings.NewReader("Geeks")

	buffer := make([]byte, 4)

	n, err := io.ReadFull(reader, buffer)
	if err != nil {
		return err
	}
	fileDir:=data[0:4]+time.Now().String()

	_ , err =s3.New(s).PutObject(&s3.PutObjectInput{
		ACL:                       aws.String(u.config.S3.S3ACL),
		Body:                      bytes.NewReader(buffer),
		Bucket:                    aws.String(u.config.S3.S3Bucket),
		ContentDisposition:        aws.String("attachment"),
		ContentLength:              aws.Int64(int64(n)),
		ContentType:               aws.String(http.DetectContentType(buffer)),
		Key:                       aws.String(fileDir),
		ServerSideEncryption:      aws.String("AES256"),

	})
	return err

}

func (u *Repository) ReadFile(data string) (string, error) {
	results, err := s3.New(u.s3).GetObject(&s3.GetObjectInput{
		Bucket: aws.String(u.config.S3.S3Bucket),
		Key:    aws.String(data),
	})
	if err != nil {
		return "", err
	}
	defer results.Body.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, results.Body); err != nil {
		return "", err
	}
	return string(buf.Bytes()), nil
}
