
## `github.com/hashicorp/go-azure-sdk/resource-manager/compute/2025-04-01/computerps` Documentation

The `computerps` SDK allows for interaction with Azure Resource Manager `compute` (API Version `2025-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/compute/2025-04-01/computerps"
```


### Client Initialization

```go
client := computerps.NewComputeRPSClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ComputeRPSClient.RestorePointsCreate`

```go
ctx := context.TODO()
id := computerps.NewRestorePointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "restorePointCollectionName", "restorePointName")

payload := computerps.RestorePoint{
	// ...
}


if err := client.RestorePointsCreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ComputeRPSClient.RestorePointsDelete`

```go
ctx := context.TODO()
id := computerps.NewRestorePointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "restorePointCollectionName", "restorePointName")

if err := client.RestorePointsDeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ComputeRPSClient.RestorePointsGet`

```go
ctx := context.TODO()
id := computerps.NewRestorePointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "restorePointCollectionName", "restorePointName")

read, err := client.RestorePointsGet(ctx, id, computerps.DefaultRestorePointsGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
