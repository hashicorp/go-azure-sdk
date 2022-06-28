
## `github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2022-05-01/environmentversion` Documentation

The `environmentversion` SDK allows for interaction with the Azure Resource Manager Service `machinelearningservices` (API Version `2022-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2022-05-01/environmentversion"
```


### Client Initialization

```go
client := environmentversion.NewEnvironmentVersionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EnvironmentVersionClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := environmentversion.NewEnvironmentVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "nameValue", "versionValue")

payload := environmentversion.EnvironmentVersionResource{
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


### Example Usage: `EnvironmentVersionClient.Delete`

```go
ctx := context.TODO()
id := environmentversion.NewEnvironmentVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "nameValue", "versionValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EnvironmentVersionClient.Get`

```go
ctx := context.TODO()
id := environmentversion.NewEnvironmentVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "nameValue", "versionValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EnvironmentVersionClient.List`

```go
ctx := context.TODO()
id := environmentversion.NewEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "nameValue")

// alternatively `client.List(ctx, id, environmentversion.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, environmentversion.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
