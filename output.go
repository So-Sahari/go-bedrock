package bedrock

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime/types"

	"github.com/So-Sahari/go-bedrock/models"
)

func (m *AWSModelConfig) processStreamingOutput(output *bedrockruntime.InvokeModelWithResponseStreamOutput, handler StreamingOutputHandler) (string, error) {
	var combinedResult string

	for event := range output.GetStream().Events() {
		switch v := event.(type) {
		case *types.ResponseStreamMemberChunk:
			// nested switch case for stream outputs. ugh
			switch {
			case strings.Contains(m.ModelID, "sonnet"), strings.Contains(m.ModelID, "haiku"):
				var pr PartialResponse
				err := json.NewDecoder(bytes.NewReader(v.Value.Bytes)).Decode(&pr)
				if err != nil {
					return combinedResult, err
				}

				if pr.Type == partialResponseTypeContentBlockDelta {
					handler(context.Background(), []byte(pr.Delta.Text))
					combinedResult += pr.Delta.Text
				}
			case strings.Contains(m.ModelID, "anthropic"):
				var resp models.ClaudeModelOutputs
				if err := json.NewDecoder(bytes.NewReader(v.Value.Bytes)).Decode(&resp); err != nil {
					return combinedResult, err
				}

				handler(context.Background(), []byte(resp.Completion))
				combinedResult += resp.Completion
			case strings.Contains(m.ModelID, "cohere"):
				var resp models.CommandModelOutput
				if err := json.NewDecoder(bytes.NewReader(v.Value.Bytes)).Decode(&resp); err != nil {
					return combinedResult, err
				}

				handler(context.Background(), []byte(resp.Generations[0].Text))
				combinedResult += resp.Generations[0].Text
			case strings.Contains(m.ModelID, "mistral"):
				var resp models.MistralResponse
				if err := json.Unmarshal([]byte(string(v.Value.Bytes)), &resp); err != nil {
					return combinedResult, err
				}

				handler(context.Background(), []byte(resp.Outputs[0].Text))
				combinedResult += resp.Outputs[0].Text
			default:
				fmt.Println("Unable to determine AWS Model")
			}

		case *types.UnknownUnionMember:
			fmt.Println("unknown tag:", v.Tag)

		default:
			fmt.Println("union is nil or unknown type")
		}
	}
	fmt.Println()

	return combinedResult, nil
}
