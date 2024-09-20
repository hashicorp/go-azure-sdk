
## `github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-04-01/componentversion` Documentation

The `componentversion` SDK allows for interaction with Azure Resource Manager `machinelearningservices` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-04-01/componentversion"
```


### Client Initialization

```go
client := componentversion.NewComponentVersionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ComponentVersionClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := componentversion.NewComponentVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "name", "version")

payload := componentversion.ComponentVersionResource{
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


### Example Usage: `ComponentVersionClient.Delete`

```go
ctx := context.TODO()
id := componentversion.NewComponentVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "name", "version")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComponentVersionClient.Get`

```go
ctx := context.TODO()
id := componentversion.NewComponentVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "name", "version")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComponentVersionClient.List`

```go
ctx := context.TODO()
id := componentversion.NewComponentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "name")

// alternatively `client.List(ctx, id, componentversion.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, componentversion.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ComponentVersionClient.Publish`

```go
ctx := context.TODO()
id := componentversion.NewComponentVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "name", "version")

payload := componentversion.DestinationAsset{
	// ...
}


if err := client.PublishThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ComponentVersionClient.RegistryComponentVersionsCreateOrUpdate`

```go
ctx := context.TODO()
id := componentversion.NewRegistryComponentVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "registryName", "componentName", "version")

payload := componentversion.ComponentVersionResource{
	// ...
}


if err := client.RegistryComponentVersionsCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ComponentVersionClient.RegistryComponentVersionsDelete`

```go
ctx := context.TODO()
id := componentversion.NewRegistryComponentVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "registryName", "componentName", "version")

if err := client.RegistryComponentVersionsDeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ComponentVersionClient.RegistryComponentVersionsGet`

```go
ctx := context.TODO()
id := componentversion.NewRegistryComponentVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "registryName", "componentName", "version")

read, err := client.RegistryComponentVersionsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComponentVersionClient.RegistryComponentVersionsList`

```go
ctx := context.TODO()
id := componentversion.NewRegistryComponentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "registryName", "componentName")

// alternatively `client.RegistryComponentVersionsList(ctx, id, componentversion.DefaultRegistryComponentVersionsListOperationOptions())` can be used to do batched pagination
items, err := client.RegistryComponentVersionsListComplete(ctx, id, componentversion.DefaultRegistryComponentVersionsListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
