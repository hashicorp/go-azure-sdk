
## `github.com/hashicorp/go-azure-sdk/resource-manager/quota/2023-02-01/quotarequests` Documentation

The `quotarequests` SDK allows for interaction with the Azure Resource Manager Service `quota` (API Version `2023-02-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/quota/2023-02-01/quotarequests"
```


### Client Initialization

```go
client := quotarequests.NewQuotaRequestsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `QuotaRequestsClient.TatusGet`

```go
ctx := context.TODO()
id := quotarequests.NewScopedQuotaRequestID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "quotaRequestValue")

read, err := client.TatusGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `QuotaRequestsClient.TatusList`

```go
ctx := context.TODO()
id := commonids.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

// alternatively `client.TatusList(ctx, id, quotarequests.DefaultTatusListOperationOptions())` can be used to do batched pagination
items, err := client.TatusListComplete(ctx, id, quotarequests.DefaultTatusListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
