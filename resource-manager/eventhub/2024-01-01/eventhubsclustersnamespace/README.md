
## `github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2024-01-01/eventhubsclustersnamespace` Documentation

The `eventhubsclustersnamespace` SDK allows for interaction with Azure Resource Manager `eventhub` (API Version `2024-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2024-01-01/eventhubsclustersnamespace"
```


### Client Initialization

```go
client := eventhubsclustersnamespace.NewEventHubsClustersNamespaceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EventHubsClustersNamespaceClient.ClustersListNamespaces`

```go
ctx := context.TODO()
id := eventhubsclustersnamespace.NewClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterValue")

read, err := client.ClustersListNamespaces(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
