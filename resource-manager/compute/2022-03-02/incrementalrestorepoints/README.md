
## `github.com/hashicorp/go-azure-sdk/resource-manager/compute/2022-03-02/incrementalrestorepoints` Documentation

The `incrementalrestorepoints` SDK allows for interaction with the Azure Resource Manager Service `compute` (API Version `2022-03-02`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/compute/2022-03-02/incrementalrestorepoints"
```


### Client Initialization

```go
client := incrementalrestorepoints.NewIncrementalRestorePointsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `IncrementalRestorePointsClient.DiskRestorePointGet`

```go
ctx := context.TODO()
id := incrementalrestorepoints.NewDiskRestorePointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "restorePointCollectionValue", "vmRestorePointValue", "diskRestorePointValue")

read, err := client.DiskRestorePointGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IncrementalRestorePointsClient.DiskRestorePointGrantAccess`

```go
ctx := context.TODO()
id := incrementalrestorepoints.NewDiskRestorePointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "restorePointCollectionValue", "vmRestorePointValue", "diskRestorePointValue")

payload := incrementalrestorepoints.GrantAccessData{
	// ...
}


if err := client.DiskRestorePointGrantAccessThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `IncrementalRestorePointsClient.DiskRestorePointListByRestorePoint`

```go
ctx := context.TODO()
id := incrementalrestorepoints.NewRestorePointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "restorePointCollectionValue", "vmRestorePointValue")

// alternatively `client.DiskRestorePointListByRestorePoint(ctx, id)` can be used to do batched pagination
items, err := client.DiskRestorePointListByRestorePointComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `IncrementalRestorePointsClient.DiskRestorePointRevokeAccess`

```go
ctx := context.TODO()
id := incrementalrestorepoints.NewDiskRestorePointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "restorePointCollectionValue", "vmRestorePointValue", "diskRestorePointValue")

if err := client.DiskRestorePointRevokeAccessThenPoll(ctx, id); err != nil {
	// handle the error
}
```
