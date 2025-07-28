
## `github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-07-15/quotainformation` Documentation

The `quotainformation` SDK allows for interaction with Azure Resource Manager `quota` (API Version `2025-07-15`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-07-15/quotainformation"
```


### Client Initialization

```go
client := quotainformation.NewQuotaInformationClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `QuotaInformationClient.QuotaCreateOrUpdate`

```go
ctx := context.TODO()
id := quotainformation.NewScopedQuotaID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "quotaName")

payload := quotainformation.CurrentQuotaLimitBase{
	// ...
}


if err := client.QuotaCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `QuotaInformationClient.QuotaGet`

```go
ctx := context.TODO()
id := quotainformation.NewScopedQuotaID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "quotaName")

read, err := client.QuotaGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `QuotaInformationClient.QuotaList`

```go
ctx := context.TODO()
id := commonids.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

// alternatively `client.QuotaList(ctx, id)` can be used to do batched pagination
items, err := client.QuotaListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `QuotaInformationClient.QuotaUpdate`

```go
ctx := context.TODO()
id := quotainformation.NewScopedQuotaID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "quotaName")

payload := quotainformation.CurrentQuotaLimitBase{
	// ...
}


if err := client.QuotaUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
