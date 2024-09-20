
## `github.com/hashicorp/go-azure-sdk/resource-manager/security/2020-01-01/alerts` Documentation

The `alerts` SDK allows for interaction with Azure Resource Manager `security` (API Version `2020-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/security/2020-01-01/alerts"
```


### Client Initialization

```go
client := alerts.NewAlertsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AlertsClient.GetResourceGroupLevelAlerts`

```go
ctx := context.TODO()
id := alerts.NewLocationAlertID("12345678-1234-9876-4563-123456789012", "example-resource-group", "ascLocation", "alertName")

read, err := client.GetResourceGroupLevelAlerts(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AlertsClient.GetSubscriptionLevelAlert`

```go
ctx := context.TODO()
id := alerts.NewAlertID("12345678-1234-9876-4563-123456789012", "ascLocation", "alertName")

read, err := client.GetSubscriptionLevelAlert(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AlertsClient.List`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AlertsClient.ListByResourceGroup`

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


### Example Usage: `AlertsClient.ListResourceGroupLevelAlertsByRegion`

```go
ctx := context.TODO()
id := alerts.NewProviderLocationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "ascLocation")

// alternatively `client.ListResourceGroupLevelAlertsByRegion(ctx, id)` can be used to do batched pagination
items, err := client.ListResourceGroupLevelAlertsByRegionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AlertsClient.ListSubscriptionLevelAlertsByRegion`

```go
ctx := context.TODO()
id := alerts.NewLocationID("12345678-1234-9876-4563-123456789012", "ascLocation")

// alternatively `client.ListSubscriptionLevelAlertsByRegion(ctx, id)` can be used to do batched pagination
items, err := client.ListSubscriptionLevelAlertsByRegionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AlertsClient.UpdateResourceGroupLevelAlertStateToDismiss`

```go
ctx := context.TODO()
id := alerts.NewLocationAlertID("12345678-1234-9876-4563-123456789012", "example-resource-group", "ascLocation", "alertName")

read, err := client.UpdateResourceGroupLevelAlertStateToDismiss(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AlertsClient.UpdateResourceGroupLevelAlertStateToReactivate`

```go
ctx := context.TODO()
id := alerts.NewLocationAlertID("12345678-1234-9876-4563-123456789012", "example-resource-group", "ascLocation", "alertName")

read, err := client.UpdateResourceGroupLevelAlertStateToReactivate(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AlertsClient.UpdateResourceGroupLevelStateToResolve`

```go
ctx := context.TODO()
id := alerts.NewLocationAlertID("12345678-1234-9876-4563-123456789012", "example-resource-group", "ascLocation", "alertName")

read, err := client.UpdateResourceGroupLevelStateToResolve(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AlertsClient.UpdateSubscriptionLevelAlertStateToDismiss`

```go
ctx := context.TODO()
id := alerts.NewAlertID("12345678-1234-9876-4563-123456789012", "ascLocation", "alertName")

read, err := client.UpdateSubscriptionLevelAlertStateToDismiss(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AlertsClient.UpdateSubscriptionLevelAlertStateToReactivate`

```go
ctx := context.TODO()
id := alerts.NewAlertID("12345678-1234-9876-4563-123456789012", "ascLocation", "alertName")

read, err := client.UpdateSubscriptionLevelAlertStateToReactivate(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AlertsClient.UpdateSubscriptionLevelStateToResolve`

```go
ctx := context.TODO()
id := alerts.NewAlertID("12345678-1234-9876-4563-123456789012", "ascLocation", "alertName")

read, err := client.UpdateSubscriptionLevelStateToResolve(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
