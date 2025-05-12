
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2025-01-01/clusterrecoverypoint` Documentation

The `clusterrecoverypoint` SDK allows for interaction with Azure Resource Manager `recoveryservicessiterecovery` (API Version `2025-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2025-01-01/clusterrecoverypoint"
```


### Client Initialization

```go
client := clusterrecoverypoint.NewClusterRecoveryPointClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ClusterRecoveryPointClient.Get`

```go
ctx := context.TODO()
id := clusterrecoverypoint.NewReplicationProtectionClusterRecoveryPointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName", "replicationFabricName", "replicationProtectionContainerName", "replicationProtectionClusterName", "recoveryPointName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
