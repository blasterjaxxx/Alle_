package service

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
	"net/http"
	"sync"
)

type ChatService struct {
}

var (
	ChatServiceDoOnce sync.Once
	ChatServiceStruct ChatService
)

func InitializeChatService() ChatService {
	ChatServiceDoOnce.Do(func() {
		ChatServiceStruct = ChatService{}
	})
	return ChatServiceStruct
}

const openaiAPIKey = "YOUR_OPEN_AI_KEY"

func (ch ChatService) ChatWithUser(c *gin.Context) {
	var request struct {
		Message string `json:"message"`
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := ch.GenerateResponse(request.Message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": response})
}

func (ch ChatService) GenerateResponse(message string) (string, error) {
	//apiKey := os.Getenv("API_KEY")

	client := openai.NewClient(openaiAPIKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: message,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
