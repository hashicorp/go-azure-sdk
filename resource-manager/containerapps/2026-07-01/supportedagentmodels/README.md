
## `github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2026-07-01/supportedagentmodels` Documentation

The `supportedagentmodels` SDK allows for interaction with Azure Resource Manager `containerapps` (API Version `2026-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2026-07-01/supportedagentmodels"
```


### Client Initialization

```go
client := supportedagentmodels.NewSupportedAgentModelsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SupportedAgentModelsClient.ListByLocation`

```go
ctx := context.TODO()
id := supportedagentmodels.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

// alternatively `client.ListByLocation(ctx, id)` can be used to do batched pagination
items, err := client.ListByLocationComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
