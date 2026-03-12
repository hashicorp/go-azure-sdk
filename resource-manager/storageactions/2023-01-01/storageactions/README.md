
## `github.com/hashicorp/go-azure-sdk/resource-manager/storageactions/2023-01-01/storageactions` Documentation

The `storageactions` SDK allows for interaction with Azure Resource Manager `storageactions` (API Version `2023-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/storageactions/2023-01-01/storageactions"
```


### Client Initialization

```go
client := storageactions.NewStorageactionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `StorageactionsClient.StorageTasksPreviewActions`

```go
ctx := context.TODO()
id := storageactions.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

payload := storageactions.StorageTaskPreviewAction{
	// ...
}


read, err := client.StorageTasksPreviewActions(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
