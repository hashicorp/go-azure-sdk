
## `github.com/hashicorp/go-azure-sdk/resource-manager/portal/2020-09-01-preview/dashboards` Documentation

The `dashboards` SDK allows for interaction with Azure Resource Manager `portal` (API Version `2020-09-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/portal/2020-09-01-preview/dashboards"
```


### Client Initialization

```go
client := dashboards.NewDashboardsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DashboardsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := dashboards.NewDashboardID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dashboardName")

payload := dashboards.Dashboard{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DashboardsClient.Delete`

```go
ctx := context.TODO()
id := dashboards.NewDashboardID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dashboardName")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DashboardsClient.Get`

```go
ctx := context.TODO()
id := dashboards.NewDashboardID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dashboardName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DashboardsClient.ListByResourceGroup`

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


### Example Usage: `DashboardsClient.ListBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ListBySubscription(ctx, id)` can be used to do batched pagination
items, err := client.ListBySubscriptionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DashboardsClient.Update`

```go
ctx := context.TODO()
id := dashboards.NewDashboardID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dashboardName")

payload := dashboards.PatchableDashboard{
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
