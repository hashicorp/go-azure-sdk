
## `github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2025-06-01/subscriptionquotaitems` Documentation

The `subscriptionquotaitems` SDK allows for interaction with Azure Resource Manager `netapp` (API Version `2025-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2025-06-01/subscriptionquotaitems"
```


### Client Initialization

```go
client := subscriptionquotaitems.NewSubscriptionQuotaItemsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SubscriptionQuotaItemsClient.NetAppResourceQuotaLimitsGet`

```go
ctx := context.TODO()
id := subscriptionquotaitems.NewQuotaLimitID("12345678-1234-9876-4563-123456789012", "locationName", "quotaLimitName")

read, err := client.NetAppResourceQuotaLimitsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SubscriptionQuotaItemsClient.NetAppResourceQuotaLimitsList`

```go
ctx := context.TODO()
id := subscriptionquotaitems.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

// alternatively `client.NetAppResourceQuotaLimitsList(ctx, id)` can be used to do batched pagination
items, err := client.NetAppResourceQuotaLimitsListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
