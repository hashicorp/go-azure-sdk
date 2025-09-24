
## `github.com/hashicorp/go-azure-sdk/resource-manager/storagesync/2022-09-01/storagesyncs` Documentation

The `storagesyncs` SDK allows for interaction with Azure Resource Manager `storagesync` (API Version `2022-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/storagesync/2022-09-01/storagesyncs"
```


### Client Initialization

```go
client := storagesyncs.NewStoragesyncsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `StoragesyncsClient.LocationOperationStatus`

```go
ctx := context.TODO()
id := storagesyncs.NewOperationID("12345678-1234-9876-4563-123456789012", "locationName", "operationId")

read, err := client.LocationOperationStatus(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StoragesyncsClient.StorageSyncServicesCheckNameAvailability`

```go
ctx := context.TODO()
id := storagesyncs.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

payload := storagesyncs.CheckNameAvailabilityParameters{
	// ...
}


read, err := client.StorageSyncServicesCheckNameAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
