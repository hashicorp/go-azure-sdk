
## `github.com/hashicorp/go-azure-sdk/resource-manager/horizondb/2026-10-01-preview/capabilities` Documentation

The `capabilities` SDK allows for interaction with Azure Resource Manager `horizondb` (API Version `2026-10-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/horizondb/2026-10-01-preview/capabilities"
```


### Client Initialization

```go
client := capabilities.NewCapabilitiesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CapabilitiesClient.Get`

```go
ctx := context.TODO()
id := capabilities.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CapabilitiesClient.List`

```go
ctx := context.TODO()
id := capabilities.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
