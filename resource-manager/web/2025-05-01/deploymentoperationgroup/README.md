
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/deploymentoperationgroup` Documentation

The `deploymentoperationgroup` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/deploymentoperationgroup"
```


### Client Initialization

```go
client := deploymentoperationgroup.NewDeploymentOperationGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DeploymentOperationGroupClient.WebAppsCreateDeploymentSlot`

```go
ctx := context.TODO()
id := deploymentoperationgroup.NewSlotDeploymentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "deploymentName")

payload := deploymentoperationgroup.Deployment{
	// ...
}


read, err := client.WebAppsCreateDeploymentSlot(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeploymentOperationGroupClient.WebAppsDeleteDeploymentSlot`

```go
ctx := context.TODO()
id := deploymentoperationgroup.NewSlotDeploymentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "deploymentName")

read, err := client.WebAppsDeleteDeploymentSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeploymentOperationGroupClient.WebAppsGetDeploymentSlot`

```go
ctx := context.TODO()
id := deploymentoperationgroup.NewSlotDeploymentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "deploymentName")

read, err := client.WebAppsGetDeploymentSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeploymentOperationGroupClient.WebAppsListDeploymentLogSlot`

```go
ctx := context.TODO()
id := deploymentoperationgroup.NewSlotDeploymentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "deploymentName")

read, err := client.WebAppsListDeploymentLogSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeploymentOperationGroupClient.WebAppsListDeploymentsSlot`

```go
ctx := context.TODO()
id := deploymentoperationgroup.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

// alternatively `client.WebAppsListDeploymentsSlot(ctx, id)` can be used to do batched pagination
items, err := client.WebAppsListDeploymentsSlotComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
