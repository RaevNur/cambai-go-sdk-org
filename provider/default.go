package provider

import (
	"context"
	"io"

	cambai "github.com/camb-ai/cambai-go-sdk"
	"github.com/camb-ai/cambai-go-sdk/client"
	"github.com/camb-ai/cambai-go-sdk/option"
)

// DefaultProvider implements TtsProvider using the official Camb.ai SDK client.
type DefaultProvider struct {
	client *client.Client
}

// NewDefaultProvider creates a new instance of DefaultProvider.
func NewDefaultProvider(apiKey string) *DefaultProvider {
	return &DefaultProvider{
		client: client.NewClient(option.WithAPIKey(apiKey)),
	}
}

// CreateTts sends a Text-to-Speech request to the Camb.ai API.
func (d *DefaultProvider) CreateTts(ctx context.Context, request *cambai.CreateTtsRequestPayload) (*cambai.CreateTtsOut, error) {
	return d.client.TextToSpeech.CreateTts(ctx, request)
}

// Tts sends a streaming Text-to-Speech request to the Camb.ai API.
func (d *DefaultProvider) Tts(ctx context.Context, request *cambai.CreateStreamTtsRequestPayload) (io.Reader, error) {
	return d.client.TextToSpeech.Tts(ctx, request)
}
