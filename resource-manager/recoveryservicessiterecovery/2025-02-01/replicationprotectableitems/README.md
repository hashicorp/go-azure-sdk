
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2025-02-01/replicationprotectableitems` Documentation

The `replicationprotectableitems` SDK allows for interaction with Azure Resource Manager `recoveryservicessiterecovery` (API Version `2025-02-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2025-02-01/replicationprotectableitems"
```


### Client Initialization

```go
client := replicationprotectableitems.NewReplicationProtectableItemsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ReplicationProtectableItemsClient.Get`

```go
ctx := context.TODO()
id := replicationprotectableitems.NewReplicationProtectableItemID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName", "replicationFabricName", "replicationProtectionContainerName", "replicationProtectableItemName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReplicationProtectableItemsClient.ListByReplicationProtectionContainers`

```go
ctx := context.TODO()
id := replicationprotectableitems.NewReplicationProtectionContainerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName", "replicationFabricName", "replicationProtectionContainerName")

// alternatively `client.ListByReplicationProtectionContainers(ctx, id, replicationprotectableitems.DefaultListByReplicationProtectionContainersOperationOptions())` can be used to do batched pagination
items, err := client.ListByReplicationProtectionContainersComplete(ctx, id, replicationprotectableitems.DefaultListByReplicationProtectionContainersOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
