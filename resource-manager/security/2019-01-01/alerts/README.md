
## `github.com/hashicorp/go-azure-sdk/resource-manager/security/2019-01-01/alerts` Documentation

The `alerts` SDK allows for interaction with the Azure Resource Manager Service `security` (API Version `2019-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/security/2019-01-01/alerts"
```


### Client Initialization

```go
client := alerts.NewAlertsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AlertsClient.GetResourceGroupLevelAlerts`

```go
ctx := context.TODO()
id := alerts.NewLocationAlertID("12345678-1234-9876-4563-123456789012", "example-resource-group", "locationValue", "alertValue")

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
id := alerts.NewAlertID("12345678-1234-9876-4563-123456789012", "locationValue", "alertValue")

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
id := alerts.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.List(ctx, id, alerts.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, alerts.DefaultListOperationOptions())
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
id := alerts.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListByResourceGroup(ctx, id, alerts.DefaultListByResourceGroupOperationOptions())` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id, alerts.DefaultListByResourceGroupOperationOptions())
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
id := alerts.NewProviderLocationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "locationValue")

// alternatively `client.ListResourceGroupLevelAlertsByRegion(ctx, id, alerts.DefaultListResourceGroupLevelAlertsByRegionOperationOptions())` can be used to do batched pagination
items, err := client.ListResourceGroupLevelAlertsByRegionComplete(ctx, id, alerts.DefaultListResourceGroupLevelAlertsByRegionOperationOptions())
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
id := alerts.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

// alternatively `client.ListSubscriptionLevelAlertsByRegion(ctx, id, alerts.DefaultListSubscriptionLevelAlertsByRegionOperationOptions())` can be used to do batched pagination
items, err := client.ListSubscriptionLevelAlertsByRegionComplete(ctx, id, alerts.DefaultListSubscriptionLevelAlertsByRegionOperationOptions())
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
id := alerts.NewLocationAlertID("12345678-1234-9876-4563-123456789012", "example-resource-group", "locationValue", "alertValue")

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
id := alerts.NewLocationAlertID("12345678-1234-9876-4563-123456789012", "example-resource-group", "locationValue", "alertValue")

read, err := client.UpdateResourceGroupLevelAlertStateToReactivate(ctx, id)
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
id := alerts.NewAlertID("12345678-1234-9876-4563-123456789012", "locationValue", "alertValue")

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
id := alerts.NewAlertID("12345678-1234-9876-4563-123456789012", "locationValue", "alertValue")

read, err := client.UpdateSubscriptionLevelAlertStateToReactivate(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
