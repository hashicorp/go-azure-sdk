
## `github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/workspacemanagerconfigurations` Documentation

The `workspacemanagerconfigurations` SDK allows for interaction with Azure Resource Manager `securityinsights` (API Version `2023-12-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/workspacemanagerconfigurations"
```


### Client Initialization

```go
client := workspacemanagerconfigurations.NewWorkspaceManagerConfigurationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WorkspaceManagerConfigurationsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := workspacemanagerconfigurations.NewWorkspaceManagerConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "workspaceManagerConfigurationName")

payload := workspacemanagerconfigurations.WorkspaceManagerConfiguration{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkspaceManagerConfigurationsClient.Delete`

```go
ctx := context.TODO()
id := workspacemanagerconfigurations.NewWorkspaceManagerConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "workspaceManagerConfigurationName")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkspaceManagerConfigurationsClient.Get`

```go
ctx := context.TODO()
id := workspacemanagerconfigurations.NewWorkspaceManagerConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "workspaceManagerConfigurationName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkspaceManagerConfigurationsClient.List`

```go
ctx := context.TODO()
id := workspacemanagerconfigurations.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName")

// alternatively `client.List(ctx, id, workspacemanagerconfigurations.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, workspacemanagerconfigurations.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
