
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/triggeredjobhistoryoperationgroup` Documentation

The `triggeredjobhistoryoperationgroup` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/triggeredjobhistoryoperationgroup"
```


### Client Initialization

```go
client := triggeredjobhistoryoperationgroup.NewTriggeredJobHistoryOperationGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TriggeredJobHistoryOperationGroupClient.WebAppsGetTriggeredWebJobHistory`

```go
ctx := context.TODO()
id := triggeredjobhistoryoperationgroup.NewHistoryID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "triggeredWebJobName", "historyName")

read, err := client.WebAppsGetTriggeredWebJobHistory(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TriggeredJobHistoryOperationGroupClient.WebAppsListTriggeredWebJobHistory`

```go
ctx := context.TODO()
id := triggeredjobhistoryoperationgroup.NewTriggeredWebJobID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "triggeredWebJobName")

// alternatively `client.WebAppsListTriggeredWebJobHistory(ctx, id)` can be used to do batched pagination
items, err := client.WebAppsListTriggeredWebJobHistoryComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
