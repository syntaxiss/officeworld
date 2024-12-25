package client

import (
	"context"
	"net/http"

	"github.com/syntaxiss/officeworld-sdk-go/pkg/config"
	"github.com/syntaxiss/officeworld-sdk-go/pkg/internal/httpclient"
)

const (
	urlBase   = "https://api.worldoffice.cloud/api/v1/documentos/"
	urlCreate = urlBase + "crearDocumentoVenta"
	urlGet    = urlBase + "getDocumentoId/" + "/{id}"
	urlUpdate = urlBase + "editarDocumentoVenta"
)

type Client interface {
	Create(ctx context.Context, request Request) (*Response, error) // Este no pide ID porque lo crea

	Get(ctx context.Context, id string) (*Response, error)

	Update(ctx context.Context, request Request) (*Response, error) // Este pide ID porque lo actualiza
}

type client struct {
	cfg *config.Config
}

func NewClient(c *config.Config) Client {
	return &client{
		cfg: c,
	}
}

func (c *client) Create(ctx context.Context, request Request) (*Response, error) {
	requestData := httpclient.RequestData{
		Body:   request,
		Method: http.MethodPost,
		URL:    urlCreate,
	}
	resource, err := httpclient.DoRequest[*Response](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) Get(ctx context.Context, id string) (*Response, error) {
	pathParams := map[string]string{
		"id": id,
	}

	requestData := httpclient.RequestData{
		Method:     http.MethodGet,
		URL:        urlGet,
		PathParams: pathParams,
	}
	resource, err := httpclient.DoRequest[*Response](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) Update(ctx context.Context, request Request) (*Response, error) {
	requestData := httpclient.RequestData{
		Body:   request,
		Method: http.MethodPut,
		URL:    urlUpdate,
	}
	resource, err := httpclient.DoRequest[*Response](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}
