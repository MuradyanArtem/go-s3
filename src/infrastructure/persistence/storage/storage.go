package storage

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// S3StorageConfig is describe configuration of s3 storage
type S3StorageConfig struct {
	Bucket string

	UploadFilePartSizeMB   int
	DownloadFilePartSizeMB int

	UploadLinkLifetime   time.Duration
	DownloadLinkLifetime time.Duration

	DownloadFileTimeout time.Duration
}

// S3Storage is describe S3 storage instance
type S3Storage struct {
	config     S3StorageConfig
	client     s3iface.S3API
	uploader   *s3manager.Uploader
	downloader *s3manager.Downloader
}

// NewS3Storage creates new S3 storage instance
func NewS3Storage(client s3iface.S3API, config S3StorageConfig) *S3Storage {
	return &S3Storage{
		config: config,
		client: client,
		uploader: s3manager.NewUploaderWithClient(client, func(u *s3manager.Uploader) {
			u.PartSize = int64(config.UploadFilePartSizeMB * 1024 * 1024)
		}),
		downloader: s3manager.NewDownloaderWithClient(client, func(d *s3manager.Downloader) {
			d.PartSize = int64(config.DownloadFilePartSizeMB * 1024 * 1024)
		}),
	}
}

// GetUploadURL creates and returns uri for upload object to S3
func (s *S3Storage) GetUploadURL(path string, size int64) (string, error) {
	req, _ := s.client.PutObjectRequest(&s3.PutObjectInput{
		Bucket:        aws.String(s.config.Bucket),
		Key:           aws.String(path),
		ContentLength: &size,
	})
	if req.Error != nil {
		return "", req.Error
	}
	return req.Presign(s.config.UploadLinkLifetime)
}

// GetDownloadURL creates and returns uri to download object from S3
func (s *S3Storage) GetDownloadURL(path string) (string, error) {
	req, _ := s.client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(s.config.Bucket),
		Key:    aws.String(path),
	})
	if req.Error != nil {
		return "", req.Error
	}
	return req.Presign(s.config.DownloadLinkLifetime)
}

// GetFileSize returns file size by path
func (s *S3Storage) GetFileSize(path string) (int, error) {
	output, err := s.client.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(s.config.Bucket),
		Key:    aws.String(path),
	})
	if err != nil {
		return 0, err
	}
	if output.ContentLength == nil {
		return 0, nil
	}

	return int(*output.ContentLength), nil
}

// DeleteFile removes file by path
func (s *S3Storage) DeleteFile(path string) error {
	_, err := s.client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(s.config.Bucket),
		Key:    aws.String(path),
	})
	return err
}

// GetRootCursor returns iterator to navigate through object in S3
func (s *S3Storage) GetRootCursor() (*Cursor, error) {
	output, err := s.client.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(s.config.Bucket),
	})
	if err != nil {
		return nil, err
	}

	return NewCursor(s.client, output), nil
}
