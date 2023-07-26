
## `github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-07-01/containers` Documentation

The `containers` SDK allows for interaction with the Azure Resource Manager Service `databoxedge` (API Version `2023-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-07-01/containers"
```


### Client Initialization

```go
client := containers.NewContainersClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ContainersClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := containers.NewContainerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceValue", "storageAccountValue", "containerValue")

payload := containers.Container{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ContainersClient.Delete`

```go
ctx := context.TODO()
id := containers.NewContainerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceValue", "storageAccountValue", "containerValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ContainersClient.Get`

```go
ctx := context.TODO()
id := containers.NewContainerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceValue", "storageAccountValue", "containerValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ContainersClient.ListByStorageAccount`

```go
ctx := context.TODO()
id := containers.NewStorageAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceValue", "storageAccountValue")

// alternatively `client.ListByStorageAccount(ctx, id)` can be used to do batched pagination
items, err := client.ListByStorageAccountComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ContainersClient.Refresh`

```go
ctx := context.TODO()
id := containers.NewContainerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceValue", "storageAccountValue", "containerValue")

if err := client.RefreshThenPoll(ctx, id); err != nil {
	// handle the error
}
```
