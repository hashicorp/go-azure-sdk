
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/teamwork` Documentation

The `teamwork` SDK allows for interaction with Microsoft Graph `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/teamwork"
```


### Client Initialization

```go
client := teamwork.NewTeamworkClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TeamworkClient.DeleteTeamwork`

```go
ctx := context.TODO()


read, err := client.DeleteTeamwork(ctx, teamwork.DefaultDeleteTeamworkOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamworkClient.GetTeamwork`

```go
ctx := context.TODO()


read, err := client.GetTeamwork(ctx, teamwork.DefaultGetTeamworkOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamworkClient.SendTeamworkActivityNotification`

```go
ctx := context.TODO()

payload := teamwork.SendTeamworkActivityNotificationRequest{
	// ...
}


read, err := client.SendTeamworkActivityNotification(ctx, payload, teamwork.DefaultSendTeamworkActivityNotificationOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamworkClient.UpdateTeamwork`

```go
ctx := context.TODO()

payload := teamwork.UserTeamwork{
	// ...
}


read, err := client.UpdateTeamwork(ctx, payload, teamwork.DefaultUpdateTeamworkOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
