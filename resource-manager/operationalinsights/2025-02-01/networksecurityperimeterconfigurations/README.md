
## `github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2025-02-01/networksecurityperimeterconfigurations` Documentation

The `networksecurityperimeterconfigurations` SDK allows for interaction with Azure Resource Manager `operationalinsights` (API Version `2025-02-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2025-02-01/networksecurityperimeterconfigurations"
```


### Client Initialization

```go
client := networksecurityperimeterconfigurations.NewNetworkSecurityPerimeterConfigurationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `NetworkSecurityPerimeterConfigurationsClient.WorkspacesGetNSP`

```go
ctx := context.TODO()
id := networksecurityperimeterconfigurations.NewNetworkSecurityPerimeterConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "networkSecurityPerimeterConfigurationName")

read, err := client.WorkspacesGetNSP(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetworkSecurityPerimeterConfigurationsClient.WorkspacesListNSP`

```go
ctx := context.TODO()
id := networksecurityperimeterconfigurations.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName")

// alternatively `client.WorkspacesListNSP(ctx, id)` can be used to do batched pagination
items, err := client.WorkspacesListNSPComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NetworkSecurityPerimeterConfigurationsClient.WorkspacesReconcileNSP`

```go
ctx := context.TODO()
id := networksecurityperimeterconfigurations.NewNetworkSecurityPerimeterConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "networkSecurityPerimeterConfigurationName")

if err := client.WorkspacesReconcileNSPThenPoll(ctx, id); err != nil {
	// handle the error
}
```
