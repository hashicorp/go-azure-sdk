
## `github.com/hashicorp/go-azure-sdk/resource-manager/databricks/2025-10-01-preview/put` Documentation

The `put` SDK allows for interaction with Azure Resource Manager `databricks` (API Version `2025-10-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/databricks/2025-10-01-preview/put"
```


### Client Initialization

```go
client := put.NewPUTClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PUTClient.PrivateEndpointConnectionsCreate`

```go
ctx := context.TODO()
id := put.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "privateEndpointConnectionName")

payload := put.PrivateEndpointConnection{
	// ...
}


if err := client.PrivateEndpointConnectionsCreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
