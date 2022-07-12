
## `github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-11-01/compute` Documentation

The `compute` SDK allows for interaction with the Azure Resource Manager Service `compute` (API Version `2021-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-11-01/compute"
```


### Client Initialization

```go
client := compute.NewComputeClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ComputeClient.RestorePointsCreate`

```go
ctx := context.TODO()
id := compute.NewRestorePointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "restorePointCollectionValue", "restorePointValue")

payload := compute.RestorePoint{
	// ...
}


if err := client.RestorePointsCreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ComputeClient.RestorePointsDelete`

```go
ctx := context.TODO()
id := compute.NewRestorePointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "restorePointCollectionValue", "restorePointValue")

if err := client.RestorePointsDeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ComputeClient.RestorePointsGet`

```go
ctx := context.TODO()
id := compute.NewRestorePointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "restorePointCollectionValue", "restorePointValue")

read, err := client.RestorePointsGet(ctx, id, compute.DefaultRestorePointsGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
