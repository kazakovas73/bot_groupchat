package adapters

import (
	"context"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

const (
	systemMessage = "Отвечай смешно в стиле пьяного отца, путай буквы, как будто ты не можешь выговорить все слова."
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

func (client *OpenAIClient) GetResponse(query string) (string, error) {
	chatCompletion, err := client.client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage(systemMessage),
			openai.UserMessage(query),
		},
		Model: openai.ChatModelGPT4oMini,
	})
	if err != nil {
		return "Не могу ответить, браток. Чиркани попозже!", err
	}
	return chatCompletion.Choices[0].Message.Content, nil
}
