
## `github.com/hashicorp/go-azure-sdk/resource-manager/search/2023-11-01/usages` Documentation

The `usages` SDK allows for interaction with Azure Resource Manager `search` (API Version `2023-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/search/2023-11-01/usages"
```


### Client Initialization

```go
client := usages.NewUsagesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UsagesClient.ListBySubscription`

```go
ctx := context.TODO()
id := usages.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

// alternatively `client.ListBySubscription(ctx, id, usages.DefaultListBySubscriptionOperationOptions())` can be used to do batched pagination
items, err := client.ListBySubscriptionComplete(ctx, id, usages.DefaultListBySubscriptionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UsagesClient.UsageBySubscriptionSku`

```go
ctx := context.TODO()
id := usages.NewUsageID("12345678-1234-9876-4563-123456789012", "locationName", "usageName")

read, err := client.UsageBySubscriptionSku(ctx, id, usages.DefaultUsageBySubscriptionSkuOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
