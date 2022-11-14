package models

import (
	"os"
	"reflect"
	"testing"
)

func TestAnswerRequest_attachResponse(t *testing.T) {
	type fields struct {
		Documents       []string
		Question        string
		SearchModel     string
		Model           string
		ExamplesContext string
		Examples        [][]string
		MaxTokens       int
		Stop            []string
		File            string
		MaxRerank       int32
		Temperature     float32
		Logprobs        interface{}
		N               int
		LogitBias       map[string]int8
		ReturnPrompt    bool
		ReturnMetadata  bool
		Expand          []string
	}
	tests := []struct {
		name   string
		fields fields
		want   Response
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AnswerRequest{
				Documents:       tt.fields.Documents,
				Question:        tt.fields.Question,
				SearchModel:     tt.fields.SearchModel,
				Model:           tt.fields.Model,
				ExamplesContext: tt.fields.ExamplesContext,
				Examples:        tt.fields.Examples,
				MaxTokens:       tt.fields.MaxTokens,
				Stop:            tt.fields.Stop,
				File:            tt.fields.File,
				MaxRerank:       tt.fields.MaxRerank,
				Temperature:     tt.fields.Temperature,
				Logprobs:        tt.fields.Logprobs,
				N:               tt.fields.N,
				LogitBias:       tt.fields.LogitBias,
				ReturnPrompt:    tt.fields.ReturnPrompt,
				ReturnMetadata:  tt.fields.ReturnMetadata,
				Expand:          tt.fields.Expand,
			}
			if got := r.attachResponse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("attachResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnswerRequest_getRequestMeta(t *testing.T) {
	type fields struct {
		Documents       []string
		Question        string
		SearchModel     string
		Model           string
		ExamplesContext string
		Examples        [][]string
		MaxTokens       int
		Stop            []string
		File            string
		MaxRerank       int32
		Temperature     float32
		Logprobs        interface{}
		N               int
		LogitBias       map[string]int8
		ReturnPrompt    bool
		ReturnMetadata  bool
		Expand          []string
	}
	type args struct {
		config RequestConfig
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		want1  string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AnswerRequest{
				Documents:       tt.fields.Documents,
				Question:        tt.fields.Question,
				SearchModel:     tt.fields.SearchModel,
				Model:           tt.fields.Model,
				ExamplesContext: tt.fields.ExamplesContext,
				Examples:        tt.fields.Examples,
				MaxTokens:       tt.fields.MaxTokens,
				Stop:            tt.fields.Stop,
				File:            tt.fields.File,
				MaxRerank:       tt.fields.MaxRerank,
				Temperature:     tt.fields.Temperature,
				Logprobs:        tt.fields.Logprobs,
				N:               tt.fields.N,
				LogitBias:       tt.fields.LogitBias,
				ReturnPrompt:    tt.fields.ReturnPrompt,
				ReturnMetadata:  tt.fields.ReturnMetadata,
				Expand:          tt.fields.Expand,
			}
			got, got1 := r.getRequestMeta(tt.args.config)
			if got != tt.want {
				t.Errorf("getRequestMeta() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getRequestMeta() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestAnswerResponse_GetBody(t *testing.T) {
	type fields struct {
		Answers           []string
		Completion        CompletionResponse
		Model             string
		Object            string
		SearchModel       string
		SelectedDocuments []Document
	}
	tests := []struct {
		name   string
		fields fields
		want   Response
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AnswerResponse{
				Answers:           tt.fields.Answers,
				Completion:        tt.fields.Completion,
				Model:             tt.fields.Model,
				Object:            tt.fields.Object,
				SearchModel:       tt.fields.SearchModel,
				SelectedDocuments: tt.fields.SelectedDocuments,
			}
			if got := r.GetBody(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBody() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClassificationRequest_attachResponse(t *testing.T) {
	type fields struct {
		Examples       [][]string
		Labels         []string
		Query          string
		File           string
		SearchModel    string
		Model          string
		Temperature    float32
		Logprobs       interface{}
		MaxExamples    int32
		LogitBias      map[string]int8
		ReturnPrompt   bool
		ReturnMetadata bool
		Expand         []string
	}
	tests := []struct {
		name   string
		fields fields
		want   Response
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ClassificationRequest{
				Examples:       tt.fields.Examples,
				Labels:         tt.fields.Labels,
				Query:          tt.fields.Query,
				File:           tt.fields.File,
				SearchModel:    tt.fields.SearchModel,
				Model:          tt.fields.Model,
				Temperature:    tt.fields.Temperature,
				Logprobs:       tt.fields.Logprobs,
				MaxExamples:    tt.fields.MaxExamples,
				LogitBias:      tt.fields.LogitBias,
				ReturnPrompt:   tt.fields.ReturnPrompt,
				ReturnMetadata: tt.fields.ReturnMetadata,
				Expand:         tt.fields.Expand,
			}
			if got := r.attachResponse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("attachResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClassificationRequest_getRequestMeta(t *testing.T) {
	type fields struct {
		Examples       [][]string
		Labels         []string
		Query          string
		File           string
		SearchModel    string
		Model          string
		Temperature    float32
		Logprobs       interface{}
		MaxExamples    int32
		LogitBias      map[string]int8
		ReturnPrompt   bool
		ReturnMetadata bool
		Expand         []string
	}
	type args struct {
		config RequestConfig
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		want1  string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ClassificationRequest{
				Examples:       tt.fields.Examples,
				Labels:         tt.fields.Labels,
				Query:          tt.fields.Query,
				File:           tt.fields.File,
				SearchModel:    tt.fields.SearchModel,
				Model:          tt.fields.Model,
				Temperature:    tt.fields.Temperature,
				Logprobs:       tt.fields.Logprobs,
				MaxExamples:    tt.fields.MaxExamples,
				LogitBias:      tt.fields.LogitBias,
				ReturnPrompt:   tt.fields.ReturnPrompt,
				ReturnMetadata: tt.fields.ReturnMetadata,
				Expand:         tt.fields.Expand,
			}
			got, got1 := r.getRequestMeta(tt.args.config)
			if got != tt.want {
				t.Errorf("getRequestMeta() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getRequestMeta() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestClassificationResponse_GetBody(t *testing.T) {
	type fields struct {
		Completion       string
		Label            string
		Model            string
		Object           string
		SearchModel      string
		SelectedExamples []ClassificationExamples
	}
	tests := []struct {
		name   string
		fields fields
		want   Response
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ClassificationResponse{
				Completion:       tt.fields.Completion,
				Label:            tt.fields.Label,
				Model:            tt.fields.Model,
				Object:           tt.fields.Object,
				SearchModel:      tt.fields.SearchModel,
				SelectedExamples: tt.fields.SelectedExamples,
			}
			if got := r.GetBody(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBody() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompletionRequest_attachResponse(t *testing.T) {
	type fields struct {
		Prompt           string
		MaxTokens        int
		Temperature      float32
		TopP             float32
		N                int
		Stream           bool
		Logprobs         int
		Stop             []string
		Echo             bool
		PresencePenalty  float32
		FrequencyPenalty float32
		BestOf           float32
		LogitBias        map[string]int8
	}
	tests := []struct {
		name   string
		fields fields
		want   Response
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &CompletionRequest{
				Prompt:           tt.fields.Prompt,
				MaxTokens:        tt.fields.MaxTokens,
				Temperature:      tt.fields.Temperature,
				TopP:             tt.fields.TopP,
				N:                tt.fields.N,
				Stream:           tt.fields.Stream,
				Logprobs:         tt.fields.Logprobs,
				Stop:             tt.fields.Stop,
				Echo:             tt.fields.Echo,
				PresencePenalty:  tt.fields.PresencePenalty,
				FrequencyPenalty: tt.fields.FrequencyPenalty,
				BestOf:           tt.fields.BestOf,
				LogitBias:        tt.fields.LogitBias,
			}
			if got := r.attachResponse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("attachResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompletionRequest_getRequestMeta(t *testing.T) {
	type fields struct {
		Prompt           string
		MaxTokens        int
		Temperature      float32
		TopP             float32
		N                int
		Stream           bool
		Logprobs         int
		Stop             []string
		Echo             bool
		PresencePenalty  float32
		FrequencyPenalty float32
		BestOf           float32
		LogitBias        map[string]int8
	}
	type args struct {
		config RequestConfig
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		want1  string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &CompletionRequest{
				Prompt:           tt.fields.Prompt,
				MaxTokens:        tt.fields.MaxTokens,
				Temperature:      tt.fields.Temperature,
				TopP:             tt.fields.TopP,
				N:                tt.fields.N,
				Stream:           tt.fields.Stream,
				Logprobs:         tt.fields.Logprobs,
				Stop:             tt.fields.Stop,
				Echo:             tt.fields.Echo,
				PresencePenalty:  tt.fields.PresencePenalty,
				FrequencyPenalty: tt.fields.FrequencyPenalty,
				BestOf:           tt.fields.BestOf,
				LogitBias:        tt.fields.LogitBias,
			}
			got, got1 := r.getRequestMeta(tt.args.config)
			if got != tt.want {
				t.Errorf("getRequestMeta() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getRequestMeta() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCompletionResponse_GetBody(t *testing.T) {
	type fields struct {
		ID      string
		Object  string
		Created int
		Model   string
		Choices []Choices
	}
	tests := []struct {
		name   string
		fields fields
		want   Response
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &CompletionResponse{
				ID:      tt.fields.ID,
				Object:  tt.fields.Object,
				Created: tt.fields.Created,
				Model:   tt.fields.Model,
				Choices: tt.fields.Choices,
			}
			if got := r.GetBody(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBody() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContentFilterRequest_attachResponse(t *testing.T) {
	type fields struct {
		Prompt           string
		MaxTokens        int
		Temperature      float32
		TopP             float32
		N                int
		Logprobs         int
		PresencePenalty  float32
		FrequencyPenalty float32
	}
	tests := []struct {
		name   string
		fields fields
		want   Response
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ContentFilterRequest{
				Prompt:           tt.fields.Prompt,
				MaxTokens:        tt.fields.MaxTokens,
				Temperature:      tt.fields.Temperature,
				TopP:             tt.fields.TopP,
				N:                tt.fields.N,
				Logprobs:         tt.fields.Logprobs,
				PresencePenalty:  tt.fields.PresencePenalty,
				FrequencyPenalty: tt.fields.FrequencyPenalty,
			}
			if got := r.attachResponse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("attachResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContentFilterRequest_getRequestMeta(t *testing.T) {
	type fields struct {
		Prompt           string
		MaxTokens        int
		Temperature      float32
		TopP             float32
		N                int
		Logprobs         int
		PresencePenalty  float32
		FrequencyPenalty float32
	}
	type args struct {
		config RequestConfig
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		want1  string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ContentFilterRequest{
				Prompt:           tt.fields.Prompt,
				MaxTokens:        tt.fields.MaxTokens,
				Temperature:      tt.fields.Temperature,
				TopP:             tt.fields.TopP,
				N:                tt.fields.N,
				Logprobs:         tt.fields.Logprobs,
				PresencePenalty:  tt.fields.PresencePenalty,
				FrequencyPenalty: tt.fields.FrequencyPenalty,
			}
			got, got1 := r.getRequestMeta(tt.args.config)
			if got != tt.want {
				t.Errorf("getRequestMeta() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getRequestMeta() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestEnginesRequest_attachResponse(t *testing.T) {
	tests := []struct {
		name string
		want Response
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &EnginesRequest{}
			if got := r.attachResponse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("attachResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnginesRequest_getRequestMeta(t *testing.T) {
	type args struct {
		config RequestConfig
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &EnginesRequest{}
			got, got1 := r.getRequestMeta(tt.args.config)
			if got != tt.want {
				t.Errorf("getRequestMeta() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getRequestMeta() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestEnginesResponse_GetBody(t *testing.T) {
	type fields struct {
		Data   []interface{}
		Object string
	}
	tests := []struct {
		name   string
		fields fields
		want   Response
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := EnginesResponse{
				Data:   tt.fields.Data,
				Object: tt.fields.Object,
			}
			if got := e.GetBody(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBody() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrorBag_Error(t *testing.T) {
	type fields struct {
		Err GptErrorResponse
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := ErrorBag{
				Err: tt.fields.Err,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrorBag_Temporary(t *testing.T) {
	type fields struct {
		Err GptErrorResponse
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := ErrorBag{
				Err: tt.fields.Err,
			}
			if got := e.Temporary(); got != tt.want {
				t.Errorf("Temporary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrorBag_Timeout(t *testing.T) {
	type fields struct {
		Err GptErrorResponse
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := ErrorBag{
				Err: tt.fields.Err,
			}
			if got := e.Timeout(); got != tt.want {
				t.Errorf("Timeout() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileRequest_attachResponse(t *testing.T) {
	type fields struct {
		File    os.File
		Purpose string
	}
	tests := []struct {
		name   string
		fields fields
		want   Response
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &FileRequest{
				File:    tt.fields.File,
				Purpose: tt.fields.Purpose,
			}
			if got := r.attachResponse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("attachResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileRequest_getRequestMeta(t *testing.T) {
	type fields struct {
		File    os.File
		Purpose string
	}
	type args struct {
		config RequestConfig
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		want1  string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &FileRequest{
				File:    tt.fields.File,
				Purpose: tt.fields.Purpose,
			}
			got, got1 := r.getRequestMeta(tt.args.config)
			if got != tt.want {
				t.Errorf("getRequestMeta() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getRequestMeta() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFileResponse_GetBody(t *testing.T) {
	type fields struct {
		File File
	}
	tests := []struct {
		name   string
		fields fields
		want   Response
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &FileResponse{
				File: tt.fields.File,
			}
			if got := r.GetBody(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBody() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilesRequest_attachResponse(t *testing.T) {
	tests := []struct {
		name string
		want Response
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &FilesRequest{}
			if got := r.attachResponse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("attachResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilesRequest_getRequestMeta(t *testing.T) {
	type args struct {
		config RequestConfig
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &FilesRequest{}
			got, got1 := r.getRequestMeta(tt.args.config)
			if got != tt.want {
				t.Errorf("getRequestMeta() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getRequestMeta() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFilesResponse_GetBody(t *testing.T) {
	type fields struct {
		Data   []File
		Object string
	}
	tests := []struct {
		name   string
		fields fields
		want   Response
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &FilesResponse{
				Data:   tt.fields.Data,
				Object: tt.fields.Object,
			}
			if got := r.GetBody(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBody() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchRequest_attachResponse(t *testing.T) {
	type fields struct {
		target         string
		Documents      []string
		Query          string
		File           string
		ReturnMetadata bool
		MaxRerank      int32
	}
	tests := []struct {
		name   string
		fields fields
		want   Response
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SearchRequest{
				target:         tt.fields.target,
				Documents:      tt.fields.Documents,
				Query:          tt.fields.Query,
				File:           tt.fields.File,
				ReturnMetadata: tt.fields.ReturnMetadata,
				MaxRerank:      tt.fields.MaxRerank,
			}
			if got := r.attachResponse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("attachResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchRequest_getRequestMeta(t *testing.T) {
	type fields struct {
		target         string
		Documents      []string
		Query          string
		File           string
		ReturnMetadata bool
		MaxRerank      int32
	}
	type args struct {
		config RequestConfig
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		want1  string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SearchRequest{
				target:         tt.fields.target,
				Documents:      tt.fields.Documents,
				Query:          tt.fields.Query,
				File:           tt.fields.File,
				ReturnMetadata: tt.fields.ReturnMetadata,
				MaxRerank:      tt.fields.MaxRerank,
			}
			got, got1 := r.getRequestMeta(tt.args.config)
			if got != tt.want {
				t.Errorf("getRequestMeta() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getRequestMeta() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSearchResponse_GetBody(t *testing.T) {
	type fields struct {
		Data   []SearchData
		Object string
	}
	tests := []struct {
		name   string
		fields fields
		want   Response
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SearchResponse{
				Data:   tt.fields.Data,
				Object: tt.fields.Object,
			}
			if got := r.GetBody(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBody() = %v, want %v", got, tt.want)
			}
		})
	}
}
