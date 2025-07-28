
## `github.com/hashicorp/go-azure-sdk/resource-manager/dynatrace/2024-04-24/monitoredsubscriptions` Documentation

The `monitoredsubscriptions` SDK allows for interaction with Azure Resource Manager `dynatrace` (API Version `2024-04-24`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/dynatrace/2024-04-24/monitoredsubscriptions"
```


### Client Initialization

```go
client := monitoredsubscriptions.NewMonitoredSubscriptionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MonitoredSubscriptionsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := monitoredsubscriptions.NewMonitorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName")

payload := monitoredsubscriptions.MonitoredSubscriptionProperties{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `MonitoredSubscriptionsClient.Delete`

```go
ctx := context.TODO()
id := monitoredsubscriptions.NewMonitorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `MonitoredSubscriptionsClient.Get`

```go
ctx := context.TODO()
id := monitoredsubscriptions.NewMonitorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MonitoredSubscriptionsClient.List`

```go
ctx := context.TODO()
id := monitoredsubscriptions.NewMonitorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MonitoredSubscriptionsClient.Update`

```go
ctx := context.TODO()
id := monitoredsubscriptions.NewMonitorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName")

payload := monitoredsubscriptions.MonitoredSubscriptionProperties{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
