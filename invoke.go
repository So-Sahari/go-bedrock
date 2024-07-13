// Package bedrock contains aws logic
package bedrock

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
)

// ClientRuntimeAPI interface is for support mocking of API calls
type ClientRuntimeAPI interface {
	InvokeModelWithResponseStream(ctx context.Context, params *bedrockruntime.InvokeModelWithResponseStreamInput, optFns ...func(*bedrockruntime.Options)) (*bedrockruntime.InvokeModelWithResponseStreamOutput, error)
}

// StreamingOutputHandler used for processing streaming output
type StreamingOutputHandler func(ctx context.Context, part []byte) error

// InvokeModel runs prompt with settings with InvokeModelWithResponseStream
func (m *AWSModelConfig) InvokeModel(ctx context.Context, api ClientRuntimeAPI, message string) (string, error) {
	contentTypeVar := "application/json"

	payload, err := m.constructPayload(message)
	if err != nil {
		return "", err
	}

	// invoke model
	output, err := api.InvokeModelWithResponseStream(ctx, &bedrockruntime.InvokeModelWithResponseStreamInput{
		ContentType: &contentTypeVar,
		ModelId:     &m.ModelID,
		Body:        payload,
	})
	if err != nil {
		return "", err
	}

	// handle stream response chunks by model
	response, err := m.processStreamingOutput(output, func(ctx context.Context, part []byte) error {
		fmt.Print(string(part))
		return nil
	})
	if err != nil {
		return "", err
	}

	return response, nil
}
