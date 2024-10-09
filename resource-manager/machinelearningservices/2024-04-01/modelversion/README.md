
## `github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-04-01/modelversion` Documentation

The `modelversion` SDK allows for interaction with Azure Resource Manager `machinelearningservices` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-04-01/modelversion"
```


### Client Initialization

```go
client := modelversion.NewModelVersionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ModelVersionClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := modelversion.NewModelVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "modelName", "versionName")

payload := modelversion.ModelVersionResource{
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


### Example Usage: `ModelVersionClient.Delete`

```go
ctx := context.TODO()
id := modelversion.NewModelVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "modelName", "versionName")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ModelVersionClient.Get`

```go
ctx := context.TODO()
id := modelversion.NewModelVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "modelName", "versionName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ModelVersionClient.List`

```go
ctx := context.TODO()
id := modelversion.NewModelID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "modelName")

// alternatively `client.List(ctx, id, modelversion.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, modelversion.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ModelVersionClient.Publish`

```go
ctx := context.TODO()
id := modelversion.NewModelVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "modelName", "versionName")

payload := modelversion.DestinationAsset{
	// ...
}


if err := client.PublishThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ModelVersionClient.RegistryModelVersionsCreateOrGetStartPendingUpload`

```go
ctx := context.TODO()
id := modelversion.NewRegistryModelVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "registryName", "modelName", "versionName")

payload := modelversion.PendingUploadRequestDto{
	// ...
}


read, err := client.RegistryModelVersionsCreateOrGetStartPendingUpload(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ModelVersionClient.RegistryModelVersionsCreateOrUpdate`

```go
ctx := context.TODO()
id := modelversion.NewRegistryModelVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "registryName", "modelName", "versionName")

payload := modelversion.ModelVersionResource{
	// ...
}


if err := client.RegistryModelVersionsCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ModelVersionClient.RegistryModelVersionsDelete`

```go
ctx := context.TODO()
id := modelversion.NewRegistryModelVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "registryName", "modelName", "versionName")

if err := client.RegistryModelVersionsDeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ModelVersionClient.RegistryModelVersionsGet`

```go
ctx := context.TODO()
id := modelversion.NewRegistryModelVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "registryName", "modelName", "versionName")

read, err := client.RegistryModelVersionsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ModelVersionClient.RegistryModelVersionsList`

```go
ctx := context.TODO()
id := modelversion.NewRegistryModelID("12345678-1234-9876-4563-123456789012", "example-resource-group", "registryName", "modelName")

// alternatively `client.RegistryModelVersionsList(ctx, id, modelversion.DefaultRegistryModelVersionsListOperationOptions())` can be used to do batched pagination
items, err := client.RegistryModelVersionsListComplete(ctx, id, modelversion.DefaultRegistryModelVersionsListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
