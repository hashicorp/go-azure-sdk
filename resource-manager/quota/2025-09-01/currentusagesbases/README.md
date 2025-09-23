
## `github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-09-01/currentusagesbases` Documentation

The `currentusagesbases` SDK allows for interaction with Azure Resource Manager `quota` (API Version `2025-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-09-01/currentusagesbases"
```


### Client Initialization

```go
client := currentusagesbases.NewCurrentUsagesBasesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CurrentUsagesBasesClient.UsagesGet`

```go
ctx := context.TODO()
id := currentusagesbases.NewScopedUsageID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "usageName")

read, err := client.UsagesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CurrentUsagesBasesClient.UsagesList`

```go
ctx := context.TODO()
id := commonids.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

// alternatively `client.UsagesList(ctx, id)` can be used to do batched pagination
items, err := client.UsagesListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
