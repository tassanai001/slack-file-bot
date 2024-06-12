package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {

	os.Setenv("SLACK_BOT_TOKEN", "<SLACK_BOT_TOKEN>")
	os.Setenv("CHANNEL_ID", "<CHANNEL_ID>")

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArray := os.Getenv("CHANNEL_ID")
	filArr := []string{"TestPDFfile.pdf"}

	for i := 0; i < len(filArr); i++ {

		fi, err := os.Stat("TestPDFfile.pdf")
		if err != nil {
			return
		}
		// get the size
		size := fi.Size()

		params := slack.UploadFileV2Parameters{
			Channel:  channelArray,
			File:     filArr[i],
			Filename: filArr[i],
			FileSize: int(size),
		}

		file, err := api.UploadFileV2(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Name: %s, URL: %s\n", file.ID, file.Title)
	}
}
