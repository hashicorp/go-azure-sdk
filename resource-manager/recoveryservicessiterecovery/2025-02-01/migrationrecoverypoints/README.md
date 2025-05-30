
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2025-02-01/migrationrecoverypoints` Documentation

The `migrationrecoverypoints` SDK allows for interaction with Azure Resource Manager `recoveryservicessiterecovery` (API Version `2025-02-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2025-02-01/migrationrecoverypoints"
```


### Client Initialization

```go
client := migrationrecoverypoints.NewMigrationRecoveryPointsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MigrationRecoveryPointsClient.Get`

```go
ctx := context.TODO()
id := migrationrecoverypoints.NewMigrationRecoveryPointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName", "replicationFabricName", "replicationProtectionContainerName", "replicationMigrationItemName", "migrationRecoveryPointName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MigrationRecoveryPointsClient.ListByReplicationMigrationItems`

```go
ctx := context.TODO()
id := migrationrecoverypoints.NewReplicationMigrationItemID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName", "replicationFabricName", "replicationProtectionContainerName", "replicationMigrationItemName")

// alternatively `client.ListByReplicationMigrationItems(ctx, id)` can be used to do batched pagination
items, err := client.ListByReplicationMigrationItemsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
