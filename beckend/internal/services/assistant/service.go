package assistant

import (
	"context"
	"errors"
	"time"

	"github.com/sashabaranov/go-openai"

	"github.com/tahmazidik/llm-assistant-web-lab/beckend/internal/models"
	dialogssvc "github.com/tahmazidik/llm-assistant-web-lab/beckend/internal/services/dialogs"
)

var ErrNoApiKey = errors.New("openai api key is empty")

type Service interface {
	Reply(ctx context.Context, dialogID models.DialogID) (string, error)
}

type service struct {
	dialogs     dialogssvc.Service
	client      *openai.Client
	model       string
	maxHistory  int
	systemIntro string
}

type Config struct {
	APIKey      string
	Model       string
	MaxHistory  int
	SystemIntro string
}

func New(dialogs dialogssvc.Service, cfg Config) (Service, error) {
	if cfg.APIKey == "" {
		return nil, ErrNoApiKey
	}
	maxHis := cfg.MaxHistory
	if maxHis <= 0 {
		maxHis = 20
	}

	intro := cfg.SystemIntro
	if intro == "" {
		intro = "You are a helpful assistant. Answer clearly and concisely."
	}

	model := cfg.Model
	if model == "" {
		model = openai.GPT4oMini
	}

	return &service{
		dialogs:     dialogs,
		client:      openai.NewClient(cfg.APIKey),
		model:       model,
		maxHistory:  maxHis,
		systemIntro: intro,
	}, nil
}

func (serv *service) Reply(ctx context.Context, dialogID models.DialogID) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	history, err := serv.dialogs.ListMessages(ctx, dialogID)

	if err != nil {
		return "", err
	}

	//Берем последние N сообщение
	if len(history) > serv.maxHistory {
		history = history[len(history)-serv.maxHistory:]
	}

	msgs := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: serv.systemIntro,
		},
	}

	for _, m := range history {
		switch m.Sender {
		case models.SenderUser:
			msgs = append(msgs, openai.ChatCompletionMessage{
				Role:    openai.ChatMessageRoleUser,
				Content: m.Content,
			})
		case models.SenderAssistant:
			msgs = append(msgs, openai.ChatCompletionMessage{
				Role:    openai.ChatMessageRoleAssistant,
				Content: m.Content,
			})
		}
	}

	resp, err := serv.client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:    serv.model,
		Messages: msgs,
	})

	if err != nil {
		return "", err
	}

	if len(resp.Choices) == 0 {
		return "", errors.New("openai: empty choices")
	}

	return resp.Choices[0].Message.Content, nil
}
