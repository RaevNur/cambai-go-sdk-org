# Camb.ai Go SDK

The official Go SDK for interacting with [Camb.ai](https://camb.ai)'s powerful voice and audio generation APIs. Create expressive speech, unique voices, and rich soundscapes with idiomatic Go.

## ✨ Features

- **Dubbing**: End-to-end video dubbing and translation.
- **Expressive Text-to-Speech**: High-fidelity speech synthesis with standard and custom voices.
- **Generative Voices**: Create entirely new voices from text descriptions.
- **Voice Cloning**: Clone voices from audio samples.
- **Soundscapes**: Generate ambient audio and sound effects.

## 📦 Installation

```bash
go get github.com/camb-ai/cambai-go-sdk
```

## 🔑 Authentication

Initialize the client with your API key. You can pass it directly or load it from an environment variable.

```go
import (
    "os"
    "github.com/camb-ai/cambai-go-sdk/client"
    "github.com/camb-ai/cambai-go-sdk/option"
)

func main() {
    c := client.NewClient(
        option.WithAPIKey(os.Getenv("CAMB_API_KEY")),
        // option.WithBaseURL("..."), // Optional: Override Base URL
    )
}
```

## 🚀 Examples

### 1. Text-to-Speech (TTS)

Generate audio from text. This example creates a task and polls for completion.

```go
package main

import (
    "context"
    "fmt"
    "os"
    "strconv"
    "time"

    "github.com/camb-ai/cambai-go-sdk"
    "github.com/camb-ai/cambai-go-sdk/client"
    "github.com/camb-ai/cambai-go-sdk/option"
)

func main() {
    c := client.NewClient(option.WithAPIKey(os.Getenv("CAMB_API_KEY")))

    // 1. Initiate TTS Task
    
    // Fetch Language ID for English (United States)
    sourceLangs, _ := c.Languages.GetSourceLanguages(context.Background(), &cambai.GetSourceLanguagesSourceLanguagesGetRequest{})
    var englishID int
    for _, l := range sourceLangs {
        if l.ShortName == "en-us" {
            englishID = l.ID
            break
        }
    }
    if englishID == 0 { englishID = 1 }

    resp, err := c.TextToSpeech.CreateTts(
        context.Background(),
        &cambai.CreateTtsRequestPayload{
            Text:     "Hello from Camb AI Go SDK!",
            VoiceID:  20303, // Standard Voice
            Language: englishID,
        },
    )
    if err != nil {
        panic(err)
    }

    // 2. Poll for Completion
    runID, _ := strconv.Atoi(resp.TaskID) // TaskID is returned as string
    fmt.Printf("TTS Started. Run ID: %d\n", runID)

    for {
        time.Sleep(1 * time.Second)
        
        status, err := c.TextToSpeech.GetTtsRunInfo(
            context.Background(), 
            &runID, 
            &cambai.GetTtsRunInfoTtsResultRunIDGetRequest{},
        )
        if err != nil {
            fmt.Printf("Polling error: %v\n", err)
            continue
        }

        // Check if we have a URL
        if status.GetTtsResultOutFileURL != nil {
            fmt.Printf("Success! Audio URL: %s\n", status.GetTtsResultOutFileURL.OutputURL)
            break
        }
        fmt.Println("Processing...")
    }
}
```

### 2. End-to-End Dubbing

Dub a video from one language to another.

```go
import (
    "context"
    "fmt"
    "time"
    "github.com/camb-ai/cambai-go-sdk"
    "github.com/camb-ai/cambai-go-sdk/client"
)

func main() {
    c := client.NewClient()

    // 1. Create Dubbing Task
    videoURL := "https://example.com/video.mp4"
    
    // Logic to find language IDs omitted for brevity, see examples/dubbing/main.go
    
    resp, _ := c.Dub.EndToEndDubbing(
        context.Background(),
        &cambai.EndToEndDubbingRequestPayload{
            VideoURL:        videoURL,
            SourceLanguage:  1, // en-us
            TargetLanguages: []cambai.Languages{2}, // fr-fr
        },
    )

    taskID := *resp.TaskID
    fmt.Printf("Dubbing started. Task ID: %s\n", taskID)

    // 2. Poll Status
    for {
        time.Sleep(5 * time.Second)
        statusResp, _ := c.Dub.GetEndToEndDubbingStatus(
            context.Background(), 
            taskID, 
            &cambai.GetEndToEndDubbingStatusDubTaskIDGetRequest{},
        )
        
        fmt.Printf("Status: %s\n", statusResp.Status)

        if statusResp.Status == cambai.TaskStatusSuccess {
            // Get Result
            runID := *statusResp.RunID
            result, _ := c.Dub.GetDubbedRunInfo(
                context.Background(), 
                &runID, 
                &cambai.GetDubbedRunInfoDubResultRunIDGetRequest{},
            )
            if result.DubbingResult != nil && result.DubbingResult.VideoURL != nil {
                fmt.Printf("Dubbed Video: %s\n", *result.DubbingResult.VideoURL)
            }
            break
        } else if statusResp.Status == cambai.TaskStatusError {
            fmt.Println("Dubbing failed.")
            break
        }
    }
}
```

### 3. Text-to-Voice

Generate a unique new voice from a description.

```go
resp, _ := c.TextToVoice.CreateTextToVoice(
    context.Background(),
    &cambai.CreateTextToVoiceRequestPayload{
        Text:             "This is a test sentence for the new voice.",
        VoiceDescription: "A deep, resonant voice suitable for narration.",
    },
)
// Poll resp.TaskID using c.TextToVoice.GetTextToVoiceStatus(...)
```

## ⚙️ Advanced

### List Available Voices

```go
voices, _ := c.VoiceCloning.ListVoices(
    context.Background(), 
    &cambai.ListVoicesListVoicesGetRequest{},
)

for _, v := range voices {
    fmt.Printf("ID: %d, Name: %s\n", v.ID, v.VoiceName)
}
```

## License

MIT
