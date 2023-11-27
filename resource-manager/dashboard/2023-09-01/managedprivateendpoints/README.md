
## `github.com/hashicorp/go-azure-sdk/resource-manager/dashboard/2023-09-01/managedprivateendpoints` Documentation

The `managedprivateendpoints` SDK allows for interaction with the Azure Resource Manager Service `dashboard` (API Version `2023-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/dashboard/2023-09-01/managedprivateendpoints"
```


### Client Initialization

```go
client := managedprivateendpoints.NewManagedPrivateEndpointsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ManagedPrivateEndpointsClient.Create`

```go
ctx := context.TODO()
id := managedprivateendpoints.NewManagedPrivateEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "grafanaValue", "managedPrivateEndpointValue")

payload := managedprivateendpoints.ManagedPrivateEndpointModel{
	// ...
}


if err := client.CreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ManagedPrivateEndpointsClient.Delete`

```go
ctx := context.TODO()
id := managedprivateendpoints.NewManagedPrivateEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "grafanaValue", "managedPrivateEndpointValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ManagedPrivateEndpointsClient.Get`

```go
ctx := context.TODO()
id := managedprivateendpoints.NewManagedPrivateEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "grafanaValue", "managedPrivateEndpointValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedPrivateEndpointsClient.List`

```go
ctx := context.TODO()
id := managedprivateendpoints.NewGrafanaID("12345678-1234-9876-4563-123456789012", "example-resource-group", "grafanaValue")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ManagedPrivateEndpointsClient.Refresh`

```go
ctx := context.TODO()
id := managedprivateendpoints.NewGrafanaID("12345678-1234-9876-4563-123456789012", "example-resource-group", "grafanaValue")

if err := client.RefreshThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ManagedPrivateEndpointsClient.Update`

```go
ctx := context.TODO()
id := managedprivateendpoints.NewManagedPrivateEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "grafanaValue", "managedPrivateEndpointValue")

payload := managedprivateendpoints.ManagedPrivateEndpointUpdateParameters{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
