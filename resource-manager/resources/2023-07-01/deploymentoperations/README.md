
## `github.com/hashicorp/go-azure-sdk/resource-manager/resources/2023-07-01/deploymentoperations` Documentation

The `deploymentoperations` SDK allows for interaction with Azure Resource Manager `resources` (API Version `2023-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/resources/2023-07-01/deploymentoperations"
```


### Client Initialization

```go
client := deploymentoperations.NewDeploymentOperationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DeploymentOperationsClient.Get`

```go
ctx := context.TODO()
id := deploymentoperations.NewResourceGroupDeploymentOperationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "deploymentName", "operationId")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeploymentOperationsClient.GetAtManagementGroupScope`

```go
ctx := context.TODO()
id := deploymentoperations.NewProviders2DeploymentOperationID("groupId", "deploymentName", "operationId")

read, err := client.GetAtManagementGroupScope(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeploymentOperationsClient.GetAtScope`

```go
ctx := context.TODO()
id := deploymentoperations.NewScopedOperationID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "deploymentName", "operationId")

read, err := client.GetAtScope(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeploymentOperationsClient.GetAtSubscriptionScope`

```go
ctx := context.TODO()
id := deploymentoperations.NewDeploymentOperationID("12345678-1234-9876-4563-123456789012", "deploymentName", "operationId")

read, err := client.GetAtSubscriptionScope(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeploymentOperationsClient.GetAtTenantScope`

```go
ctx := context.TODO()
id := deploymentoperations.NewOperationID("deploymentName", "operationId")

read, err := client.GetAtTenantScope(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeploymentOperationsClient.List`

```go
ctx := context.TODO()
id := deploymentoperations.NewResourceGroupDeploymentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "deploymentName")

// alternatively `client.List(ctx, id, deploymentoperations.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, deploymentoperations.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DeploymentOperationsClient.ListAtManagementGroupScope`

```go
ctx := context.TODO()
id := deploymentoperations.NewProviders2DeploymentID("groupId", "deploymentName")

// alternatively `client.ListAtManagementGroupScope(ctx, id, deploymentoperations.DefaultListAtManagementGroupScopeOperationOptions())` can be used to do batched pagination
items, err := client.ListAtManagementGroupScopeComplete(ctx, id, deploymentoperations.DefaultListAtManagementGroupScopeOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DeploymentOperationsClient.ListAtScope`

```go
ctx := context.TODO()
id := deploymentoperations.NewScopedDeploymentID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "deploymentName")

// alternatively `client.ListAtScope(ctx, id, deploymentoperations.DefaultListAtScopeOperationOptions())` can be used to do batched pagination
items, err := client.ListAtScopeComplete(ctx, id, deploymentoperations.DefaultListAtScopeOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DeploymentOperationsClient.ListAtSubscriptionScope`

```go
ctx := context.TODO()
id := deploymentoperations.NewProviderDeploymentID("12345678-1234-9876-4563-123456789012", "deploymentName")

// alternatively `client.ListAtSubscriptionScope(ctx, id, deploymentoperations.DefaultListAtSubscriptionScopeOperationOptions())` can be used to do batched pagination
items, err := client.ListAtSubscriptionScopeComplete(ctx, id, deploymentoperations.DefaultListAtSubscriptionScopeOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DeploymentOperationsClient.ListAtTenantScope`

```go
ctx := context.TODO()
id := deploymentoperations.NewDeploymentID("deploymentName")

// alternatively `client.ListAtTenantScope(ctx, id, deploymentoperations.DefaultListAtTenantScopeOperationOptions())` can be used to do batched pagination
items, err := client.ListAtTenantScopeComplete(ctx, id, deploymentoperations.DefaultListAtTenantScopeOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
