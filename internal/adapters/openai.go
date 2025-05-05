package adapters

import (
	"context"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

const (
	systemMessage = "Act like a drunk dad. Tell jokes. Response in Russian."
	defaultAnswer = "Sorry, I couldn't process your request. Please try again later."
)

type OpenAIClient struct {
	client *openai.Client
}

func NewOpenAIClient(api_key string) (*OpenAIClient, error) {
	client := openai.NewClient(
		option.WithAPIKey(api_key),
	)

	return &OpenAIClient{client: &client}, nil
}

func (client *OpenAIClient) GetResponse(ctx context.Context, query string) (string, error) {
	chatCompletion, err := client.client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage(systemMessage),
			openai.UserMessage(query),
		},
		Model: openai.ChatModelGPT4oMini,
	})
	if err != nil {
		return "", err
	}
	return chatCompletion.Choices[0].Message.Content, nil
}
