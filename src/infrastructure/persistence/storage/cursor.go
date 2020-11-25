package storage

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

// Cursor is iterator for s3 files
type Cursor struct {
	client s3iface.S3API
	index  int
	output *s3.ListObjectsV2Output
	err    error
}

// NewCursor create cursor instance
func NewCursor(client s3iface.S3API, output *s3.ListObjectsV2Output) *Cursor {
	return &Cursor{
		client: client,
		output: output,
		index:  -1,
	}
}

// Next update cursor position to the next object
func (c *Cursor) Next() bool {
	if c.output == nil || c.err != nil {
		return false
	}

	c.index++
	if c.index < len(c.output.Contents) {
		return true
	}

	if c.output.IsTruncated == nil || !*c.output.IsTruncated {
		return false
	}

	c.index = -1
	c.output, c.err = c.client.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket:            c.output.Name,
		Prefix:            c.output.Prefix,
		MaxKeys:           c.output.MaxKeys,
		ContinuationToken: c.output.NextContinuationToken,
	})
	return c.err == nil
}

// Object returns current object
func (c *Cursor) Object() *s3.Object {
	return c.output.Contents[c.index]
}

// Err returns error from iterator
func (c *Cursor) Err() error {
	return c.err
}
