
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/meteamwork` Documentation

The `meteamwork` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/meteamwork"
```


### Client Initialization

```go
client := meteamwork.NewMeTeamworkClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeTeamworkClient.CreateMeTeamworkSendActivityNotification`

```go
ctx := context.TODO()

payload := meteamwork.CreateMeTeamworkSendActivityNotificationRequest{
	// ...
}


read, err := client.CreateMeTeamworkSendActivityNotification(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeTeamworkClient.DeleteMeTeamwork`

```go
ctx := context.TODO()


read, err := client.DeleteMeTeamwork(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeTeamworkClient.GetMeTeamwork`

```go
ctx := context.TODO()


read, err := client.GetMeTeamwork(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeTeamworkClient.UpdateMeTeamwork`

```go
ctx := context.TODO()

payload := meteamwork.UserTeamwork{
	// ...
}


read, err := client.UpdateMeTeamwork(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
