package main

import (
	"fmt"
	"os"
	"github.com/slack-go/slack"
)

func main() {

	os.Setenv("SLACK_BOT_TOKEN", "xoxb-3471069331845-3474048268690-XYut45U3yQxwrDuqUFLqIWjb")
	os.Setenv("CHANNEL_ID", "C03DXURF3AN")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"gofile.pdf"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("name: %s , URL:%s\n", file.Name, file.URL)
	}
}
