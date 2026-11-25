
## `github.com/hashicorp/go-azure-sdk/resource-manager/storagemover/2026-05-01/connections` Documentation

The `connections` SDK allows for interaction with Azure Resource Manager `storagemover` (API Version `2026-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/storagemover/2026-05-01/connections"
```


### Client Initialization

```go
client := connections.NewConnectionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ConnectionsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := connections.NewConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageMoverName", "connectionName")

payload := connections.Connection{
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


### Example Usage: `ConnectionsClient.Delete`

```go
ctx := context.TODO()
id := connections.NewConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageMoverName", "connectionName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ConnectionsClient.Get`

```go
ctx := context.TODO()
id := connections.NewConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageMoverName", "connectionName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConnectionsClient.List`

```go
ctx := context.TODO()
id := connections.NewStorageMoverID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageMoverName")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
