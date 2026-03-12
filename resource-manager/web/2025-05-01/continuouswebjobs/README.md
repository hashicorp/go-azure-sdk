
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/continuouswebjobs` Documentation

The `continuouswebjobs` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/continuouswebjobs"
```


### Client Initialization

```go
client := continuouswebjobs.NewContinuousWebJobsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ContinuousWebJobsClient.WebAppsDeleteContinuousWebJob`

```go
ctx := context.TODO()
id := continuouswebjobs.NewContinuousWebJobID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "continuousWebJobName")

read, err := client.WebAppsDeleteContinuousWebJob(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ContinuousWebJobsClient.WebAppsGetContinuousWebJob`

```go
ctx := context.TODO()
id := continuouswebjobs.NewContinuousWebJobID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "continuousWebJobName")

read, err := client.WebAppsGetContinuousWebJob(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ContinuousWebJobsClient.WebAppsListContinuousWebJobs`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

// alternatively `client.WebAppsListContinuousWebJobs(ctx, id)` can be used to do batched pagination
items, err := client.WebAppsListContinuousWebJobsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ContinuousWebJobsClient.WebAppsStartContinuousWebJob`

```go
ctx := context.TODO()
id := continuouswebjobs.NewContinuousWebJobID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "continuousWebJobName")

read, err := client.WebAppsStartContinuousWebJob(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ContinuousWebJobsClient.WebAppsStopContinuousWebJob`

```go
ctx := context.TODO()
id := continuouswebjobs.NewContinuousWebJobID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "continuousWebJobName")

read, err := client.WebAppsStopContinuousWebJob(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
