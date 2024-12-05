
## `github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/onlinedeployment` Documentation

The `onlinedeployment` SDK allows for interaction with Azure Resource Manager `machinelearningservices` (API Version `2024-10-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/onlinedeployment"
```


### Client Initialization

```go
client := onlinedeployment.NewOnlineDeploymentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `OnlineDeploymentClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := onlinedeployment.NewOnlineEndpointDeploymentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "onlineEndpointName", "deploymentName")

payload := onlinedeployment.OnlineDeploymentTrackedResource{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `OnlineDeploymentClient.Delete`

```go
ctx := context.TODO()
id := onlinedeployment.NewOnlineEndpointDeploymentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "onlineEndpointName", "deploymentName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `OnlineDeploymentClient.Get`

```go
ctx := context.TODO()
id := onlinedeployment.NewOnlineEndpointDeploymentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "onlineEndpointName", "deploymentName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnlineDeploymentClient.GetLogs`

```go
ctx := context.TODO()
id := onlinedeployment.NewOnlineEndpointDeploymentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "onlineEndpointName", "deploymentName")

payload := onlinedeployment.DeploymentLogsRequest{
	// ...
}


read, err := client.GetLogs(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnlineDeploymentClient.List`

```go
ctx := context.TODO()
id := onlinedeployment.NewOnlineEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "onlineEndpointName")

// alternatively `client.List(ctx, id, onlinedeployment.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, onlinedeployment.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OnlineDeploymentClient.ListSkus`

```go
ctx := context.TODO()
id := onlinedeployment.NewOnlineEndpointDeploymentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "onlineEndpointName", "deploymentName")

// alternatively `client.ListSkus(ctx, id, onlinedeployment.DefaultListSkusOperationOptions())` can be used to do batched pagination
items, err := client.ListSkusComplete(ctx, id, onlinedeployment.DefaultListSkusOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OnlineDeploymentClient.Update`

```go
ctx := context.TODO()
id := onlinedeployment.NewOnlineEndpointDeploymentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "onlineEndpointName", "deploymentName")

payload := onlinedeployment.PartialMinimalTrackedResourceWithSku{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
