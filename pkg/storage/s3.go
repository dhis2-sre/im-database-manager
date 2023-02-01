package storage

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func NewS3Client(client AWSS3Client, uploader AWSS3Uploader) (*S3Client, error) {
	return &S3Client{client, uploader}, nil
}

type S3Client struct {
	client   AWSS3Client
	uploader AWSS3Uploader
}

type AWSS3Client interface {
	CopyObject(ctx context.Context, params *s3.CopyObjectInput, optFns ...func(*s3.Options)) (*s3.CopyObjectOutput, error)
	DeleteObject(ctx context.Context, params *s3.DeleteObjectInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectOutput, error)
	GetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error)
}

type AWSS3Uploader interface {
	Upload(ctx context.Context, input *s3.PutObjectInput, opts ...func(*manager.Uploader)) (*manager.UploadOutput, error)
}

func (s S3Client) Copy(bucket string, source string, destination string) error {
	_, err := s.client.CopyObject(context.TODO(), &s3.CopyObjectInput{
		Bucket:     aws.String(bucket),
		CopySource: aws.String(bucket + "/" + source),
		Key:        aws.String(destination),
		ACL:        types.ObjectCannedACLPrivate,
	})

	if err != nil {
		return fmt.Errorf("error copying object from %q to %q: %s", source, destination, err)
	}
	return nil
}

func (s S3Client) Upload(bucket string, key string, body *bytes.Buffer) error {
	_, err := s.uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(body.Bytes()),
		ACL:    types.ObjectCannedACLPrivate,
	})

	if err != nil {
		return fmt.Errorf("error uploading object to bucket %q using key %q: %s", bucket, key, err)
	}
	return nil
}

func (s S3Client) Delete(bucket string, key string) error {
	_, err := s.client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		return fmt.Errorf("error deleting object from bucket %q using key %q: %s", bucket, key, err)
	}
	return nil
}

func (s S3Client) Download(bucket string, key string, dst io.Writer, cb func(contentLength int64)) error {
	object, err := s.client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return fmt.Errorf("error downloading object from bucket %q using key %q: %s", bucket, key, err)
	}

	cb(object.ContentLength)

	_, err = io.Copy(dst, object.Body)

	return err
}
