
## `github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-03-01/usagesinformation` Documentation

The `usagesinformation` SDK allows for interaction with Azure Resource Manager `quota` (API Version `2025-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/quota/2025-03-01/usagesinformation"
```


### Client Initialization

```go
client := usagesinformation.NewUsagesInformationClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UsagesInformationClient.UsagesGet`

```go
ctx := context.TODO()
id := usagesinformation.NewScopedUsageID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "usageName")

read, err := client.UsagesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UsagesInformationClient.UsagesList`

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
