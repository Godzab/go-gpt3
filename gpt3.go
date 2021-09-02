//Package gpt3 provides access to the the GPT3 completions Api
//along with new beta APIs for classification, enhanced search, and question answering.
//
//The underlying structure is defined along a request / response interface pattern with a
//singular call to the client.
//The request is initialised as per required parameters an example being:
//
//	req := gpt3.CompletionRequest{
//		Prompt:      string(query),
//		MaxTokens:   60,
//		TopP:        1,
//		Temperature: 0.3,
//		FrequencyPenalty: 0.5,
//		PresencePenalty: 0,
//		Stop: []string{"You:"},
//	}
//
//The content filter endpoint is used to validate a prompt in order to safeguard responses ushered back to the enduser.
//The request object should always have the following parameters:
//
// reformattedPrompt := fmt.Sprintf("<|endoftext|>[%s]\n--\nLabel:", string(query))
//
// req := gpt3.ContentFilterRequest{
//		Prompt:      reformattedPrompt,
//		MaxTokens:   1,
//		TopP:        0,
//		Temperature: 0,
//		Logprobs: 10,
//	}
//
// The Response is the same format as that of the Completions request with the following entries:
//
// 0 => text is safe
// 1 => This text is sensitive. This means that the text could be talking about a sensitive topic, something political,
//		religious, or talking about a protected class such as race or nationality.
// 2 => This text is unsafe. This means that the text contains profane language, prejudiced or hateful language,
//      something that could be NSFW, or text that portrays certain groups/people in a harmful manner.
//
// Code Generation:
//
// Added to the completions API are the codex engines for code generation.
// The Codex model series is a descendant of our base GPT-3 series that’s been trained on both
// natural language and billions of lines of code.

package gpt3

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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
		//Perform single action initialisations
	})

	a.apiKey = os.Getenv(apiKeyName)
	a.engines = append(a.engines, engines...)
	return a
}
