package config

import (
	"fmt"

	"github.com/syntaxiss/officeworld-sdk-go/pkg/requester"
)

type Option func(*Config) error

// WithHTTPClient allow to do requests using received http client.
func WithHTTPClient(r requester.Requester) Option {
	return func(c *Config) error {
		if r == nil {
			return fmt.Errorf("received http client can't be nil")
		}
		c.Requester = r
		return nil
	}
}
