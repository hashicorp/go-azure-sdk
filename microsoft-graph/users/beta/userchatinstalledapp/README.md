
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/userchatinstalledapp` Documentation

The `userchatinstalledapp` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/userchatinstalledapp"
```


### Client Initialization

```go
client := userchatinstalledapp.NewUserChatInstalledAppClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserChatInstalledAppClient.CreateUserByIdChatByIdInstalledApp`

```go
ctx := context.TODO()
id := userchatinstalledapp.NewUserChatID("userIdValue", "chatIdValue")

payload := userchatinstalledapp.TeamsAppInstallation{
	// ...
}


read, err := client.CreateUserByIdChatByIdInstalledApp(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserChatInstalledAppClient.CreateUserByIdChatByIdInstalledAppByIdUpgrade`

```go
ctx := context.TODO()
id := userchatinstalledapp.NewUserChatInstalledAppID("userIdValue", "chatIdValue", "teamsAppInstallationIdValue")

payload := userchatinstalledapp.CreateUserByIdChatByIdInstalledAppByIdUpgradeRequest{
	// ...
}


read, err := client.CreateUserByIdChatByIdInstalledAppByIdUpgrade(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserChatInstalledAppClient.DeleteUserByIdChatByIdInstalledAppById`

```go
ctx := context.TODO()
id := userchatinstalledapp.NewUserChatInstalledAppID("userIdValue", "chatIdValue", "teamsAppInstallationIdValue")

read, err := client.DeleteUserByIdChatByIdInstalledAppById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserChatInstalledAppClient.GetUserByIdChatByIdInstalledAppById`

```go
ctx := context.TODO()
id := userchatinstalledapp.NewUserChatInstalledAppID("userIdValue", "chatIdValue", "teamsAppInstallationIdValue")

read, err := client.GetUserByIdChatByIdInstalledAppById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserChatInstalledAppClient.GetUserByIdChatByIdInstalledAppCount`

```go
ctx := context.TODO()
id := userchatinstalledapp.NewUserChatID("userIdValue", "chatIdValue")

read, err := client.GetUserByIdChatByIdInstalledAppCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserChatInstalledAppClient.ListUserByIdChatByIdInstalledApps`

```go
ctx := context.TODO()
id := userchatinstalledapp.NewUserChatID("userIdValue", "chatIdValue")

// alternatively `client.ListUserByIdChatByIdInstalledApps(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdChatByIdInstalledAppsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserChatInstalledAppClient.UpdateUserByIdChatByIdInstalledAppById`

```go
ctx := context.TODO()
id := userchatinstalledapp.NewUserChatInstalledAppID("userIdValue", "chatIdValue", "teamsAppInstallationIdValue")

payload := userchatinstalledapp.TeamsAppInstallation{
	// ...
}


read, err := client.UpdateUserByIdChatByIdInstalledAppById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
