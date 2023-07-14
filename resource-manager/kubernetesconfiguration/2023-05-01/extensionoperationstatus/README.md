
## `github.com/hashicorp/go-azure-sdk/resource-manager/kubernetesconfiguration/2023-05-01/extensionoperationstatus` Documentation

The `extensionoperationstatus` SDK allows for interaction with the Azure Resource Manager Service `kubernetesconfiguration` (API Version `2023-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/kubernetesconfiguration/2023-05-01/extensionoperationstatus"
```


### Client Initialization

```go
client := extensionoperationstatus.NewExtensionOperationStatusClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ExtensionOperationStatusClient.OperationStatusGet`

```go
ctx := context.TODO()
id := extensionoperationstatus.NewOperationScopedID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "extensionValue", "operationIdValue")

read, err := client.OperationStatusGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
