
## `github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/workspacemanagedsqlserverdedicatedsqlminimaltlssettings` Documentation

The `workspacemanagedsqlserverdedicatedsqlminimaltlssettings` SDK allows for interaction with the Azure Resource Manager Service `synapse` (API Version `2021-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/workspacemanagedsqlserverdedicatedsqlminimaltlssettings"
```


### Client Initialization

```go
client := workspacemanagedsqlserverdedicatedsqlminimaltlssettings.NewWorkspaceManagedSqlServerDedicatedSQLminimalTlsSettingsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WorkspaceManagedSqlServerDedicatedSQLminimalTlsSettingsClient.Get`

```go
ctx := context.TODO()
id := workspacemanagedsqlserverdedicatedsqlminimaltlssettings.NewDedicatedSQLMinimalTLSSettingID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "dedicatedSQLMinimalTLSSettingValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkspaceManagedSqlServerDedicatedSQLminimalTlsSettingsClient.List`

```go
ctx := context.TODO()
id := workspacemanagedsqlserverdedicatedsqlminimaltlssettings.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WorkspaceManagedSqlServerDedicatedSQLminimalTlsSettingsClient.Update`

```go
ctx := context.TODO()
id := workspacemanagedsqlserverdedicatedsqlminimaltlssettings.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue")

payload := workspacemanagedsqlserverdedicatedsqlminimaltlssettings.DedicatedSQLMinimalTLSSettings{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
