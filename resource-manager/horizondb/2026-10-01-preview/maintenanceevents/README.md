
## `github.com/hashicorp/go-azure-sdk/resource-manager/horizondb/2026-10-01-preview/maintenanceevents` Documentation

The `maintenanceevents` SDK allows for interaction with Azure Resource Manager `horizondb` (API Version `2026-10-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/horizondb/2026-10-01-preview/maintenanceevents"
```


### Client Initialization

```go
client := maintenanceevents.NewMaintenanceEventsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MaintenanceEventsClient.ApplyNow`

```go
ctx := context.TODO()
id := maintenanceevents.NewMaintenanceEventID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "poolName", "maintenanceEventId")

if err := client.ApplyNowThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `MaintenanceEventsClient.Get`

```go
ctx := context.TODO()
id := maintenanceevents.NewMaintenanceEventID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "poolName", "maintenanceEventId")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MaintenanceEventsClient.List`

```go
ctx := context.TODO()
id := maintenanceevents.NewPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "poolName")

// alternatively `client.List(ctx, id, maintenanceevents.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, maintenanceevents.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MaintenanceEventsClient.Reschedule`

```go
ctx := context.TODO()
id := maintenanceevents.NewMaintenanceEventID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "poolName", "maintenanceEventId")

payload := maintenanceevents.MaintenanceEventRescheduleRequest{
	// ...
}


if err := client.RescheduleThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
