
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2024-04-01/clusterrecoverypoints` Documentation

The `clusterrecoverypoints` SDK allows for interaction with Azure Resource Manager `recoveryservicessiterecovery` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2024-04-01/clusterrecoverypoints"
```


### Client Initialization

```go
client := clusterrecoverypoints.NewClusterRecoveryPointsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ClusterRecoveryPointsClient.ListByReplicationProtectionCluster`

```go
ctx := context.TODO()
id := clusterrecoverypoints.NewReplicationProtectionClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "replicationFabricValue", "replicationProtectionContainerValue", "replicationProtectionClusterValue")

// alternatively `client.ListByReplicationProtectionCluster(ctx, id)` can be used to do batched pagination
items, err := client.ListByReplicationProtectionClusterComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
