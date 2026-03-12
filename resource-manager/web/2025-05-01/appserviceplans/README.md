
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/appserviceplans` Documentation

The `appserviceplans` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/appserviceplans"
```


### Client Initialization

```go
client := appserviceplans.NewAppServicePlansClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AppServicePlansClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := commonids.NewAppServicePlanID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverFarmName")

payload := appserviceplans.AppServicePlan{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `AppServicePlansClient.Delete`

```go
ctx := context.TODO()
id := commonids.NewAppServicePlanID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverFarmName")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AppServicePlansClient.Get`

```go
ctx := context.TODO()
id := commonids.NewAppServicePlanID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverFarmName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AppServicePlansClient.GetServerFarmInstanceDetails`

```go
ctx := context.TODO()
id := commonids.NewAppServicePlanID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverFarmName")

read, err := client.GetServerFarmInstanceDetails(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AppServicePlansClient.GetServerFarmRdpPassword`

```go
ctx := context.TODO()
id := commonids.NewAppServicePlanID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverFarmName")

read, err := client.GetServerFarmRdpPassword(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AppServicePlansClient.GetServerFarmSkus`

```go
ctx := context.TODO()
id := commonids.NewAppServicePlanID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverFarmName")

read, err := client.GetServerFarmSkus(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AppServicePlansClient.List`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.List(ctx, id, appserviceplans.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, appserviceplans.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AppServicePlansClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AppServicePlansClient.ListCapabilities`

```go
ctx := context.TODO()
id := commonids.NewAppServicePlanID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverFarmName")

read, err := client.ListCapabilities(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AppServicePlansClient.ListHybridConnections`

```go
ctx := context.TODO()
id := commonids.NewAppServicePlanID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverFarmName")

// alternatively `client.ListHybridConnections(ctx, id)` can be used to do batched pagination
items, err := client.ListHybridConnectionsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AppServicePlansClient.ListUsages`

```go
ctx := context.TODO()
id := commonids.NewAppServicePlanID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverFarmName")

// alternatively `client.ListUsages(ctx, id, appserviceplans.DefaultListUsagesOperationOptions())` can be used to do batched pagination
items, err := client.ListUsagesComplete(ctx, id, appserviceplans.DefaultListUsagesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AppServicePlansClient.ListWebApps`

```go
ctx := context.TODO()
id := commonids.NewAppServicePlanID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverFarmName")

// alternatively `client.ListWebApps(ctx, id, appserviceplans.DefaultListWebAppsOperationOptions())` can be used to do batched pagination
items, err := client.ListWebAppsComplete(ctx, id, appserviceplans.DefaultListWebAppsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AppServicePlansClient.RebootWorker`

```go
ctx := context.TODO()
id := appserviceplans.NewWorkerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverFarmName", "workerName")

read, err := client.RebootWorker(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AppServicePlansClient.RecycleManagedInstanceWorker`

```go
ctx := context.TODO()
id := appserviceplans.NewWorkerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverFarmName", "workerName")

read, err := client.RecycleManagedInstanceWorker(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AppServicePlansClient.RestartWebApps`

```go
ctx := context.TODO()
id := commonids.NewAppServicePlanID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverFarmName")

read, err := client.RestartWebApps(ctx, id, appserviceplans.DefaultRestartWebAppsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AppServicePlansClient.Update`

```go
ctx := context.TODO()
id := commonids.NewAppServicePlanID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverFarmName")

payload := appserviceplans.AppServicePlanPatchResource{
	// ...
}


read, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
