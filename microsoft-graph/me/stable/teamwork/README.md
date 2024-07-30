
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/teamwork` Documentation

The `teamwork` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/teamwork"
```


### Client Initialization

```go
client := teamwork.NewTeamworkClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TeamworkClient.CreateTeamworkSendActivityNotification`

```go
ctx := context.TODO()

payload := teamwork.CreateTeamworkSendActivityNotificationRequest{
	// ...
}


read, err := client.CreateTeamworkSendActivityNotification(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamworkClient.DeleteTeamwork`

```go
ctx := context.TODO()


read, err := client.DeleteTeamwork(ctx)
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


read, err := client.GetTeamwork(ctx)
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


read, err := client.UpdateTeamwork(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
