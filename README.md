
# Slack-Bot-API 
A bot-api written in GO lang to upload files from remote to the slack chat channel  

## Demo
You should be admin of the slack chat channel  
Add the bot by inviting it.

![local](https://github.com/prashant54singh/GO_Projects/blob/main/slack/Screenshot%202022-08-28%20215807.jpg?raw=true)  
Now Your file should be updated like this   

![hello](https://github.com/prashant54singh/GO_Projects/blob/main/slack/Screenshot%202022-08-28%20215904.jpg?raw=true)

## Deployment  

To deploy this project install GO lang

Go to the src of the GO file location and open cmd
```
go mod init github.com/gorilla/mux
go get github.com/gorilla/mux
```
run these to install gorilla  
then run a slack package installer to talk to slack network
```
"github.com/slack-go/slack"
```
run...
```bash
-go build main.go
```
```
-go run main.go
```


