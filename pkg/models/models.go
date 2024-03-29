package models

import (
	"fmt"
	"os"
)

const (
	DAVINCI               = "text-davinci-002"
	CURIE                 = "curie"
	BABBAGE               = "babbage"
	BABBAGE_INSTRUCT_BETA = "text-babbage-001"
	ADA                   = "ada"
	ADA_INSTRUCT_BETA     = "text-ada-001"
	CURIE_INSTRUCT_BETA   = "text-curie-001"
	DAVINCI_INSTRUCT_BETA = "text-davinci-001"

	// CURSING_FILTER_V6 Content filters moderate output and input to the api to
	//avoid negative content generation
	CURSING_FILTER_V6       = "cursing-filter-v6"
	CONTENT_FILTER_DEV      = "content-filter-dev"
	CONTENT_FILTER_ALPHA_C4 = "content-filter-alpha-c4"

	// DAVINCI_CODEX Codex engines for code generation.
	//Davinci Codex is more capable, particularly for translating natural language to code
	DAVINCI_CODEX = "davinci-codex"

	// CUSHMAN_CODEX Cushman Codex is almost as capable, but slightly faster.
	//This speed advantage may be preferable for real-time applications.
	CUSHMAN_CODEX = "cushman-codex"
)

const (
	getRequest  = "GET"
	postRequest = "POST"
)

const (
	ANSWERS         = "answers"
	SEARCH          = "search"
	CLASSIFICATIONS = "classifications"
)

type RequestConfig struct {
	EndpointVersion, BaseUrl, Engine string
}

type Request interface {
	AttachResponse() Response
	GetRequestMeta(config RequestConfig) (string, string)
}

type Response interface {
	GetBody() Response
}

type Document struct {
	Document int     `json:"document"`
	Object   string  `json:"object,omitempty"`
	Score    float64 `json:"score,omitempty"`
	Text     string  `json:"text"`
}

type File struct {
	ID        string `json:"id"`
	Object    string `json:"object"`
	Bytes     int    `json:"bytes"`
	CreatedAt int    `json:"created_at"`
	Filename  string `json:"filename"`
	Purpose   string `json:"purpose"`
}

type Choices struct {
	Text         string        `json:"text"`
	Index        int           `json:"index"`
	Logprobs     LogprobResult `json:"logprobs"`
	FinishReason string        `json:"finish_reason"`
}

type LogprobResult struct {
	Tokens        []string             `json:"tokens"`
	TokenLogprobs []float32            `json:"token_logprobs"`
	TopLogprobs   []map[string]float32 `json:"top_logprobs"`
	TextOffset    []int                `json:"text_offset"`
}

type SearchData struct {
	Document
	Object string  `json:"object"`
	Score  float64 `json:"score"`
}

type ClassificationExamples struct {
	Document
	Label string `json:"label"`
}

type Engine struct {
	ID     string `json:"id"`
	Object string `json:"object"`
	Owner  string `json:"owner"`
	Ready  bool   `json:"ready"`
}

// Files models
type FilesRequest struct{}

type FilesResponse struct {
	Data   []File `json:"data"`
	Object string `json:"object"`
}

func (r *FilesRequest) AttachResponse() Response {
	resp := &FilesResponse{}
	return resp
}

func (r *FilesResponse) GetBody() Response {
	return r
}

func (r *FilesRequest) GetRequestMeta(config RequestConfig) (string, string) {
	return getRequest, fmt.Sprintf("%s/%s/files", config.BaseUrl, config.EndpointVersion)
}

// File models
type FileRequest struct {
	File    os.File `json:"file"`
	Purpose string  `json:"purpose"`
}

type FileResponse struct {
	File
}

func (r *FileRequest) AttachResponse() Response {
	resp := &FileResponse{}
	return resp
}

func (r *FileResponse) GetBody() Response {
	return r
}

func (r *FileRequest) GetRequestMeta(config RequestConfig) (string, string) {
	return postRequest, fmt.Sprintf("%s/%s/files", config.BaseUrl, config.EndpointVersion)
}

// CompletionRequest Completion model structures
type CompletionRequest struct {
	Prompt           string          `json:"prompt"`
	MaxTokens        int             `json:"max_tokens"`
	Temperature      float32         `json:"temperature,omitempty"`
	TopP             float32         `json:"top_p,omitempty"`
	N                int             `json:"n,omitempty"`
	Stream           bool            `json:"stream"`
	Logprobs         int             `json:"logprobs,omitempty"`
	Stop             []string        `json:"stop,omitempty"`
	Echo             bool            `json:"echo,omitempty"`
	PresencePenalty  float32         `json:"presence_penalty,omitempty"`
	FrequencyPenalty float32         `json:"frequency_penalty,omitempty"`
	BestOf           float32         `json:"best_of,omitempty"`
	LogitBias        map[string]int8 `json:"logit_bias,omitempty"`
}

type CompletionResponse struct {
	ID      string    `json:"id"`
	Object  string    `json:"object"`
	Created int       `json:"created"`
	Model   string    `json:"model"`
	Choices []Choices `json:"choices"`
}

func (r *CompletionRequest) AttachResponse() Response {
	resp := &CompletionResponse{}
	return resp
}

func (r *CompletionRequest) GetRequestMeta(config RequestConfig) (string, string) {
	return postRequest, fmt.Sprintf("%s/%s/engines/%s/completions", config.BaseUrl, config.EndpointVersion, config.Engine)
}

func (r *CompletionResponse) GetBody() Response {
	return r
}

// ContentFilterRequest Content filter model structures
type ContentFilterRequest struct {
	Prompt           string  `json:"prompt"`
	MaxTokens        int     `json:"max_tokens"`
	Temperature      float32 `json:"temperature,omitempty"`
	TopP             float32 `json:"top_p,omitempty"`
	N                int     `json:"n,omitempty"`
	Logprobs         int     `json:"logprobs,omitempty"`
	PresencePenalty  float32 `json:"presence_penalty,omitempty"`
	FrequencyPenalty float32 `json:"frequency_penalty,omitempty"`
}

func (r *ContentFilterRequest) AttachResponse() Response {
	resp := &CompletionResponse{}
	return resp
}

func (r *ContentFilterRequest) GetRequestMeta(config RequestConfig) (string, string) {
	return postRequest, fmt.Sprintf("%s/%s/engines/content-filter-alpha-c4/completions", config.BaseUrl, config.EndpointVersion)
}

// SearchRequest Search Model structures
type SearchRequest struct {
	target         string
	Documents      []string `json:"documents,omitempty"`
	Query          string   `json:"query"`
	File           string   `json:"file,omitempty"`
	ReturnMetadata bool     `json:"return_metadata"`
	MaxRerank      int32    `json:"max_rerank,omitempty"`
}

type SearchResponse struct {
	Data   []SearchData `json:"data"`
	Object string       `json:"object"`
}

func (r *SearchRequest) AttachResponse() Response {
	resp := &SearchResponse{}
	return resp
}

func (r *SearchRequest) GetRequestMeta(config RequestConfig) (string, string) {
	return postRequest, fmt.Sprintf("%s/%s/engines/%s/search", config.BaseUrl, config.EndpointVersion, config.Engine)
}

func (r *SearchResponse) GetBody() Response {
	return r
}

type EnginesRequest struct{}

type EnginesResponse struct {
	Data   []interface{} `json:"data"`
	Object string        `json:"object"`
}

func (e EnginesResponse) GetBody() Response {
	return e
}

func (r *EnginesRequest) AttachResponse() Response {
	resp := &EnginesResponse{}
	return resp
}

func (r *EnginesRequest) GetRequestMeta(config RequestConfig) (string, string) {
	return getRequest, fmt.Sprintf("%s/%s/engines", config.BaseUrl, config.EndpointVersion)
}

// ClassificationRequest Classification Model structures
type ClassificationRequest struct {
	Examples       [][]string      `json:"examples"`
	Labels         []string        `json:"labels"`
	Query          string          `json:"query"`
	File           string          `json:"file"`
	SearchModel    string          `json:"search_model"`
	Model          string          `json:"model"`
	Temperature    float32         `json:"temperature"`
	Logprobs       interface{}     `json:"logprobs,omitempty"`
	MaxExamples    int32           `json:"max_examples"`
	LogitBias      map[string]int8 `json:"logit_bias,omitempty"`
	ReturnPrompt   bool            `json:"return_prompt,omitempty"`
	ReturnMetadata bool            `json:"return_metadata,omitempty"`
	Expand         []string        `json:"expand,omitempty"`
}

type ClassificationResponse struct {
	Completion       string                   `json:"completion"`
	Label            string                   `json:"label"`
	Model            string                   `json:"model"`
	Object           string                   `json:"object"`
	SearchModel      string                   `json:"search_model"`
	SelectedExamples []ClassificationExamples `json:"selected_examples"`
}

func (r *ClassificationRequest) AttachResponse() Response {
	resp := &ClassificationResponse{}
	return resp
}

func (r *ClassificationRequest) GetRequestMeta(config RequestConfig) (string, string) {
	return postRequest, fmt.Sprintf("%s/%s/classifications", config.BaseUrl, config.EndpointVersion)
}

func (r *ClassificationResponse) GetBody() Response {
	return r
}

// AnswerRequest Answer Model structures
type AnswerRequest struct {
	Documents       []string        `json:"documents"`
	Question        string          `json:"question"`
	SearchModel     string          `json:"search_model"`
	Model           string          `json:"model"`
	ExamplesContext string          `json:"examples_context"`
	Examples        [][]string      `json:"examples"`
	MaxTokens       int             `json:"max_tokens"`
	Stop            []string        `json:"stop"`
	File            string          `json:"file,omitempty"`
	MaxRerank       int32           `json:"max_rerank"`
	Temperature     float32         `json:"temperature"`
	Logprobs        interface{}     `json:"logprobs,omitempty"`
	N               int             `json:"n,omitempty"`
	LogitBias       map[string]int8 `json:"logit_bias,omitempty"`
	ReturnPrompt    bool            `json:"return_prompt"`
	ReturnMetadata  bool            `json:"return_metadata"`
	Expand          []string        `json:"expand,omitempty"`
}

type AnswerResponse struct {
	Answers           []string           `json:"answers"`
	Completion        CompletionResponse `json:"completion"`
	Model             string             `json:"model"`
	Object            string             `json:"object"`
	SearchModel       string             `json:"search_model"`
	SelectedDocuments []Document         `json:"selected_documents"`
}

func (r *AnswerRequest) AttachResponse() Response {
	resp := &AnswerResponse{}
	return resp
}

func (r *AnswerRequest) GetRequestMeta(config RequestConfig) (string, string) {
	return postRequest, fmt.Sprintf("%s/%s/answers", config.BaseUrl, config.EndpointVersion)
}

func (r *AnswerResponse) GetBody() Response {
	return r
}

// GptErrorResponse Error handling for client calls
type GptErrorResponse struct {
	Code    interface{} `json:"code"`
	Message string      `json:"message"`
	Param   string      `json:"param"`
	Type    string      `json:"type"`
}

type ErrorBag struct {
	Err GptErrorResponse `json:"error"`
}

func (e ErrorBag) Error() string {
	return fmt.Sprintf("[GPT ERROR] %v:  %s %s %v",
		e.Err.Code, e.Err.Type, e.Err.Param, e.Err.Message)
}

func (e ErrorBag) Timeout() bool {
	return true
}

func (e ErrorBag) Temporary() bool {
	return true
}
