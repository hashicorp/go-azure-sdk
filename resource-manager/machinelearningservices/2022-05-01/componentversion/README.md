
## `github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2022-05-01/componentversion` Documentation

The `componentversion` SDK allows for interaction with the Azure Resource Manager Service `machinelearningservices` (API Version `2022-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2022-05-01/componentversion"
```


### Client Initialization

```go
client := componentversion.NewComponentVersionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ComponentVersionClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := componentversion.NewComponentVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "componentValue", "versionValue")

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
id := componentversion.NewComponentVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "componentValue", "versionValue")

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
id := componentversion.NewComponentVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "componentValue", "versionValue")

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
id := componentversion.NewComponentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "componentValue")

// alternatively `client.List(ctx, id, componentversion.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, componentversion.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
