
## `github.com/hashicorp/go-azure-sdk/resource-manager/management/2023-04-01/subscriptionundermanagementgroups` Documentation

The `subscriptionundermanagementgroups` SDK allows for interaction with Azure Resource Manager `management` (API Version `2023-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/management/2023-04-01/subscriptionundermanagementgroups"
```


### Client Initialization

```go
client := subscriptionundermanagementgroups.NewSubscriptionUnderManagementGroupsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SubscriptionUnderManagementGroupsClient.ManagementGroupSubscriptionsCreate`

```go
ctx := context.TODO()
id := subscriptionundermanagementgroups.NewSubscriptionID("groupId", "12345678-1234-9876-4563-123456789012")

read, err := client.ManagementGroupSubscriptionsCreate(ctx, id, subscriptionundermanagementgroups.DefaultManagementGroupSubscriptionsCreateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SubscriptionUnderManagementGroupsClient.ManagementGroupSubscriptionsDelete`

```go
ctx := context.TODO()
id := subscriptionundermanagementgroups.NewSubscriptionID("groupId", "12345678-1234-9876-4563-123456789012")

read, err := client.ManagementGroupSubscriptionsDelete(ctx, id, subscriptionundermanagementgroups.DefaultManagementGroupSubscriptionsDeleteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SubscriptionUnderManagementGroupsClient.ManagementGroupSubscriptionsGetSubscription`

```go
ctx := context.TODO()
id := subscriptionundermanagementgroups.NewSubscriptionID("groupId", "12345678-1234-9876-4563-123456789012")

read, err := client.ManagementGroupSubscriptionsGetSubscription(ctx, id, subscriptionundermanagementgroups.DefaultManagementGroupSubscriptionsGetSubscriptionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SubscriptionUnderManagementGroupsClient.ManagementGroupSubscriptionsGetSubscriptionsUnderManagementGroup`

```go
ctx := context.TODO()
id := commonids.NewManagementGroupID("groupId")

// alternatively `client.ManagementGroupSubscriptionsGetSubscriptionsUnderManagementGroup(ctx, id)` can be used to do batched pagination
items, err := client.ManagementGroupSubscriptionsGetSubscriptionsUnderManagementGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
