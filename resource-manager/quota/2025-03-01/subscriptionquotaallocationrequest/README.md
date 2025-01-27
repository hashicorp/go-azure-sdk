
## `github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-03-01/subscriptionquotaallocationrequest` Documentation

The `subscriptionquotaallocationrequest` SDK allows for interaction with Azure Resource Manager `quota` (API Version `2025-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-03-01/subscriptionquotaallocationrequest"
```


### Client Initialization

```go
client := subscriptionquotaallocationrequest.NewSubscriptionQuotaAllocationRequestClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SubscriptionQuotaAllocationRequestClient.GroupQuotaSubscriptionAllocationRequestGet`

```go
ctx := context.TODO()
id := subscriptionquotaallocationrequest.NewQuotaAllocationRequestID("managementGroupId", "12345678-1234-9876-4563-123456789012", "groupQuotaName", "resourceProviderName", "allocationId")

read, err := client.GroupQuotaSubscriptionAllocationRequestGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SubscriptionQuotaAllocationRequestClient.GroupQuotaSubscriptionAllocationRequestList`

```go
ctx := context.TODO()
id := subscriptionquotaallocationrequest.NewGroupQuotaResourceProviderID("managementGroupId", "12345678-1234-9876-4563-123456789012", "groupQuotaName", "resourceProviderName")

// alternatively `client.GroupQuotaSubscriptionAllocationRequestList(ctx, id, subscriptionquotaallocationrequest.DefaultGroupQuotaSubscriptionAllocationRequestListOperationOptions())` can be used to do batched pagination
items, err := client.GroupQuotaSubscriptionAllocationRequestListComplete(ctx, id, subscriptionquotaallocationrequest.DefaultGroupQuotaSubscriptionAllocationRequestListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
