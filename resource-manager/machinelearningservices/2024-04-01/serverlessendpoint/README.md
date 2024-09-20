
## `github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-04-01/serverlessendpoint` Documentation

The `serverlessendpoint` SDK allows for interaction with Azure Resource Manager `machinelearningservices` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-04-01/serverlessendpoint"
```


### Client Initialization

```go
client := serverlessendpoint.NewServerlessEndpointClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ServerlessEndpointClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := serverlessendpoint.NewServerlessEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "name")

payload := serverlessendpoint.ServerlessEndpointTrackedResource{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ServerlessEndpointClient.Delete`

```go
ctx := context.TODO()
id := serverlessendpoint.NewServerlessEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "name")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ServerlessEndpointClient.Get`

```go
ctx := context.TODO()
id := serverlessendpoint.NewServerlessEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "name")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServerlessEndpointClient.List`

```go
ctx := context.TODO()
id := serverlessendpoint.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName")

// alternatively `client.List(ctx, id, serverlessendpoint.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, serverlessendpoint.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ServerlessEndpointClient.ListKeys`

```go
ctx := context.TODO()
id := serverlessendpoint.NewServerlessEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "name")

read, err := client.ListKeys(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServerlessEndpointClient.RegenerateKeys`

```go
ctx := context.TODO()
id := serverlessendpoint.NewServerlessEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "name")

payload := serverlessendpoint.RegenerateEndpointKeysRequest{
	// ...
}


if err := client.RegenerateKeysThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ServerlessEndpointClient.Update`

```go
ctx := context.TODO()
id := serverlessendpoint.NewServerlessEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "name")

payload := serverlessendpoint.PartialMinimalTrackedResourceWithSkuAndIdentity{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
