# Camb.ai Go SDK

The official Go SDK for [Camb.ai](https://camb.ai).

## Installation

```bash
go get github.com/camb-ai/cambai-go-sdk
```

## Usage

### Client Initialization

```go
import (
    "context"
    "fmt"
    "os"

    "github.com/camb-ai/cambai-go-sdk/client"
    "github.com/camb-ai/cambai-go-sdk/option"
)

func main() {
    client := client.NewClient(
        option.WithAPIKey(os.Getenv("CAMB_API_KEY")),
    )
}
```

### Text-to-Speech (TTS)

```go
import (
    "context"
    "fmt"
    "os"
    "github.com/camb-ai/cambai-go-sdk/client"
    "github.com/camb-ai/cambai-go-sdk/option"
    sdk "github.com/camb-ai/cambai-go-sdk"
)

func main() {
    c := client.NewClient(
        option.WithAPIKey(os.Getenv("CAMB_API_KEY")),
    )

    // Create TTS task
    resp, err := c.TextToSpeech.CreateTts(
        context.TODO(),
        &sdk.CreateTtsRequestPayload{
            Text: "Hello world",
            VoiceID: 20303,
            Language: sdk.LanguagesEnUs,
        },
    )
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("Task ID: %s\n", resp.TaskID)
}
```

## Requirements

- Go 1.18+

## License

MIT
