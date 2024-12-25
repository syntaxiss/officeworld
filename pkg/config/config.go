package config

import (
	"fmt"

	"github.com/syntaxiss/officeworld-sdk-go/pkg/internal/defaultrequester"
	"github.com/syntaxiss/officeworld-sdk-go/pkg/requester"
)

// Config allows you to send custom settings and API authentication
type Configuracion struct {
	Requester requester.Requester

	Authorization string
}

// New returns a new Config.
func New(authorization string, opts ...Option) (*Configuracion, error) {
	config := &Configuracion{
		Authorization: authorization,
		Requester:     defaultrequester.New(),
	}

	// Apply all the functional options to configure the client.
	for _, opt := range opts {
		if err := opt(config); err != nil {
			return nil, fmt.Errorf("fail to build config: %w", err)
		}
	}

	return config, nil
}
