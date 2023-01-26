package storage

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type S3Client interface {
	Copy(bucket string, source string, destination string) error
	Upload(bucket string, key string, body *bytes.Buffer) error
	Delete(bucket string, key string) error
	Download(bucket string, key string, dst io.Writer, cb func(contentLength int64)) error
}

func NewS3Client(client AWSS3Client) (S3Client, error) {
	return s3Client{client}, nil
}

type AWSS3Client interface {
	CopyObject(ctx context.Context, params *s3.CopyObjectInput, optFns ...func(*s3.Options)) (*s3.CopyObjectOutput, error)
	DeleteObject(ctx context.Context, params *s3.DeleteObjectInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectOutput, error)
	GetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error)

	PutObject(context.Context, *s3.PutObjectInput, ...func(*s3.Options)) (*s3.PutObjectOutput, error)
	UploadPart(context.Context, *s3.UploadPartInput, ...func(*s3.Options)) (*s3.UploadPartOutput, error)
	CreateMultipartUpload(context.Context, *s3.CreateMultipartUploadInput, ...func(*s3.Options)) (*s3.CreateMultipartUploadOutput, error)
	CompleteMultipartUpload(context.Context, *s3.CompleteMultipartUploadInput, ...func(*s3.Options)) (*s3.CompleteMultipartUploadOutput, error)
	AbortMultipartUpload(context.Context, *s3.AbortMultipartUploadInput, ...func(*s3.Options)) (*s3.AbortMultipartUploadOutput, error)
}

type s3Client struct {
	client AWSS3Client
}

func (s s3Client) Copy(bucket string, source string, destination string) error {
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

func (s s3Client) Upload(bucket string, key string, body *bytes.Buffer) error {
	uploader := manager.NewUploader(s.client)

	_, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
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

func (s s3Client) Delete(bucket string, key string) error {
	_, err := s.client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		return fmt.Errorf("error deleting object from bucket %q using key %q: %s", bucket, key, err)
	}
	return nil
}

func (s s3Client) Download(bucket string, key string, dst io.Writer, cb func(contentLength int64)) error {
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
