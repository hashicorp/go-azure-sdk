
## `github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-09-01/subscriptionquotaallocationslists` Documentation

The `subscriptionquotaallocationslists` SDK allows for interaction with Azure Resource Manager `quota` (API Version `2025-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-09-01/subscriptionquotaallocationslists"
```


### Client Initialization

```go
client := subscriptionquotaallocationslists.NewSubscriptionQuotaAllocationsListsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SubscriptionQuotaAllocationsListsClient.GroupQuotaSubscriptionAllocationList`

```go
ctx := context.TODO()
id := subscriptionquotaallocationslists.NewQuotaAllocationID("managementGroupId", "12345678-1234-9876-4563-123456789012", "groupQuotaName", "resourceProviderName", "quotaAllocationName")

read, err := client.GroupQuotaSubscriptionAllocationList(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SubscriptionQuotaAllocationsListsClient.GroupQuotaSubscriptionAllocationRequestUpdate`

```go
ctx := context.TODO()
id := subscriptionquotaallocationslists.NewQuotaAllocationID("managementGroupId", "12345678-1234-9876-4563-123456789012", "groupQuotaName", "resourceProviderName", "quotaAllocationName")

payload := subscriptionquotaallocationslists.SubscriptionQuotaAllocationsList{
	// ...
}


if err := client.GroupQuotaSubscriptionAllocationRequestUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
