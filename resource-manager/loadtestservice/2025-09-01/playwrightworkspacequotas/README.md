
## `github.com/hashicorp/go-azure-sdk/resource-manager/loadtestservice/2025-09-01/playwrightworkspacequotas` Documentation

The `playwrightworkspacequotas` SDK allows for interaction with Azure Resource Manager `loadtestservice` (API Version `2025-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/loadtestservice/2025-09-01/playwrightworkspacequotas"
```


### Client Initialization

```go
client := playwrightworkspacequotas.NewPlaywrightWorkspaceQuotasClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PlaywrightWorkspaceQuotasClient.Get`

```go
ctx := context.TODO()
id := playwrightworkspacequotas.NewPlaywrightWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "playwrightWorkspaceName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PlaywrightWorkspaceQuotasClient.ListByPlaywrightWorkspace`

```go
ctx := context.TODO()
id := playwrightworkspacequotas.NewPlaywrightWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "playwrightWorkspaceName")

// alternatively `client.ListByPlaywrightWorkspace(ctx, id)` can be used to do batched pagination
items, err := client.ListByPlaywrightWorkspaceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
