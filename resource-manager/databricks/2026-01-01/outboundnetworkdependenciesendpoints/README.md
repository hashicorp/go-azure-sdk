
## `github.com/hashicorp/go-azure-sdk/resource-manager/databricks/2026-01-01/outboundnetworkdependenciesendpoints` Documentation

The `outboundnetworkdependenciesendpoints` SDK allows for interaction with Azure Resource Manager `databricks` (API Version `2026-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/databricks/2026-01-01/outboundnetworkdependenciesendpoints"
```


### Client Initialization

```go
client := outboundnetworkdependenciesendpoints.NewOutboundNetworkDependenciesEndpointsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `OutboundNetworkDependenciesEndpointsClient.List`

```go
ctx := context.TODO()
id := outboundnetworkdependenciesendpoints.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName")

read, err := client.List(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
