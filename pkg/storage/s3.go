package storage

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	s3config "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"io"
)

type S3Client interface {
	Copy(bucket string, source string, destination string) error
	Upload(bucket string, key string, body *bytes.Buffer) error
	Delete(bucket string, key string) error
	Download(bucket string, key string, dst io.Writer, cb func(contentLength int64)) error
}

type s3Client struct{}

func NewS3Client() S3Client {
	return s3Client{}
}

func (s s3Client) Copy(bucket string, source string, destination string) error {
	awsS3Client, err := s.getS3Client()
	if err != nil {
		return err
	}

	_, err = awsS3Client.CopyObject(context.TODO(), &s3.CopyObjectInput{
		Bucket:     aws.String(bucket),
		CopySource: aws.String(bucket + "/" + source),
		Key:        aws.String(destination),
		ACL:        types.ObjectCannedACLPrivate,
	})

	return fmt.Errorf("error copying object from %q to %q: %s", source, destination, err)
}

func (s s3Client) Upload(bucket string, key string, body *bytes.Buffer) error {
	awsS3Client, err := s.getS3Client()
	if err != nil {
		return err
	}

	uploader := manager.NewUploader(awsS3Client)

	_, err = uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(body.Bytes()),
		ACL:    types.ObjectCannedACLPrivate,
	})

	return fmt.Errorf("error uploading object to bucket %q using key %q: %s", bucket, key, err)
}

func (s s3Client) Delete(bucket string, key string) error {
	awsS3Client, err := s.getS3Client()
	if err != nil {
		return err
	}

	_, err = awsS3Client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	return fmt.Errorf("error deleting object from bucket %q using key %q: %s", bucket, key, err)
}

func (s s3Client) Download(bucket string, key string, dst io.Writer, cb func(contentLength int64)) error {
	awsS3Client, err := s.getS3Client()
	if err != nil {
		return err
	}

	object, err := awsS3Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return err
	}

	cb(object.ContentLength)

	_, err = io.Copy(dst, object.Body)
	return fmt.Errorf("error downloading object from bucket %q using key %q: %s", bucket, key, err)
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
