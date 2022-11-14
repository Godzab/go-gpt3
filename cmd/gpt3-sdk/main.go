//Package gpt3 provides access to the GPT3 completions Api
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

package gpt3_sdk
