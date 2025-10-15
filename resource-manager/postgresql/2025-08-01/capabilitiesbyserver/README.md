
## `github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2025-08-01/capabilitiesbyserver` Documentation

The `capabilitiesbyserver` SDK allows for interaction with Azure Resource Manager `postgresql` (API Version `2025-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2025-08-01/capabilitiesbyserver"
```


### Client Initialization

```go
client := capabilitiesbyserver.NewCapabilitiesByServerClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CapabilitiesByServerClient.List`

```go
ctx := context.TODO()
id := capabilitiesbyserver.NewFlexibleServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "flexibleServerName")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
