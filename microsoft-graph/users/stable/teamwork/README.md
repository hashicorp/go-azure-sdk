
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/teamwork` Documentation

The `teamwork` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/teamwork"
```


### Client Initialization

```go
client := teamwork.NewTeamworkClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TeamworkClient.DeleteTeamwork`

```go
ctx := context.TODO()
id := teamwork.NewUserID("userIdValue")

read, err := client.DeleteTeamwork(ctx, id, teamwork.DefaultDeleteTeamworkOperationOptions())
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

read, err := client.GetTeamwork(ctx, id, teamwork.DefaultGetTeamworkOperationOptions())
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
id := teamwork.NewUserID("userIdValue")

payload := teamwork.SendTeamworkActivityNotificationRequest{
	// ...
}


read, err := client.SendTeamworkActivityNotification(ctx, id, payload)
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
