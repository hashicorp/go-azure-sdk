
## `github.com/hashicorp/go-azure-sdk/resource-manager/security/2017-08-01-preview/workspacesettings` Documentation

The `workspacesettings` SDK allows for interaction with the Azure Resource Manager Service `security` (API Version `2017-08-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/security/2017-08-01-preview/workspacesettings"
```


### Client Initialization

```go
client := workspacesettings.NewWorkspaceSettingsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WorkspaceSettingsClient.WorkspaceSettingsCreate`

```go
ctx := context.TODO()
id := workspacesettings.NewWorkspaceSettingID("12345678-1234-9876-4563-123456789012", "workspaceSettingValue")

payload := workspacesettings.WorkspaceSetting{
	// ...
}


read, err := client.WorkspaceSettingsCreate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkspaceSettingsClient.WorkspaceSettingsDelete`

```go
ctx := context.TODO()
id := workspacesettings.NewWorkspaceSettingID("12345678-1234-9876-4563-123456789012", "workspaceSettingValue")

read, err := client.WorkspaceSettingsDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkspaceSettingsClient.WorkspaceSettingsGet`

```go
ctx := context.TODO()
id := workspacesettings.NewWorkspaceSettingID("12345678-1234-9876-4563-123456789012", "workspaceSettingValue")

read, err := client.WorkspaceSettingsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkspaceSettingsClient.WorkspaceSettingsList`

```go
ctx := context.TODO()
id := workspacesettings.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.WorkspaceSettingsList(ctx, id)` can be used to do batched pagination
items, err := client.WorkspaceSettingsListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WorkspaceSettingsClient.WorkspaceSettingsUpdate`

```go
ctx := context.TODO()
id := workspacesettings.NewWorkspaceSettingID("12345678-1234-9876-4563-123456789012", "workspaceSettingValue")

payload := workspacesettings.WorkspaceSetting{
	// ...
}


read, err := client.WorkspaceSettingsUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
