package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Godzab/go-gpt3/internal"
	"github.com/Godzab/go-gpt3/pkg/models"
	"io"
	"net/http"
)

type Client interface {
	Call(request models.Request) (*models.Response, error)
	Setup(...string)
}

type Gpt3Client struct {
	apiKey     string
	apiBaseUrl string
	apiVersion string
	engines    []string
}

func (a *Gpt3Client) Call(request models.Request) (*models.Response, error) {
	var err error
	var req *http.Request

	config := models.RequestConfig{
		EndpointVersion: a.apiVersion,
		BaseUrl:         a.apiBaseUrl,
		Engine:          a.engines[0],
	}

	jsonStr, err := json.Marshal(request)

	if err != nil {
		_ = fmt.Errorf("Request marshalling error: %s\n", err)
		return nil, err
	}

	req, err = a.instantiateRequestObject(request, config, req, err, jsonStr)

	if err != nil {
		_ = fmt.Errorf("Http Request creation error: %s\n", err)
		return nil, err
	}

	prepareRequestHeaders(a.apiKey, req)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		_ = fmt.Errorf("Http request error: %s\n", err)
		return nil, err
	}

	defer resp.Body.Close()

	respObj := request.AttachResponse()
	data, err := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		_ = fmt.Errorf("Http response user error %d\n", resp.StatusCode)
		errObj := models.ErrorBag{}
		json.Unmarshal(data, &errObj)
		return nil, errObj
	}

	if err := json.Unmarshal(data, respObj); err != nil {
		_ = fmt.Errorf("Http response unmarshal error: %s\n", err)
		return nil, err
	}
	return &respObj, nil
}

func (a *Gpt3Client) instantiateRequestObject(request models.Request, config models.RequestConfig, req *http.Request, err error, jsonStr []byte) (*http.Request, error) {
	requestMethod, requestUrl := request.GetRequestMeta(config)

	if requestMethod == "GET" {
		req, err = http.NewRequest(requestMethod, requestUrl, nil)
	} else {
		req, err = http.NewRequest(requestMethod, requestUrl, bytes.NewBuffer(jsonStr))
	}
	return req, err
}

func prepareRequestHeaders(apiKey string, req *http.Request) {
	authHeader := fmt.Sprintf("Bearer %s", apiKey)
	req.Header.Set("Authorization", authHeader)
	req.Header.Set("Content-Type", "application/json")
}

func (a *Gpt3Client) Setup(engines ...string) {
	a.apiKey = internal.Config.Gpt3ApiKey
	a.apiVersion = internal.Config.Gpt3ApiVersion
	a.apiBaseUrl = internal.Config.Gpt3BaseUrl
	a.engines = append(a.engines, engines...)
}
