
## `github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-09-01/groupquotasubscriptionids` Documentation

The `groupquotasubscriptionids` SDK allows for interaction with Azure Resource Manager `quota` (API Version `2025-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-09-01/groupquotasubscriptionids"
```


### Client Initialization

```go
client := groupquotasubscriptionids.NewGroupQuotaSubscriptionIdsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupQuotaSubscriptionIdsClient.GroupQuotaSubscriptionsCreateOrUpdate`

```go
ctx := context.TODO()
id := groupquotasubscriptionids.NewSubscriptionID("managementGroupId", "groupQuotaName", "12345678-1234-9876-4563-123456789012")

if err := client.GroupQuotaSubscriptionsCreateOrUpdateThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `GroupQuotaSubscriptionIdsClient.GroupQuotaSubscriptionsDelete`

```go
ctx := context.TODO()
id := groupquotasubscriptionids.NewSubscriptionID("managementGroupId", "groupQuotaName", "12345678-1234-9876-4563-123456789012")

if err := client.GroupQuotaSubscriptionsDeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `GroupQuotaSubscriptionIdsClient.GroupQuotaSubscriptionsGet`

```go
ctx := context.TODO()
id := groupquotasubscriptionids.NewSubscriptionID("managementGroupId", "groupQuotaName", "12345678-1234-9876-4563-123456789012")

read, err := client.GroupQuotaSubscriptionsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupQuotaSubscriptionIdsClient.GroupQuotaSubscriptionsList`

```go
ctx := context.TODO()
id := groupquotasubscriptionids.NewGroupQuotaID("managementGroupId", "groupQuotaName")

// alternatively `client.GroupQuotaSubscriptionsList(ctx, id)` can be used to do batched pagination
items, err := client.GroupQuotaSubscriptionsListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupQuotaSubscriptionIdsClient.GroupQuotaSubscriptionsUpdate`

```go
ctx := context.TODO()
id := groupquotasubscriptionids.NewSubscriptionID("managementGroupId", "groupQuotaName", "12345678-1234-9876-4563-123456789012")

if err := client.GroupQuotaSubscriptionsUpdateThenPoll(ctx, id); err != nil {
	// handle the error
}
```
