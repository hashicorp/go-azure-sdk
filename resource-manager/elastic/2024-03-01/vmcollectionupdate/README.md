
## `github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2024-03-01/vmcollectionupdate` Documentation

The `vmcollectionupdate` SDK allows for interaction with Azure Resource Manager `elastic` (API Version `2024-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2024-03-01/vmcollectionupdate"
```


### Client Initialization

```go
client := vmcollectionupdate.NewVMCollectionUpdateClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `VMCollectionUpdateClient.VMCollectionUpdate`

```go
ctx := context.TODO()
id := vmcollectionupdate.NewMonitorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorValue")

payload := vmcollectionupdate.VMCollectionUpdate{
	// ...
}


read, err := client.VMCollectionUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
