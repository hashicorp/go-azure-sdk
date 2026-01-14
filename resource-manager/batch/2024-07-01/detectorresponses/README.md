
## `github.com/hashicorp/go-azure-sdk/resource-manager/batch/2024-07-01/detectorresponses` Documentation

The `detectorresponses` SDK allows for interaction with Azure Resource Manager `batch` (API Version `2024-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/batch/2024-07-01/detectorresponses"
```


### Client Initialization

```go
client := detectorresponses.NewDetectorResponsesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DetectorResponsesClient.BatchAccountGetDetector`

```go
ctx := context.TODO()
id := detectorresponses.NewDetectorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "batchAccountName", "detectorId")

read, err := client.BatchAccountGetDetector(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DetectorResponsesClient.BatchAccountListDetectors`

```go
ctx := context.TODO()
id := detectorresponses.NewBatchAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "batchAccountName")

// alternatively `client.BatchAccountListDetectors(ctx, id)` can be used to do batched pagination
items, err := client.BatchAccountListDetectorsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
