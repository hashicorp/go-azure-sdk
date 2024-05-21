
## `github.com/hashicorp/go-azure-sdk/resource-manager/resources/2015-11-01/deploymentoperations` Documentation

The `deploymentoperations` SDK allows for interaction with the Azure Resource Manager Service `resources` (API Version `2015-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/resources/2015-11-01/deploymentoperations"
```


### Client Initialization

```go
client := deploymentoperations.NewDeploymentOperationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DeploymentOperationsClient.Get`

```go
ctx := context.TODO()
id := deploymentoperations.NewOperationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "deploymentValue", "operationIdValue")

read, err := client.Get(ctx, id)
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
id := deploymentoperations.NewDeploymentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "deploymentValue")

// alternatively `client.List(ctx, id, deploymentoperations.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, deploymentoperations.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
