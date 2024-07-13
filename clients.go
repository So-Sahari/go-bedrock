package bedrock

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/bedrock"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
)

// CreateBedrockClient creates a bedrock client based on region
func CreateBedrockClient(ctx context.Context, region string) (*bedrock.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(region),
	)
	if err != nil {
		return nil, fmt.Errorf("error in create bedrock client: %w", err)
	}

	return bedrock.NewFromConfig(cfg), nil
}

// CreateBedrockruntimeClient creates a bedrockruntime client based on region
func CreateBedrockruntimeClient(ctx context.Context, region string) (*bedrockruntime.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(region),
	)
	if err != nil {
		return nil, fmt.Errorf("error in create bedrockruntime client: %w", err)
	}

	return bedrockruntime.NewFromConfig(cfg), nil
}
