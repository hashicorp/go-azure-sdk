
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/userteamwork` Documentation

The `userteamwork` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/userteamwork"
```


### Client Initialization

```go
client := userteamwork.NewUserTeamworkClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserTeamworkClient.CreateUserByIdTeamworkSendActivityNotification`

```go
ctx := context.TODO()
id := userteamwork.NewUserID("userIdValue")

payload := userteamwork.CreateUserByIdTeamworkSendActivityNotificationRequest{
	// ...
}


read, err := client.CreateUserByIdTeamworkSendActivityNotification(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserTeamworkClient.DeleteUserByIdTeamwork`

```go
ctx := context.TODO()
id := userteamwork.NewUserID("userIdValue")

read, err := client.DeleteUserByIdTeamwork(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserTeamworkClient.GetUserByIdTeamwork`

```go
ctx := context.TODO()
id := userteamwork.NewUserID("userIdValue")

read, err := client.GetUserByIdTeamwork(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserTeamworkClient.UpdateUserByIdTeamwork`

```go
ctx := context.TODO()
id := userteamwork.NewUserID("userIdValue")

payload := userteamwork.UserTeamwork{
	// ...
}


read, err := client.UpdateUserByIdTeamwork(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
