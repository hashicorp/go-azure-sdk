
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-06-01/replicationlogicalnetworks` Documentation

The `replicationlogicalnetworks` SDK allows for interaction with Azure Resource Manager `recoveryservicessiterecovery` (API Version `2023-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-06-01/replicationlogicalnetworks"
```


### Client Initialization

```go
client := replicationlogicalnetworks.NewReplicationLogicalNetworksClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ReplicationLogicalNetworksClient.Get`

```go
ctx := context.TODO()
id := replicationlogicalnetworks.NewReplicationLogicalNetworkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName", "fabricName", "logicalNetworkName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReplicationLogicalNetworksClient.ListByReplicationFabrics`

```go
ctx := context.TODO()
id := replicationlogicalnetworks.NewReplicationFabricID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName", "fabricName")

// alternatively `client.ListByReplicationFabrics(ctx, id)` can be used to do batched pagination
items, err := client.ListByReplicationFabricsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
