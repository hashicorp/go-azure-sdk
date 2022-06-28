
## `github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2022-05-01/environmentcontainer` Documentation

The `environmentcontainer` SDK allows for interaction with the Azure Resource Manager Service `machinelearningservices` (API Version `2022-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2022-05-01/environmentcontainer"
```


### Client Initialization

```go
client := environmentcontainer.NewEnvironmentContainerClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `EnvironmentContainerClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := environmentcontainer.NewEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "nameValue")

payload := environmentcontainer.EnvironmentContainerResource{
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


### Example Usage: `EnvironmentContainerClient.Delete`

```go
ctx := context.TODO()
id := environmentcontainer.NewEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "nameValue")
read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EnvironmentContainerClient.Get`

```go
ctx := context.TODO()
id := environmentcontainer.NewEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "nameValue")
read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EnvironmentContainerClient.List`

```go
ctx := context.TODO()
id := environmentcontainer.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue")
// alternatively `client.List(ctx, id, environmentcontainer.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, environmentcontainer.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
