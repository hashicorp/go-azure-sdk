
## `github.com/hashicorp/go-azure-sdk/resource-manager/resources/2015-11-01/deployments` Documentation

The `deployments` SDK allows for interaction with the Azure Resource Manager Service `resources` (API Version `2015-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/resources/2015-11-01/deployments"
```


### Client Initialization

```go
client := deployments.NewDeploymentsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DeploymentsClient.CalculateTemplateHash`

```go
ctx := context.TODO()
var payload interface{}

read, err := client.CalculateTemplateHash(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeploymentsClient.Cancel`

```go
ctx := context.TODO()
id := deployments.NewProviderDeploymentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "deploymentValue")

read, err := client.Cancel(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeploymentsClient.CheckExistence`

```go
ctx := context.TODO()
id := deployments.NewProviderDeploymentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "deploymentValue")

read, err := client.CheckExistence(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeploymentsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := deployments.NewProviderDeploymentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "deploymentValue")

payload := deployments.Deployment{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DeploymentsClient.Delete`

```go
ctx := context.TODO()
id := deployments.NewProviderDeploymentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "deploymentValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `DeploymentsClient.Get`

```go
ctx := context.TODO()
id := deployments.NewProviderDeploymentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "deploymentValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeploymentsClient.List`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.List(ctx, id, deployments.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, deployments.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DeploymentsClient.Validate`

```go
ctx := context.TODO()
id := deployments.NewProviderDeploymentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "deploymentValue")

payload := deployments.Deployment{
	// ...
}


read, err := client.Validate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
