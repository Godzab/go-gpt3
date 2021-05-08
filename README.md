# go-gpt3
Package gpt3 provides access to the the GPT3 completions Api
along with new beta APIs for classification, enhanced search, and question answering.

Additionally, support for newer Beta Engines is added broadening the scope in which you can query the APIs from your Golang application/packages.


The underlying structure is defined along a request / response interface pattern with a
singular call to the client.
The request is initialised as per required parameters an example being:


```go
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
```
Full examples can be found and run at `go run internal/examples/main.go`

You can alternate between the different engines through the call to `setup` on the client to switch engines on execution.

```go
//Initialises a client
cl := gpt3.ApiClient{}
//Stipulate engines to use.
cl.Setup(gpt3.ADA, gpt3.DAVINCI) 
//Execute request
response, err := cl.Call(&req)
```

***

## Support
- List Engines API
- Completion API
- Document Search API
- Files List API
- Answers API
- Classification API
- Enhanced Search API