# Slack File Upload Bot
This is a fullstack project built with HTML/CSS frontend and Golang backend. A simple To-Do app where one can create, edit and delete tasks.


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
os.Setenv("SLACK_BOT_TOKEN", "xoxb-3471069331845-3474048268690-XYut45U3yQxwrDuqUFLqIWjb")
os.Setenv("CHANNEL_ID", "C03DXURF3AN")
```
* Add your Bot to your channel in slack

* Run the code to upload the file
```
go run main.go
```
