package s3

import "github.com/aws/aws-sdk-go/aws/session"

type Usecase interface {
	AddFileToS3(s *session.Session, fileDir string) error
	AddFileToS3Stream(s *session.Session, data string) error
	ReadFile(data string) (string, error)
}
