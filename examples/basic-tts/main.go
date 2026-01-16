package main

import (
	"context"
	"fmt"
	"os"
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
		context.TODO(),
		&cambai.CreateTtsRequestPayload{
			Text:     "Hello from Go SDK!",
			VoiceID:  20303, // Standard voice
			Language: cambai.LanguagesEnUs,
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("TTS Task Created! ID: %s\n", resp.TaskID)
	
	// Parse Task ID
	var runID int
	fmt.Sscanf(resp.TaskID, "%d", &runID)

	fmt.Printf("Polling for Run ID: %d...\n", runID)

	for {
		time.Sleep(1 * time.Second)
		status, err := c.TextToSpeech.GetTtsRunInfo(context.TODO(), &runID, &cambai.GetTtsRunInfoTtsResultRunIDGetRequest{})
		if err != nil {
			fmt.Printf("Error polling: %v\n", err)
			continue
		}
		
		fmt.Printf("Status: %v\n", status.String)
		if status.GetTtsResultOutFileURL != nil {
			fmt.Printf("Success! URL: %s\n", status.GetTtsResultOutFileURL.OutputURL)
			break
		}
	}
}
