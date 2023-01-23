
## `github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2022-05-01/componentcontainer` Documentation

The `componentcontainer` SDK allows for interaction with the Azure Resource Manager Service `machinelearningservices` (API Version `2022-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2022-05-01/componentcontainer"
```


### Client Initialization

```go
client := componentcontainer.NewComponentContainerClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ComponentContainerClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := componentcontainer.NewComponentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "componentValue")

payload := componentcontainer.ComponentContainerResource{
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


### Example Usage: `ComponentContainerClient.Delete`

```go
ctx := context.TODO()
id := componentcontainer.NewComponentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "componentValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComponentContainerClient.Get`

```go
ctx := context.TODO()
id := componentcontainer.NewComponentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "componentValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComponentContainerClient.List`

```go
ctx := context.TODO()
id := componentcontainer.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue")

// alternatively `client.List(ctx, id, componentcontainer.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, componentcontainer.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
