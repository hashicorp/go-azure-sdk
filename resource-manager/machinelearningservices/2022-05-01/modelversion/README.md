
## `github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2022-05-01/modelversion` Documentation

The `modelversion` SDK allows for interaction with the Azure Resource Manager Service `machinelearningservices` (API Version `2022-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2022-05-01/modelversion"
```


### Client Initialization

```go
client := modelversion.NewModelVersionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `ModelVersionClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := modelversion.NewModelVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "nameValue", "versionValue")

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
id := modelversion.NewModelVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "nameValue", "versionValue")
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
id := modelversion.NewModelVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "nameValue", "versionValue")
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
id := modelversion.NewModelID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "nameValue")
// alternatively `client.List(ctx, id, modelversion.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, modelversion.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
