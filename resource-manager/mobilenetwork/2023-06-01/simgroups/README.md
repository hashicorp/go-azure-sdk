
## `github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2023-06-01/simgroups` Documentation

The `simgroups` SDK allows for interaction with the Azure Resource Manager Service `mobilenetwork` (API Version `2023-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2023-06-01/simgroups"
```


### Client Initialization

```go
client := simgroups.NewSIMGroupsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SIMGroupsClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := simgroups.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SIMGroupsClient.ListBySubscription`

```go
ctx := context.TODO()
id := simgroups.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ListBySubscription(ctx, id)` can be used to do batched pagination
items, err := client.ListBySubscriptionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
