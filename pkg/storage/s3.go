package storage

import (
	"bytes"
	"context"
	s3config "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Client interface {
	Upload(bucket string, key string, body *bytes.Buffer) error
}

func ProvideS3Client() S3Client {
	return s3Client{}
}

type s3Client struct {
}

func (s s3Client) Upload(bucket string, key string, body *bytes.Buffer) error {
	awsS3Client, err := s.getS3Client()
	if err != nil {
		return err
	}

	uploader := manager.NewUploader(awsS3Client)

	b := bytes.NewReader(body.Bytes())

	_, err = uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &key,
		Body:   b,
	})

	return err
}

func (s s3Client) getS3Client() (*s3.Client, error) {
	// TODO: Region? Why isn't our bucket in eu-north-1?
	cfg, err := s3config.LoadDefaultConfig(context.TODO(), s3config.WithRegion("eu-west-1"))
	if err != nil {
		return nil, err
	}

	awsS3Client := s3.NewFromConfig(cfg)
	return awsS3Client, err
}
