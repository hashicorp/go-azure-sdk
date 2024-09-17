
## `github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-10-01-preview/settings` Documentation

The `settings` SDK allows for interaction with Azure Resource Manager `securityinsights` (API Version `2022-10-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-10-01-preview/settings"
```


### Client Initialization

```go
client := settings.NewSettingsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SettingsClient.ProductSettingsDelete`

```go
ctx := context.TODO()
id := settings.NewSettingID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "settingValue")

read, err := client.ProductSettingsDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SettingsClient.ProductSettingsGet`

```go
ctx := context.TODO()
id := settings.NewSettingID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "settingValue")

read, err := client.ProductSettingsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SettingsClient.ProductSettingsList`

```go
ctx := context.TODO()
id := settings.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue")

read, err := client.ProductSettingsList(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SettingsClient.ProductSettingsUpdate`

```go
ctx := context.TODO()
id := settings.NewSettingID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "settingValue")

payload := settings.Settings{
	// ...
}


read, err := client.ProductSettingsUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
