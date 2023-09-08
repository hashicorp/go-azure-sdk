
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mechatinstalledapp` Documentation

The `mechatinstalledapp` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mechatinstalledapp"
```


### Client Initialization

```go
client := mechatinstalledapp.NewMeChatInstalledAppClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeChatInstalledAppClient.CreateMeChatByIdInstalledApp`

```go
ctx := context.TODO()
id := mechatinstalledapp.NewMeChatID("chatIdValue")

payload := mechatinstalledapp.TeamsAppInstallation{
	// ...
}


read, err := client.CreateMeChatByIdInstalledApp(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeChatInstalledAppClient.CreateMeChatByIdInstalledAppByIdUpgrade`

```go
ctx := context.TODO()
id := mechatinstalledapp.NewMeChatInstalledAppID("chatIdValue", "teamsAppInstallationIdValue")

payload := mechatinstalledapp.CreateMeChatByIdInstalledAppByIdUpgradeRequest{
	// ...
}


read, err := client.CreateMeChatByIdInstalledAppByIdUpgrade(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeChatInstalledAppClient.DeleteMeChatByIdInstalledAppById`

```go
ctx := context.TODO()
id := mechatinstalledapp.NewMeChatInstalledAppID("chatIdValue", "teamsAppInstallationIdValue")

read, err := client.DeleteMeChatByIdInstalledAppById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeChatInstalledAppClient.GetMeChatByIdInstalledAppById`

```go
ctx := context.TODO()
id := mechatinstalledapp.NewMeChatInstalledAppID("chatIdValue", "teamsAppInstallationIdValue")

read, err := client.GetMeChatByIdInstalledAppById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeChatInstalledAppClient.GetMeChatByIdInstalledAppCount`

```go
ctx := context.TODO()
id := mechatinstalledapp.NewMeChatID("chatIdValue")

read, err := client.GetMeChatByIdInstalledAppCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeChatInstalledAppClient.ListMeChatByIdInstalledApps`

```go
ctx := context.TODO()
id := mechatinstalledapp.NewMeChatID("chatIdValue")

// alternatively `client.ListMeChatByIdInstalledApps(ctx, id)` can be used to do batched pagination
items, err := client.ListMeChatByIdInstalledAppsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeChatInstalledAppClient.UpdateMeChatByIdInstalledAppById`

```go
ctx := context.TODO()
id := mechatinstalledapp.NewMeChatInstalledAppID("chatIdValue", "teamsAppInstallationIdValue")

payload := mechatinstalledapp.TeamsAppInstallation{
	// ...
}


read, err := client.UpdateMeChatByIdInstalledAppById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
