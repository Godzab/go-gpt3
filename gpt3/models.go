package gpt3

import (
	"fmt"
)

const (
	DAVINCI                 = "davinci"
	CURIE                   = "curie"
	BABBAGE                 = "babbage"
	ADA                     = "ada"
	CURIE_INSTRUCT_BETA     = "curie-instruct-beta"
	DAVINCI_INSTRUCT_BETA   = "davinci-instruct-beta"
	CURSING_FILTER_V6       = "cursing-filter-v6"
	CONTENT_FILTER_DEV      = "content-filter-dev"
	CONTENT_FILTER_ALPHA_C4 = "content-filter-alpha"
)

type RequestConfig struct {
	endpointVersion, baseUrl, engine string
}

type Request interface {
	attachResponse() Response
	getRequestUrl(config RequestConfig) string
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

// Files models
type Files struct{}

type FilesResponse struct {
	Data   []File `json:"data"`
	Object string `json:"object"`
}

func (r *Files) attachResponse() Response {
	resp := &FilesResponse{}
	return resp
}

func (r *FilesResponse) GetBody() Response {
	return r
}

func (r *Files) getRequestUrl(config RequestConfig) string {
	return fmt.Sprintf("%s/%s/files", config.baseUrl, config.endpointVersion)
}

// CompletionRequest Completion model structures
type CompletionRequest struct {
	Prompt           string          `json:"prompt"`
	MaxTokens        int             `json:"max_tokens"`
	Temperature      float32         `json:"temperature"`
	TopP             float32         `json:"top_p"`
	N                int             `json:"n"`
	Stream           bool            `json:"stream"`
	Logprobs         int             `json:"logprobs"`
	Stop             []string        `json:"stop"`
	Echo             bool            `json:"echo"`
	PresencePenalty  float32         `json:"presence_penalty"`
	FrequencyPenalty float32         `json:"frequency_penalty"`
	BestOf           float32         `json:"best_of"`
	LogitBias        map[string]int8 `json:"logit_bias"`
}

type CompletionResponse struct {
	ID      string    `json:"id"`
	Object  string    `json:"object"`
	Created int       `json:"created"`
	Model   string    `json:"model"`
	Choices []Choices `json:"choices"`
}

func (r *CompletionRequest) attachResponse() Response {
	resp := &CompletionResponse{}
	return resp
}

func (r *CompletionRequest) getRequestUrl(config RequestConfig) string {
	return fmt.Sprintf("%s/%s/engines/%s/completions", config.baseUrl, config.endpointVersion, config.engine)
}

func (r *CompletionResponse) GetBody() Response {
	return r
}

// SearchRequest Search Model structures
type SearchRequest struct {
	target         string
	Documents      []string `json:"documents"`
	Query          string   `json:"query"`
	File           string   `json:"file"`
	ReturnMetadata bool     `json:"return_metadata"`
	MaxRerank      int32    `json:"max_rerank"`
}

type SearchResponse struct {
	Data   []SearchData `json:"data"`
	Object string       `json:"object"`
}

func (r *SearchRequest) attachResponse() Response {
	resp := &SearchResponse{}
	return resp
}

func (r *SearchRequest) getRequestUrl(config RequestConfig) string {
	return fmt.Sprintf("%s/%s/engines/%s/search", config.baseUrl, config.endpointVersion, config.engine)
}

func (r *SearchResponse) GetBody() Response {
	return r
}

// ClassificationRequest Classification Model structures
type ClassificationRequest struct {
	target         string
	Examples       [][]string      `json:"examples"`
	Labels         []string        `json:"labels"`
	Query          string          `json:"query"`
	File           string          `json:"file"`
	SearchModel    string          `json:"search_model"`
	Model          string          `json:"model"`
	Temperature    float32         `json:"temperature"`
	Logprobs       interface{}     `json:"logprobs"`
	MaxExamples    int32           `json:"max_examples"`
	LogitBias      map[string]int8 `json:"logit_bias"`
	ReturnPrompt   bool            `json:"return_prompt"`
	ReturnMetadata bool            `json:"return_metadata"`
	Expand         []interface{}   `json:"expand"`
}

type ClassificationResponse struct {
	Completion       string                   `json:"completion"`
	Label            string                   `json:"label"`
	Model            string                   `json:"model"`
	Object           string                   `json:"object"`
	SearchModel      string                   `json:"search_model"`
	SelectedExamples []ClassificationExamples `json:"selected_examples"`
}

func (r *ClassificationRequest) attachResponse() Response {
	resp := &ClassificationResponse{}
	return resp
}

func (r *ClassificationRequest) getRequestUrl(config RequestConfig) string {
	return fmt.Sprintf("%s/%s/classifications", config.baseUrl, config.endpointVersion)
}

func (r *ClassificationResponse) GetBody() Response {
	return r
}

// AnswerRequest Answer Model structures
type AnswerRequest struct {
	target          string
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

func (r *AnswerRequest) attachResponse() Response {
	resp := &AnswerResponse{}
	return resp
}

func (r *AnswerRequest) getRequestUrl(config RequestConfig) string {
	return fmt.Sprintf("%s/%s/answers", config.baseUrl, config.endpointVersion)
}

func (r *AnswerResponse) GetBody() Response {
	return r
}

type Error struct {
	Code    interface{} `json:"code"`
	Message string      `json:"message"`
	Param   interface{} `json:"param"`
	Type    string      `json:"type"`
}
