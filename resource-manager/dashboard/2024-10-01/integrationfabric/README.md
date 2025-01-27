
## `github.com/hashicorp/go-azure-sdk/resource-manager/dashboard/2024-10-01/integrationfabric` Documentation

The `integrationfabric` SDK allows for interaction with Azure Resource Manager `dashboard` (API Version `2024-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/dashboard/2024-10-01/integrationfabric"
```


### Client Initialization

```go
client := integrationfabric.NewIntegrationFabricClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `IntegrationFabricClient.Create`

```go
ctx := context.TODO()
id := integrationfabric.NewIntegrationFabricID("12345678-1234-9876-4563-123456789012", "example-resource-group", "grafanaName", "integrationFabricName")

payload := integrationfabric.IntegrationFabric{
	// ...
}


if err := client.CreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `IntegrationFabricClient.Delete`

```go
ctx := context.TODO()
id := integrationfabric.NewIntegrationFabricID("12345678-1234-9876-4563-123456789012", "example-resource-group", "grafanaName", "integrationFabricName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `IntegrationFabricClient.Get`

```go
ctx := context.TODO()
id := integrationfabric.NewIntegrationFabricID("12345678-1234-9876-4563-123456789012", "example-resource-group", "grafanaName", "integrationFabricName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IntegrationFabricClient.List`

```go
ctx := context.TODO()
id := integrationfabric.NewGrafanaID("12345678-1234-9876-4563-123456789012", "example-resource-group", "grafanaName")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `IntegrationFabricClient.Update`

```go
ctx := context.TODO()
id := integrationfabric.NewIntegrationFabricID("12345678-1234-9876-4563-123456789012", "example-resource-group", "grafanaName", "integrationFabricName")

payload := integrationfabric.IntegrationFabricUpdateParameters{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
