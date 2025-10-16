
## `github.com/hashicorp/go-azure-sdk/resource-manager/dashboard/2025-08-01/integrationfabrics` Documentation

The `integrationfabrics` SDK allows for interaction with Azure Resource Manager `dashboard` (API Version `2025-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/dashboard/2025-08-01/integrationfabrics"
```


### Client Initialization

```go
client := integrationfabrics.NewIntegrationFabricsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `IntegrationFabricsClient.Create`

```go
ctx := context.TODO()
id := integrationfabrics.NewIntegrationFabricID("12345678-1234-9876-4563-123456789012", "example-resource-group", "grafanaName", "integrationFabricName")

payload := integrationfabrics.IntegrationFabric{
	// ...
}


if err := client.CreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `IntegrationFabricsClient.Delete`

```go
ctx := context.TODO()
id := integrationfabrics.NewIntegrationFabricID("12345678-1234-9876-4563-123456789012", "example-resource-group", "grafanaName", "integrationFabricName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `IntegrationFabricsClient.Get`

```go
ctx := context.TODO()
id := integrationfabrics.NewIntegrationFabricID("12345678-1234-9876-4563-123456789012", "example-resource-group", "grafanaName", "integrationFabricName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IntegrationFabricsClient.List`

```go
ctx := context.TODO()
id := integrationfabrics.NewGrafanaID("12345678-1234-9876-4563-123456789012", "example-resource-group", "grafanaName")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `IntegrationFabricsClient.Update`

```go
ctx := context.TODO()
id := integrationfabrics.NewIntegrationFabricID("12345678-1234-9876-4563-123456789012", "example-resource-group", "grafanaName", "integrationFabricName")

payload := integrationfabrics.IntegrationFabricUpdateParameters{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
