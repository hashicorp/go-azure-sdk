
## `github.com/hashicorp/go-azure-sdk/resource-manager/redhatopenshift/2023-09-04/syncsets` Documentation

The `syncsets` SDK allows for interaction with Azure Resource Manager `redhatopenshift` (API Version `2023-09-04`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/redhatopenshift/2023-09-04/syncsets"
```


### Client Initialization

```go
client := syncsets.NewSyncSetsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SyncSetsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := syncsets.NewSyncSetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "openShiftClusterValue", "syncSetValue")

payload := syncsets.SyncSet{
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


### Example Usage: `SyncSetsClient.Delete`

```go
ctx := context.TODO()
id := syncsets.NewSyncSetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "openShiftClusterValue", "syncSetValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SyncSetsClient.Get`

```go
ctx := context.TODO()
id := syncsets.NewSyncSetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "openShiftClusterValue", "syncSetValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SyncSetsClient.List`

```go
ctx := context.TODO()
id := syncsets.NewOpenShiftClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "openShiftClusterValue")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SyncSetsClient.Update`

```go
ctx := context.TODO()
id := syncsets.NewSyncSetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "openShiftClusterValue", "syncSetValue")

payload := syncsets.SyncSetUpdate{
	// ...
}


read, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
