package client

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
)

// Client wraps EC2Metadata clients
type Client struct {
	EC2Metadata *ec2metadata.EC2Metadata
}

// GetList returns list of entries by path
func (c *Client) GetList(p string) ([]string, error) {
	m, err := c.EC2Metadata.GetMetadata(p)

	if err != nil {
		return nil, err
	}

	return strings.Split(m, "\n"), nil
}

// GetString returns single entry by path
func (c *Client) GetString(p string) (string, error) {
	m, err := c.EC2Metadata.GetMetadata(p)

	if err != nil {
		return "", err
	}

	return m, nil
}

// ResolvePath returns all entries matching path pattern
func (c *Client) ResolvePath(p string) ([]string, error) {
	p = strings.Trim(p, "/")

	// wrap GetList to ignore 404 errors
	f := func(p string) ([]string, error) {
		list, err := c.GetList(p)

		if reqErr, ok := err.(awserr.RequestFailure); ok {
			if reqErr.StatusCode() != 404 {
				return nil, err
			}
		}

		return list, nil
	}

	return PathResolver(p, f)
}

// New returns new client
func New() (*Client, error) {
	sess, err := session.NewSession()

	if err != nil {
		return nil, err
	}

	return &Client{ec2metadata.New(sess)}, nil
}
