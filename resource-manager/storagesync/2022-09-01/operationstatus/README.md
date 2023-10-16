
## `github.com/hashicorp/go-azure-sdk/resource-manager/storagesync/2022-09-01/operationstatus` Documentation

The `operationstatus` SDK allows for interaction with the Azure Resource Manager Service `storagesync` (API Version `2022-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/storagesync/2022-09-01/operationstatus"
```


### Client Initialization

```go
client := operationstatus.NewOperationStatusClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `OperationStatusClient.LocationOperationStatus`

```go
ctx := context.TODO()
id := operationstatus.NewOperationID("12345678-1234-9876-4563-123456789012", "locationValue", "operationIdValue")

read, err := client.LocationOperationStatus(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OperationStatusClient.OperationStatusGet`

```go
ctx := context.TODO()
id := operationstatus.NewWorkflowOperationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "locationValue", "workflowIdValue", "operationIdValue")

read, err := client.OperationStatusGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
