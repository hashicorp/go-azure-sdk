
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2024-04-01/operation` Documentation

The `operation` SDK allows for interaction with Azure Resource Manager `recoveryservicesbackup` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2024-04-01/operation"
```


### Client Initialization

```go
client := operation.NewOperationClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `OperationClient.Validate`

```go
ctx := context.TODO()
id := operation.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName")

payload := operation.ValidateOperationRequestResource{
	// ...
}


read, err := client.Validate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
