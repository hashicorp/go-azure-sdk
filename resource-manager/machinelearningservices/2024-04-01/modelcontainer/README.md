
## `github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-04-01/modelcontainer` Documentation

The `modelcontainer` SDK allows for interaction with Azure Resource Manager `machinelearningservices` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-04-01/modelcontainer"
```


### Client Initialization

```go
client := modelcontainer.NewModelContainerClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ModelContainerClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := modelcontainer.NewModelID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "modelName")

payload := modelcontainer.ModelContainerResource{
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


### Example Usage: `ModelContainerClient.Delete`

```go
ctx := context.TODO()
id := modelcontainer.NewModelID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "modelName")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ModelContainerClient.Get`

```go
ctx := context.TODO()
id := modelcontainer.NewModelID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "modelName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ModelContainerClient.List`

```go
ctx := context.TODO()
id := modelcontainer.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName")

// alternatively `client.List(ctx, id, modelcontainer.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, modelcontainer.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ModelContainerClient.RegistryModelContainersCreateOrUpdate`

```go
ctx := context.TODO()
id := modelcontainer.NewRegistryModelID("12345678-1234-9876-4563-123456789012", "example-resource-group", "registryName", "modelName")

payload := modelcontainer.ModelContainerResource{
	// ...
}


if err := client.RegistryModelContainersCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ModelContainerClient.RegistryModelContainersDelete`

```go
ctx := context.TODO()
id := modelcontainer.NewRegistryModelID("12345678-1234-9876-4563-123456789012", "example-resource-group", "registryName", "modelName")

if err := client.RegistryModelContainersDeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ModelContainerClient.RegistryModelContainersGet`

```go
ctx := context.TODO()
id := modelcontainer.NewRegistryModelID("12345678-1234-9876-4563-123456789012", "example-resource-group", "registryName", "modelName")

read, err := client.RegistryModelContainersGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ModelContainerClient.RegistryModelContainersList`

```go
ctx := context.TODO()
id := modelcontainer.NewRegistryID("12345678-1234-9876-4563-123456789012", "example-resource-group", "registryName")

// alternatively `client.RegistryModelContainersList(ctx, id, modelcontainer.DefaultRegistryModelContainersListOperationOptions())` can be used to do batched pagination
items, err := client.RegistryModelContainersListComplete(ctx, id, modelcontainer.DefaultRegistryModelContainersListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
