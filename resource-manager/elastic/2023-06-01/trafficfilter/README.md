
## `github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2023-06-01/trafficfilter` Documentation

The `trafficfilter` SDK allows for interaction with Azure Resource Manager `elastic` (API Version `2023-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2023-06-01/trafficfilter"
```


### Client Initialization

```go
client := trafficfilter.NewTrafficFilterClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TrafficFilterClient.AllTrafficFilterslist`

```go
ctx := context.TODO()
id := trafficfilter.NewMonitorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName")

read, err := client.AllTrafficFilterslist(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TrafficFilterClient.AssociateTrafficFilterAssociate`

```go
ctx := context.TODO()
id := trafficfilter.NewMonitorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName")

if err := client.AssociateTrafficFilterAssociateThenPoll(ctx, id, trafficfilter.DefaultAssociateTrafficFilterAssociateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `TrafficFilterClient.CreateAndAssociateIPFilterCreate`

```go
ctx := context.TODO()
id := trafficfilter.NewMonitorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName")

if err := client.CreateAndAssociateIPFilterCreateThenPoll(ctx, id, trafficfilter.DefaultCreateAndAssociateIPFilterCreateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `TrafficFilterClient.CreateAndAssociatePLFilterCreate`

```go
ctx := context.TODO()
id := trafficfilter.NewMonitorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName")

if err := client.CreateAndAssociatePLFilterCreateThenPoll(ctx, id, trafficfilter.DefaultCreateAndAssociatePLFilterCreateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `TrafficFilterClient.Delete`

```go
ctx := context.TODO()
id := trafficfilter.NewMonitorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName")

read, err := client.Delete(ctx, id, trafficfilter.DefaultDeleteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TrafficFilterClient.DetachAndDeleteTrafficFilterDelete`

```go
ctx := context.TODO()
id := trafficfilter.NewMonitorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName")

read, err := client.DetachAndDeleteTrafficFilterDelete(ctx, id, trafficfilter.DefaultDetachAndDeleteTrafficFilterDeleteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TrafficFilterClient.DetachTrafficFilterUpdate`

```go
ctx := context.TODO()
id := trafficfilter.NewMonitorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName")

if err := client.DetachTrafficFilterUpdateThenPoll(ctx, id, trafficfilter.DefaultDetachTrafficFilterUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `TrafficFilterClient.ListAssociatedTrafficFilterslist`

```go
ctx := context.TODO()
id := trafficfilter.NewMonitorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName")

read, err := client.ListAssociatedTrafficFilterslist(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
