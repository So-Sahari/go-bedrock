package bedrock

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/So-Sahari/go-bedrock/models"
)

func (m *AWSModelConfig) constructPayload(message string) ([]byte, error) {
	switch {
	case strings.Contains(m.ModelID, "sonnet"), strings.Contains(m.ModelID, "haiku"):
		body := models.ClaudeMessagesInput{
			AnthropicVersion: "bedrock-2023-05-31",
			Messages: []models.ClaudeMessage{
				{
					Role: "user",
					Content: []models.ClaudeContent{
						{
							Type: "text",
							Text: message,
						},
					},
				},
			},
			MaxTokens:   m.MaxTokens,
			Temperature: m.Temperature,
			TopP:        m.TopP,
			TopK:        m.TopK,
		}

		payload, err := json.Marshal(body)
		if err != nil {
			return []byte{}, err
		}
		return payload, nil

	case strings.Contains(m.ModelID, "anthropic"):
		body := models.ClaudeModelInputs{
			Prompt:            fmt.Sprintf("\n\nHuman: %s\n\nAssistant:", message),
			MaxTokensToSample: m.MaxTokens,
			Temperature:       m.Temperature,
			TopP:              m.TopP,
			TopK:              m.TopK,
		}

		payload, err := json.Marshal(body)
		if err != nil {
			return []byte{}, err
		}
		return payload, nil

	case strings.Contains(m.ModelID, "cohere"):
		body := models.CommandModelInput{
			Prompt:            message,
			MaxTokensToSample: m.MaxTokens,
			Temperature:       m.Temperature,
			TopP:              m.TopP,
			TopK:              m.TopK,
			StopSequences:     []string{`""`},
			ReturnLiklihoods:  "NONE",
			NumGenerations:    1,
		}

		payload, err := json.Marshal(body)
		if err != nil {
			return []byte{}, err
		}
		return payload, nil

	case strings.Contains(m.ModelID, "mistral"):
		// handle the default being higher than the model allows
		if m.TopK > 200 {
			m.TopK = 200
		}

		body := models.MistralRequest{
			Prompt:      message,
			MaxTokens:   m.MaxTokens,
			Temperature: m.Temperature,
			TopP:        m.TopP,
			TopK:        m.TopK,
		}

		payload, err := json.Marshal(body)
		if err != nil {
			return []byte{}, err
		}
		return payload, nil

	default:
		fmt.Println("ModelID not provided or unknown")
	}

	return []byte{}, fmt.Errorf("unable to construct payload for model: %s", m.ModelID)
}
