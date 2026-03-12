
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/appserviceenvironments` Documentation

The `appserviceenvironments` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/appserviceenvironments"
```


### Client Initialization

```go
client := appserviceenvironments.NewAppServiceEnvironmentsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AppServiceEnvironmentsClient.CreateOrUpdateWorkerPool`

```go
ctx := context.TODO()
id := appserviceenvironments.NewWorkerPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName", "workerPoolName")

payload := appserviceenvironments.WorkerPoolResource{
	// ...
}


if err := client.CreateOrUpdateWorkerPoolThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `AppServiceEnvironmentsClient.GetWorkerPool`

```go
ctx := context.TODO()
id := appserviceenvironments.NewWorkerPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName", "workerPoolName")

read, err := client.GetWorkerPool(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AppServiceEnvironmentsClient.ListWebWorkerMetricDefinitions`

```go
ctx := context.TODO()
id := appserviceenvironments.NewWorkerPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName", "workerPoolName")

// alternatively `client.ListWebWorkerMetricDefinitions(ctx, id)` can be used to do batched pagination
items, err := client.ListWebWorkerMetricDefinitionsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AppServiceEnvironmentsClient.ListWebWorkerUsages`

```go
ctx := context.TODO()
id := appserviceenvironments.NewWorkerPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName", "workerPoolName")

// alternatively `client.ListWebWorkerUsages(ctx, id)` can be used to do batched pagination
items, err := client.ListWebWorkerUsagesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AppServiceEnvironmentsClient.ListWorkerPoolInstanceMetricDefinitions`

```go
ctx := context.TODO()
id := appserviceenvironments.NewWorkerPoolInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName", "workerPoolName", "instanceName")

// alternatively `client.ListWorkerPoolInstanceMetricDefinitions(ctx, id)` can be used to do batched pagination
items, err := client.ListWorkerPoolInstanceMetricDefinitionsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AppServiceEnvironmentsClient.ListWorkerPoolSkus`

```go
ctx := context.TODO()
id := appserviceenvironments.NewWorkerPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName", "workerPoolName")

// alternatively `client.ListWorkerPoolSkus(ctx, id)` can be used to do batched pagination
items, err := client.ListWorkerPoolSkusComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AppServiceEnvironmentsClient.ListWorkerPools`

```go
ctx := context.TODO()
id := commonids.NewAppServiceEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName")

// alternatively `client.ListWorkerPools(ctx, id)` can be used to do batched pagination
items, err := client.ListWorkerPoolsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AppServiceEnvironmentsClient.UpdateWorkerPool`

```go
ctx := context.TODO()
id := appserviceenvironments.NewWorkerPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName", "workerPoolName")

payload := appserviceenvironments.WorkerPoolResource{
	// ...
}


read, err := client.UpdateWorkerPool(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
