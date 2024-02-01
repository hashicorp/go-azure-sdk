
## `github.com/hashicorp/go-azure-sdk/resource-manager/managedservices/2022-10-01/operations` Documentation

The `operations` SDK allows for interaction with the Azure Resource Manager Service `managedservices` (API Version `2022-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/managedservices/2022-10-01/operations"
```


### Client Initialization

```go
client := operations.NewOperationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `OperationsClient.WithScopeList`

```go
ctx := context.TODO()
id := commonids.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

read, err := client.WithScopeList(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
