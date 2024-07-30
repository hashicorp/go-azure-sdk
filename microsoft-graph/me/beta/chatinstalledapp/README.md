
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/chatinstalledapp` Documentation

The `chatinstalledapp` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/chatinstalledapp"
```


### Client Initialization

```go
client := chatinstalledapp.NewChatInstalledAppClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ChatInstalledAppClient.CreateChatInstalledApp`

```go
ctx := context.TODO()
id := chatinstalledapp.NewMeChatID("chatIdValue")

payload := chatinstalledapp.TeamsAppInstallation{
	// ...
}


read, err := client.CreateChatInstalledApp(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChatInstalledAppClient.CreateChatInstalledAppUpgrade`

```go
ctx := context.TODO()
id := chatinstalledapp.NewMeChatIdInstalledAppID("chatIdValue", "teamsAppInstallationIdValue")

payload := chatinstalledapp.CreateChatInstalledAppUpgradeRequest{
	// ...
}


read, err := client.CreateChatInstalledAppUpgrade(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChatInstalledAppClient.DeleteChatInstalledApp`

```go
ctx := context.TODO()
id := chatinstalledapp.NewMeChatIdInstalledAppID("chatIdValue", "teamsAppInstallationIdValue")

read, err := client.DeleteChatInstalledApp(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChatInstalledAppClient.GetChatInstalledApp`

```go
ctx := context.TODO()
id := chatinstalledapp.NewMeChatIdInstalledAppID("chatIdValue", "teamsAppInstallationIdValue")

read, err := client.GetChatInstalledApp(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChatInstalledAppClient.GetChatInstalledAppCount`

```go
ctx := context.TODO()
id := chatinstalledapp.NewMeChatID("chatIdValue")

read, err := client.GetChatInstalledAppCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChatInstalledAppClient.ListChatInstalledApps`

```go
ctx := context.TODO()
id := chatinstalledapp.NewMeChatID("chatIdValue")

// alternatively `client.ListChatInstalledApps(ctx, id)` can be used to do batched pagination
items, err := client.ListChatInstalledAppsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ChatInstalledAppClient.UpdateChatInstalledApp`

```go
ctx := context.TODO()
id := chatinstalledapp.NewMeChatIdInstalledAppID("chatIdValue", "teamsAppInstallationIdValue")

payload := chatinstalledapp.TeamsAppInstallation{
	// ...
}


read, err := client.UpdateChatInstalledApp(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
