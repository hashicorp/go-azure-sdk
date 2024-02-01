
## `github.com/hashicorp/go-azure-sdk/resource-manager/healthcareapis/2023-12-01/collection` Documentation

The `collection` SDK allows for interaction with the Azure Resource Manager Service `healthcareapis` (API Version `2023-12-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/healthcareapis/2023-12-01/collection"
```


### Client Initialization

```go
client := collection.NewCollectionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CollectionClient.ServicesList`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

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
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ServicesListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.ServicesListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
