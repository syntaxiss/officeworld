package httpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/syntaxiss/officeworld-sdk-go/pkg/config"
)

var (
	pathParamRegexp = regexp.MustCompile(`{[^{}]*}`)
)

type RequestData struct {
	Body any

	Method      string
	URL         string
	PathParams  map[string]string
	QueryParams map[string]string
}

func DoRequest[T any](ctx context.Context, config *config.Configuracion, requestData RequestData) (T, error) {
	var resource T

	req, err := createRequest(ctx, config, requestData)
	if err != nil {
		return resource, err
	}

	b, err := Send(config.Requester, req)
	if err != nil {
		return resource, err
	}

	return unmarshal(b, resource)
}

func createRequest(ctx context.Context, config *config.Configuracion, requestData RequestData) (*http.Request, error) {
	body, err := marshal(requestData.Body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, requestData.Method, requestData.URL, body)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	setHeaders(req, config)
	if err = setPathParams(req, requestData.PathParams); err != nil {
		return nil, err
	}
	setQueryParams(req, requestData.QueryParams)

	return req, nil
}

func setHeaders(req *http.Request, config *config.Configuracion) {
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Authorization", "WO "+config.Authorization)
}

func setPathParams(req *http.Request, params map[string]string) error {
	pathURL := req.URL.Path

	for k, v := range params {
		pathParam := "{" + k + "}"
		pathURL = strings.Replace(pathURL, pathParam, v, 1)
	}

	matches := pathParamRegexp.FindAllString(pathURL, -1)
	if matches != nil {
		return fmt.Errorf("the following parameters weren't replaced: %v", matches)
	}

	req.URL.Path = pathURL
	return nil
}

func setQueryParams(req *http.Request, params map[string]string) {
	if len(params) == 0 {
		return
	}

	queryParams := url.Values{}
	for k, v := range params {
		queryParams.Add(k, v)
	}
	req.URL.RawQuery = queryParams.Encode()
}

func marshal(body any) (io.Reader, error) {
	if body == nil {
		return nil, nil
	}

	b, err := json.Marshal(&body)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request body: %w", err)
	}

	return strings.NewReader(string(b)), nil
}

func unmarshal[T any](b []byte, response T) (T, error) {
	if err := json.Unmarshal(b, &response); err != nil {
		return response, err
	}

	return response, nil
}
