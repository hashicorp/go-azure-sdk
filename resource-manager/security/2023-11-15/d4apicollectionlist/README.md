
## `github.com/hashicorp/go-azure-sdk/resource-manager/security/2023-11-15/d4apicollectionlist` Documentation

The `d4apicollectionlist` SDK allows for interaction with Azure Resource Manager `security` (API Version `2023-11-15`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/security/2023-11-15/d4apicollectionlist"
```


### Client Initialization

```go
client := d4apicollectionlist.NewD4APICollectionListClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `D4APICollectionListClient.APICollectionsListByAzureApiManagementService`

```go
ctx := context.TODO()
id := d4apicollectionlist.NewServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue")

// alternatively `client.APICollectionsListByAzureApiManagementService(ctx, id)` can be used to do batched pagination
items, err := client.APICollectionsListByAzureApiManagementServiceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `D4APICollectionListClient.APICollectionsListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.APICollectionsListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.APICollectionsListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `D4APICollectionListClient.APICollectionsListBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.APICollectionsListBySubscription(ctx, id)` can be used to do batched pagination
items, err := client.APICollectionsListBySubscriptionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
