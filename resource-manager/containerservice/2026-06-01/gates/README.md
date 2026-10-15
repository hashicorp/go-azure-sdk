
## `github.com/hashicorp/go-azure-sdk/resource-manager/containerservice/2026-06-01/gates` Documentation

The `gates` SDK allows for interaction with Azure Resource Manager `containerservice` (API Version `2026-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/containerservice/2026-06-01/gates"
```


### Client Initialization

```go
client := gates.NewGatesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GatesClient.Get`

```go
ctx := context.TODO()
id := gates.NewGateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "fleetName", "gateName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GatesClient.ListByFleet`

```go
ctx := context.TODO()
id := gates.NewFleetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "fleetName")

// alternatively `client.ListByFleet(ctx, id, gates.DefaultListByFleetOperationOptions())` can be used to do batched pagination
items, err := client.ListByFleetComplete(ctx, id, gates.DefaultListByFleetOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GatesClient.Update`

```go
ctx := context.TODO()
id := gates.NewGateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "fleetName", "gateName")

payload := gates.GatePatch{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload, gates.DefaultUpdateOperationOptions()); err != nil {
	// handle the error
}
```
