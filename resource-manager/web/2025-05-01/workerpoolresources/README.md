
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/workerpoolresources` Documentation

The `workerpoolresources` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/workerpoolresources"
```


### Client Initialization

```go
client := workerpoolresources.NewWorkerPoolResourcesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WorkerPoolResourcesClient.AppServiceEnvironmentsCreateOrUpdateMultiRolePool`

```go
ctx := context.TODO()
id := commonids.NewAppServiceEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName")

payload := workerpoolresources.WorkerPoolResource{
	// ...
}


if err := client.AppServiceEnvironmentsCreateOrUpdateMultiRolePoolThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `WorkerPoolResourcesClient.AppServiceEnvironmentsGetMultiRolePool`

```go
ctx := context.TODO()
id := commonids.NewAppServiceEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName")

read, err := client.AppServiceEnvironmentsGetMultiRolePool(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkerPoolResourcesClient.AppServiceEnvironmentsListMultiRoleMetricDefinitions`

```go
ctx := context.TODO()
id := commonids.NewAppServiceEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName")

// alternatively `client.AppServiceEnvironmentsListMultiRoleMetricDefinitions(ctx, id)` can be used to do batched pagination
items, err := client.AppServiceEnvironmentsListMultiRoleMetricDefinitionsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WorkerPoolResourcesClient.AppServiceEnvironmentsListMultiRolePoolInstanceMetricDefinitions`

```go
ctx := context.TODO()
id := workerpoolresources.NewDefaultInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName", "instanceName")

// alternatively `client.AppServiceEnvironmentsListMultiRolePoolInstanceMetricDefinitions(ctx, id)` can be used to do batched pagination
items, err := client.AppServiceEnvironmentsListMultiRolePoolInstanceMetricDefinitionsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WorkerPoolResourcesClient.AppServiceEnvironmentsListMultiRolePoolSkus`

```go
ctx := context.TODO()
id := commonids.NewAppServiceEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName")

// alternatively `client.AppServiceEnvironmentsListMultiRolePoolSkus(ctx, id)` can be used to do batched pagination
items, err := client.AppServiceEnvironmentsListMultiRolePoolSkusComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WorkerPoolResourcesClient.AppServiceEnvironmentsListMultiRolePools`

```go
ctx := context.TODO()
id := commonids.NewAppServiceEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName")

// alternatively `client.AppServiceEnvironmentsListMultiRolePools(ctx, id)` can be used to do batched pagination
items, err := client.AppServiceEnvironmentsListMultiRolePoolsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WorkerPoolResourcesClient.AppServiceEnvironmentsListMultiRoleUsages`

```go
ctx := context.TODO()
id := commonids.NewAppServiceEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName")

// alternatively `client.AppServiceEnvironmentsListMultiRoleUsages(ctx, id)` can be used to do batched pagination
items, err := client.AppServiceEnvironmentsListMultiRoleUsagesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WorkerPoolResourcesClient.AppServiceEnvironmentsUpdateMultiRolePool`

```go
ctx := context.TODO()
id := commonids.NewAppServiceEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName")

payload := workerpoolresources.WorkerPoolResource{
	// ...
}


read, err := client.AppServiceEnvironmentsUpdateMultiRolePool(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
