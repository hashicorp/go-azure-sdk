
## `github.com/hashicorp/go-azure-sdk/resource-manager/batch/2023-05-01/batchmanagements` Documentation

The `batchmanagements` SDK allows for interaction with Azure Resource Manager `batch` (API Version `2023-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/batch/2023-05-01/batchmanagements"
```


### Client Initialization

```go
client := batchmanagements.NewBatchManagementsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BatchManagementsClient.BatchAccountGetDetector`

```go
ctx := context.TODO()
id := batchmanagements.NewDetectorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "batchAccountValue", "detectorIdValue")

read, err := client.BatchAccountGetDetector(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BatchManagementsClient.BatchAccountListDetectors`

```go
ctx := context.TODO()
id := batchmanagements.NewBatchAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "batchAccountValue")

// alternatively `client.BatchAccountListDetectors(ctx, id)` can be used to do batched pagination
items, err := client.BatchAccountListDetectorsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BatchManagementsClient.LocationCheckNameAvailability`

```go
ctx := context.TODO()
id := batchmanagements.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

payload := batchmanagements.CheckNameAvailabilityParameters{
	// ...
}


read, err := client.LocationCheckNameAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
