package usecase

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"gitlab.com/Spide_IT/spide_it/upload/internal/s3"
)


type Usecase struct {
	repo s3.Repository
}

func NewUsecase(repo s3.Repository) *Usecase{
	return &Usecase{repo:repo}
}

func (u *Usecase) AddFileToS3(s *session.Session, fileDir string) error{
	return u.repo.AddFileToS3(s,fileDir)
}

func (u *Usecase) AddFileToS3Stream(s *session.Session, data string) error{
	return u.repo.AddFileToS3Stream(s,data)
}

func (u *Usecase) ReadFile(data string) (string, error){
	return u.repo.ReadFile(data)
}