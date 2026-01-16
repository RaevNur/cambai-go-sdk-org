package provider

import (
	"context"

	cambai "github.com/camb-ai/cambai-go-sdk"
)

// TtsProvider defines the interface for Text-to-Speech operations.
// This allows for swapping the default Camb.ai implementation with custom providers (e.g., Baseten, Vertex AI).
type TtsProvider interface {
	CreateTts(ctx context.Context, request *cambai.CreateTtsRequestPayload) (*cambai.CreateTtsOut, error)
}
