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
		fmt.Println("Please set CAMB_API_KEY")
		return
	}

	c := client.NewClient(option.WithAPIKey(apiKey))

	fmt.Println("Generating Voice...")
	
	resp, err := c.TextToVoice.CreateTextToVoice(
		context.Background(),
		&cambai.CreateTextToVoiceRequestPayload{
			Text:             "This is a sample text for voice generation.",
			VoiceDescription: "A calm, deep male voice suitable for meditation.",
		},
	)
	if err != nil {
		panic(err)
	}

	if resp.TaskID == nil {
		panic("TaskID is nil")
	}
	taskID := *resp.TaskID
	fmt.Printf("Task Started: %s\n", taskID)

	for {
		time.Sleep(2 * time.Second)
		status, err := c.TextToVoice.GetTextToVoiceStatus(
			context.Background(),
			taskID,
			&cambai.GetTextToVoiceStatusTextToVoiceTaskIDGetRequest{},
		)
		if err != nil {
			fmt.Printf("Error polling: %v\n", err)
			continue
		}

		fmt.Printf("Status: %v\n", status.Status)

		if status.Status == cambai.TaskStatusSuccess {
			if status.RunID != nil {
				runID := *status.RunID
				res, err := c.TextToVoice.GetTextToVoiceResult(
					context.Background(),
					&runID,
					&cambai.GetTextToVoiceResultTextToVoiceResultRunIDGetRequest{},
				)
				if err != nil {
					panic(err)
				}
				fmt.Printf("Success! Generated Voice Previews: %v\n", res.Previews)
			}
			break
		} else if status.Status == cambai.TaskStatusError {
			fmt.Println("Task failed.")
			break
		}
	}
}
