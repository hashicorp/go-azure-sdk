
## `github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2022-01-10-preview/datastores` Documentation

The `datastores` SDK allows for interaction with the Azure Resource Manager Service `connectedvmware` (API Version `2022-01-10-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2022-01-10-preview/datastores"
```


### Client Initialization

```go
client := datastores.NewDataStoresClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DataStoresClient.Create`

```go
ctx := context.TODO()
id := datastores.NewDataStoreID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataStoreValue")

payload := datastores.Datastore{
	// ...
}


if err := client.CreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DataStoresClient.Delete`

```go
ctx := context.TODO()
id := datastores.NewDataStoreID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataStoreValue")

if err := client.DeleteThenPoll(ctx, id, datastores.DefaultDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `DataStoresClient.Get`

```go
ctx := context.TODO()
id := datastores.NewDataStoreID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataStoreValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataStoresClient.List`

```go
ctx := context.TODO()
id := datastores.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DataStoresClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := datastores.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DataStoresClient.Update`

```go
ctx := context.TODO()
id := datastores.NewDataStoreID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataStoreValue")

payload := datastores.ResourcePatch{
	// ...
}


read, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
