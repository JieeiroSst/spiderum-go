package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"log"
	"sync"
)

type S3 struct {
	s3 *session.Session
}

var (
	instance *S3
	once     sync.Once
)

func GetS3ConnInstance(region string) *S3{
	once.Do(func() {
		session,err:=session.NewSession(&aws.Config{Region:aws.String(region)})
		if err!=nil{
			log.Println(err)
		}
		instance = &S3{s3:session}
	})
	return instance
}

func NewS3(region string) *session.Session {
	return GetS3ConnInstance(region).s3
}
