
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/teamwork` Documentation

The `teamwork` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/teamwork"
```


### Client Initialization

```go
client := teamwork.NewTeamworkClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TeamworkClient.CreateTeamworkSendActivityNotification`

```go
ctx := context.TODO()
id := teamwork.NewUserID("userIdValue")

payload := teamwork.CreateTeamworkSendActivityNotificationRequest{
	// ...
}


read, err := client.CreateTeamworkSendActivityNotification(ctx, id, payload)
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
id := teamwork.NewUserID("userIdValue")

read, err := client.DeleteTeamwork(ctx, id)
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
id := teamwork.NewUserID("userIdValue")

read, err := client.GetTeamwork(ctx, id)
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
id := teamwork.NewUserID("userIdValue")

payload := teamwork.UserTeamwork{
	// ...
}


read, err := client.UpdateTeamwork(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
