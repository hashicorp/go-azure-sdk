
## `github.com/hashicorp/go-azure-sdk/resource-manager/healthcareapis/2021-06-01-preview/collection` Documentation

The `collection` SDK allows for interaction with the Azure Resource Manager Service `healthcareapis` (API Version `2021-06-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/healthcareapis/2021-06-01-preview/collection"
```


### Client Initialization

```go
client := collection.NewCollectionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CollectionClient.ServicesList`

```go
ctx := context.TODO()
id := collection.NewSubscriptionID()

// alternatively `client.ServicesList(ctx, id)` can be used to do batched pagination
items, err := client.ServicesListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CollectionClient.ServicesListByResourceGroup`

```go
ctx := context.TODO()
id := collection.NewResourceGroupID()

// alternatively `client.ServicesListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.ServicesListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
