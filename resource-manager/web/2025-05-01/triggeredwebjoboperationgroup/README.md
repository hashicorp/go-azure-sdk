
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/triggeredwebjoboperationgroup` Documentation

The `triggeredwebjoboperationgroup` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/triggeredwebjoboperationgroup"
```


### Client Initialization

```go
client := triggeredwebjoboperationgroup.NewTriggeredWebJobOperationGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TriggeredWebJobOperationGroupClient.WebAppsDeleteTriggeredWebJob`

```go
ctx := context.TODO()
id := triggeredwebjoboperationgroup.NewTriggeredWebJobID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "triggeredWebJobName")

read, err := client.WebAppsDeleteTriggeredWebJob(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TriggeredWebJobOperationGroupClient.WebAppsGetTriggeredWebJob`

```go
ctx := context.TODO()
id := triggeredwebjoboperationgroup.NewTriggeredWebJobID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "triggeredWebJobName")

read, err := client.WebAppsGetTriggeredWebJob(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TriggeredWebJobOperationGroupClient.WebAppsListTriggeredWebJobs`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

// alternatively `client.WebAppsListTriggeredWebJobs(ctx, id)` can be used to do batched pagination
items, err := client.WebAppsListTriggeredWebJobsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TriggeredWebJobOperationGroupClient.WebAppsRunTriggeredWebJob`

```go
ctx := context.TODO()
id := triggeredwebjoboperationgroup.NewTriggeredWebJobID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "triggeredWebJobName")

read, err := client.WebAppsRunTriggeredWebJob(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
