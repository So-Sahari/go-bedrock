package bedrock

// SupportedModels contains all supported models implemented
var SupportedModels = []string{"sonnet", "anthropic", "cohere", "mistral"}

// AWSModelConfig contains all settings for AWS Models
type AWSModelConfig struct {
	ModelID     string  // ModelID is the ID of the model to invoke
	Temperature float64 // Temperature is part of model settings (Anthropic, Cohere)
	TopP        float64 // TopP is part of model settings (Anthropic, Cohere)
	TopK        int     // TopK is part of model settings (Anthropic, Cohere)
	MaxTokens   int     // MaxTokens is part of model settings (Anthropic, Cohere)
}

// NewModel is a factory for AWS Models
func NewModel(modelID string, temp, topP float64, topK, tokens int) AWSModelConfig {
	return AWSModelConfig{
		ModelID:     modelID,
		Temperature: temp,
		TopP:        topP,
		TopK:        topK,
		MaxTokens:   tokens,
	}
}
