package client

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
)

type Client struct {
	EC2Metadata *ec2metadata.EC2Metadata
}

func (c *Client) GetList(p string) ([]string, error) {
	m, err := c.EC2Metadata.GetMetadata(p)

	if err != nil {
		return nil, err
	}

	return strings.Split(m, "\n"), nil
}

func (c *Client) GetString(p string) (string, error) {
	m, err := c.EC2Metadata.GetMetadata(p)

	if err != nil {
		return "", err
	}

	return m, nil
}

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

func New() (*Client, error) {
	sess, err := session.NewSession()

	if err != nil {
		return nil, err
	}

	return &Client{ec2metadata.New(sess)}, nil
}
