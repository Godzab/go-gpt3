package gpt3

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
)

var once sync.Once

const (
	baseUrl            = "https://api.openai.com"
	defaultVersion     = "v1"
	apiKeyName         = "OPENAI_API_KEY"
	apiKeyMissingError = "Api key required. Please ensure env variable %s is set."
)

type Client interface {
	Call(request Request) (*Response, error)
	Setup(...string)
}

type ApiClient struct {
	apiKey  string
	engines []string
}

func (a ApiClient) Call(request Request) (*Response, error) {
	var err error
	var req *http.Request

	config := RequestConfig{
		endpointVersion: defaultVersion,
		baseUrl:         baseUrl,
		engine:          a.engines[0],
	}

	jsonStr, err := json.Marshal(request)
	if err != nil {
		_ = fmt.Errorf("Request marshalling error: %s\n", err)
		return nil, err
	}

	requestMethod, requestUrl := request.getRequestMeta(config)

	if requestMethod == getRequest{
		req, err = http.NewRequest(requestMethod, requestUrl, nil)
	} else {
		req, err = http.NewRequest(requestMethod, requestUrl, bytes.NewBuffer(jsonStr))
	}

	if err != nil {
		_ = fmt.Errorf("Http Request creation error: %s\n", err)
		return nil, err
	}

	authHeader := fmt.Sprintf("Bearer %s", a.apiKey)
	req.Header.Set("Authorization", authHeader)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		_ = fmt.Errorf("Http request error: %s\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	respObj := request.attachResponse()
	data, err := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK{
		_ = fmt.Errorf("Http response user error %d\n", resp.StatusCode)
		errObj := ErrorBag{}
		json.Unmarshal(data, &errObj)
		return nil, errObj
	}

	if err := json.Unmarshal(data, respObj); err != nil {
		_ = fmt.Errorf("Http response unmarshal error: %s\n", err)
		return nil, err
	}
	return &respObj, nil
}

func (a *ApiClient) Setup(engines ...string) *ApiClient {
	once.Do(func() {
		a.apiKey = os.Getenv(apiKeyName)
		if a.apiKey == "" {
			log.Fatalf(apiKeyMissingError, apiKeyName)
		}
	})
	//TODO: For concurrent calls to different engines for the same request
	a.engines = append(a.engines, engines...)
	return a
}
