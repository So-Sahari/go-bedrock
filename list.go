// Package bedrock contains aws logic
package bedrock

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/bedrock"
	"github.com/aws/aws-sdk-go-v2/service/bedrock/types"
)

// FoundationModel contains model fields
type FoundationModel struct {
	Name     string
	Provider string
	ID       string
	Modality string
}

// ClientAPI is used to interface with bedrock client
type ClientAPI interface {
	ListFoundationModels(ctx context.Context, params *bedrock.ListFoundationModelsInput, optFns ...func(*bedrock.Options)) (*bedrock.ListFoundationModelsOutput, error)
}

func ListModels(ctx context.Context, api ClientAPI) ([]FoundationModel, error) {
	var output []FoundationModel

	response, err := api.ListFoundationModels(ctx, &bedrock.ListFoundationModelsInput{
		ByInferenceType:  types.InferenceTypeOnDemand,
		ByOutputModality: types.ModelModalityText,
	})
	if err != nil {
		return output, err
	}

	filteredModels := []types.FoundationModelSummary{}
	for _, model := range response.ModelSummaries {
		for _, m := range SupportedModels {
			if strings.Contains(*model.ModelId, m) {
				filteredModels = append(filteredModels, model)
			}
		}
	}

	for _, sm := range filteredModels {
		output = append(output, FoundationModel{
			Name:     *sm.ModelName,
			Provider: *sm.ProviderName,
			ID:       *sm.ModelId,
			Modality: fmt.Sprintf("%v", sm.OutputModalities),
		})
	}
	return output, nil
}
