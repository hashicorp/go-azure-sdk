
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-11-01/managedinstanceprivatelinkresources` Documentation

The `managedinstanceprivatelinkresources` SDK allows for interaction with the Azure Resource Manager Service `sql` (API Version `2021-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-11-01/managedinstanceprivatelinkresources"
```


### Client Initialization

```go
client := managedinstanceprivatelinkresources.NewManagedInstancePrivateLinkResourcesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ManagedInstancePrivateLinkResourcesClient.Get`

```go
ctx := context.TODO()
id := managedinstanceprivatelinkresources.NewManagedInstancePrivateLinkResourceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceValue", "privateLinkResourceValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedInstancePrivateLinkResourcesClient.ListByManagedInstance`

```go
ctx := context.TODO()
id := managedinstanceprivatelinkresources.NewSqlManagedInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceValue")

// alternatively `client.ListByManagedInstance(ctx, id)` can be used to do batched pagination
items, err := client.ListByManagedInstanceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
