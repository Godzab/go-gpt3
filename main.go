package main

import (
	"fmt"
	"github.com/Godzab/go-gpt3/gpt3"
	"io/ioutil"
)

func main() {
	answersCall()
}


func answersCall(){
	examples := make([][]string, 1)
	data1 := []string{"What is human life expectancy in the United States?","78 years."}
	examples[0] = data1

	req := gpt3.AnswerRequest{
		Documents:       []string{"Puppy A is happy.","Puppy B is sad."},
		Question:        "which puppy is happy?",
		SearchModel:     gpt3.ADA,
		Model:           gpt3.CURIE,
		ExamplesContext: "In 2017, U.S. life expectancy was 78.6 years.",
		Examples:        examples,
		MaxTokens:       5,
		Stop:            []string{"\n", "<|endoftext|>"},
		Logprobs:        1,
		N:               1,
	}

	cl := gpt3.ApiClient{}
	cl.Setup(gpt3.ADA, gpt3.DAVINCI)

	response, err := cl.Call(&req)
	if err != nil {
		panic(err)
	}

	data := *response
	results, _ := data.(*gpt3.AnswerResponse)
	fmt.Println(results)
}

func completionCall(){
	query, err := ioutil.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}
	req := gpt3.CompletionRequest{
		Prompt:      string(query),
		MaxTokens:   64,
		TopP:        1,
		Temperature: 1,
		N: 1,
		BestOf: 1,
		FrequencyPenalty: 0.5,
		PresencePenalty: 0.6,
	}

	cl := gpt3.ApiClient{}
	cl.Setup(gpt3.DAVINCI_INSTRUCT_BETA, gpt3.DAVINCI)

	response, err := cl.Call(&req)
	if err != nil {
		panic(err)
	}

	data := *response
	results, _ := data.(*gpt3.CompletionResponse)

	for _,t  := range results.Choices{
		fmt.Println(t.Text)
	}
}