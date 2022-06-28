
## `github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2018-01-01-preview/eventhubsclustersnamespace` Documentation

The `eventhubsclustersnamespace` SDK allows for interaction with the Azure Resource Manager Service `eventhub` (API Version `2018-01-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2018-01-01-preview/eventhubsclustersnamespace"
```


### Client Initialization

```go
client := eventhubsclustersnamespace.NewEventHubsClustersNamespaceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
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
