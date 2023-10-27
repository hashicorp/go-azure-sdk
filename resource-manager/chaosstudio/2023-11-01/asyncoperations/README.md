
## `github.com/hashicorp/go-azure-sdk/resource-manager/chaosstudio/2023-11-01/asyncoperations` Documentation

The `asyncoperations` SDK allows for interaction with the Azure Resource Manager Service `chaosstudio` (API Version `2023-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/chaosstudio/2023-11-01/asyncoperations"
```


### Client Initialization

```go
client := asyncoperations.NewAsyncOperationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AsyncOperationsClient.OperationStatusesGet`

```go
ctx := context.TODO()
id := asyncoperations.NewOperationsStatusID("12345678-1234-9876-4563-123456789012", "locationValue", "asyncOperationIdValue")

read, err := client.OperationStatusesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
