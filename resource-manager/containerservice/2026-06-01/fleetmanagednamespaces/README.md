
## `github.com/hashicorp/go-azure-sdk/resource-manager/containerservice/2026-06-01/fleetmanagednamespaces` Documentation

The `fleetmanagednamespaces` SDK allows for interaction with Azure Resource Manager `containerservice` (API Version `2026-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/containerservice/2026-06-01/fleetmanagednamespaces"
```


### Client Initialization

```go
client := fleetmanagednamespaces.NewFleetManagedNamespacesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `FleetManagedNamespacesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := fleetmanagednamespaces.NewManagedNamespaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "fleetName", "managedNamespaceName")

payload := fleetmanagednamespaces.FleetManagedNamespace{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload, fleetmanagednamespaces.DefaultCreateOrUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `FleetManagedNamespacesClient.Delete`

```go
ctx := context.TODO()
id := fleetmanagednamespaces.NewManagedNamespaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "fleetName", "managedNamespaceName")

if err := client.DeleteThenPoll(ctx, id, fleetmanagednamespaces.DefaultDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `FleetManagedNamespacesClient.Get`

```go
ctx := context.TODO()
id := fleetmanagednamespaces.NewManagedNamespaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "fleetName", "managedNamespaceName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `FleetManagedNamespacesClient.ListByFleet`

```go
ctx := context.TODO()
id := fleetmanagednamespaces.NewFleetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "fleetName")

// alternatively `client.ListByFleet(ctx, id)` can be used to do batched pagination
items, err := client.ListByFleetComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `FleetManagedNamespacesClient.Update`

```go
ctx := context.TODO()
id := fleetmanagednamespaces.NewManagedNamespaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "fleetName", "managedNamespaceName")

payload := fleetmanagednamespaces.FleetManagedNamespacePatch{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload, fleetmanagednamespaces.DefaultUpdateOperationOptions()); err != nil {
	// handle the error
}
```
