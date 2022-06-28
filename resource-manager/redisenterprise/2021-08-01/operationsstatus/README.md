
## `github.com/hashicorp/go-azure-sdk/resource-manager/redisenterprise/2021-08-01/operationsstatus` Documentation

The `operationsstatus` SDK allows for interaction with the Azure Resource Manager Service `redisenterprise` (API Version `2021-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/redisenterprise/2021-08-01/operationsstatus"
```


### Client Initialization

```go
client := operationsstatus.NewOperationsStatusClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `OperationsStatusClient.Get`

```go
ctx := context.TODO()
id := operationsstatus.NewOperationsStatuID("12345678-1234-9876-4563-123456789012", "locationValue", "operationIdValue")
read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
