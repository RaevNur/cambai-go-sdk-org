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
	apiKey := os.Getenv("CAMB_API_KEY")
	if apiKey == "" {
		fmt.Println("Please set CAMB_API_KEY environment variable")
		return
	}

	c := client.NewClient(
		option.WithAPIKey(apiKey),
	)

	fmt.Println("Sending TTS request...")
	resp, err := c.TextToSpeech.CreateTts(
		context.Background(),
		&cambai.CreateTtsRequestPayload{
			Text:     "Hello from Go SDK!",
			VoiceID:  20303, // Standard voice
			Language: cambai.LanguagesEnUs,
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("TTS Task Created! TaskID: %s\n", resp.TaskID)

	// Parse Task ID (which is returned as string from CreateTts but required as int for GetRunInfo)
	runID, err := strconv.Atoi(resp.TaskID)
	if err != nil {
		panic(fmt.Errorf("failed to parse run ID: %v", err))
	}

	fmt.Printf("Polling for Run ID: %d...\n", runID)

	for {
		time.Sleep(2 * time.Second)
		status, err := c.TextToSpeech.GetTtsRunInfo(context.Background(), &runID, &cambai.GetTtsRunInfoTtsResultRunIDGetRequest{})
		if err != nil {
			fmt.Printf("Error polling: %v\n", err)
			continue
		}

		fmt.Printf("Status Response Received.\n")
		// The status response is a union type. We check fields.
		if status.GetTtsResultOutFileURL != nil {
			fmt.Printf("Success! Audio URL: %s\n", status.GetTtsResultOutFileURL.OutputURL)
			break
		}

		if status.String != "" {
			fmt.Printf("Status: %s\n", status.String)
			if status.String == "SUCCESS" {
				// Wait, if it is success, we expect FileURL.
				// Maybe the union type handling depends on actual JSON structure.
			}
		}
	}
}
