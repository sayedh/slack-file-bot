# Slack File Upload Bot
This is a project built with Golang. A simple api where one can upload a file to their slack channel.


## Technologies Used
* Golang
* Slack


## Dependencies

* Install Go - https://go.dev/doc/install


## Executing program

* Download the repository to your computer and go to project file
```
git clone https://github.com/sayedh/slack-file-bot
cd slack-file-bot
go mod init github.com/sayedh/slack-file-bot
go get "github.com/slack-go/slack"
go mod tidy
code .
```
* Add your Bot Token and Channel ID into the main.go code
```
os.Setenv("SLACK_BOT_TOKEN", "<insert here>")
os.Setenv("CHANNEL_ID", "<insert here>")
```
* Add your Bot to your channel in slack

* Run the code to upload the file
```
go run main.go
```
