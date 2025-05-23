
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-08-01/managedinstanceoperations` Documentation

The `managedinstanceoperations` SDK allows for interaction with Azure Resource Manager `sql` (API Version `2023-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-08-01/managedinstanceoperations"
```


### Client Initialization

```go
client := managedinstanceoperations.NewManagedInstanceOperationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ManagedInstanceOperationsClient.ListByManagedInstance`

```go
ctx := context.TODO()
id := commonids.NewSqlManagedInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceName")

// alternatively `client.ListByManagedInstance(ctx, id)` can be used to do batched pagination
items, err := client.ListByManagedInstanceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
