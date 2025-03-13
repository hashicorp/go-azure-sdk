
## `github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/componentcontainer` Documentation

The `componentcontainer` SDK allows for interaction with Azure Resource Manager `machinelearningservices` (API Version `2024-10-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/componentcontainer"
```


### Client Initialization

```go
client := componentcontainer.NewComponentContainerClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ComponentContainerClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := componentcontainer.NewComponentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "componentName")

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
id := componentcontainer.NewComponentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "componentName")

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
id := componentcontainer.NewComponentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "componentName")

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
id := componentcontainer.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName")

// alternatively `client.List(ctx, id, componentcontainer.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, componentcontainer.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ComponentContainerClient.RegistryComponentContainersCreateOrUpdate`

```go
ctx := context.TODO()
id := componentcontainer.NewRegistryComponentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "registryName", "componentName")

payload := componentcontainer.ComponentContainerResource{
	// ...
}


if err := client.RegistryComponentContainersCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ComponentContainerClient.RegistryComponentContainersDelete`

```go
ctx := context.TODO()
id := componentcontainer.NewRegistryComponentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "registryName", "componentName")

if err := client.RegistryComponentContainersDeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ComponentContainerClient.RegistryComponentContainersGet`

```go
ctx := context.TODO()
id := componentcontainer.NewRegistryComponentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "registryName", "componentName")

read, err := client.RegistryComponentContainersGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComponentContainerClient.RegistryComponentContainersList`

```go
ctx := context.TODO()
id := componentcontainer.NewRegistryID("12345678-1234-9876-4563-123456789012", "example-resource-group", "registryName")

// alternatively `client.RegistryComponentContainersList(ctx, id, componentcontainer.DefaultRegistryComponentContainersListOperationOptions())` can be used to do batched pagination
items, err := client.RegistryComponentContainersListComplete(ctx, id, componentcontainer.DefaultRegistryComponentContainersListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
