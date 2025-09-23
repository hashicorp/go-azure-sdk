
## `github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-09-01/quotarequeststatus` Documentation

The `quotarequeststatus` SDK allows for interaction with Azure Resource Manager `quota` (API Version `2025-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-09-01/quotarequeststatus"
```


### Client Initialization

```go
client := quotarequeststatus.NewQuotaRequestStatusClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `QuotaRequestStatusClient.Get`

```go
ctx := context.TODO()
id := quotarequeststatus.NewScopedQuotaRequestID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "quotaRequestName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `QuotaRequestStatusClient.List`

```go
ctx := context.TODO()
id := commonids.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

// alternatively `client.List(ctx, id, quotarequeststatus.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, quotarequeststatus.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
