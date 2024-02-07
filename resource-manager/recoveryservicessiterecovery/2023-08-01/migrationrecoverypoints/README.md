
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-08-01/migrationrecoverypoints` Documentation

The `migrationrecoverypoints` SDK allows for interaction with the Azure Resource Manager Service `recoveryservicessiterecovery` (API Version `2023-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-08-01/migrationrecoverypoints"
```


### Client Initialization

```go
client := migrationrecoverypoints.NewMigrationRecoveryPointsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MigrationRecoveryPointsClient.Get`

```go
ctx := context.TODO()
id := migrationrecoverypoints.NewMigrationRecoveryPointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "replicationFabricValue", "replicationProtectionContainerValue", "replicationMigrationItemValue", "migrationRecoveryPointValue")

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
id := migrationrecoverypoints.NewReplicationMigrationItemID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "replicationFabricValue", "replicationProtectionContainerValue", "replicationMigrationItemValue")

// alternatively `client.ListByReplicationMigrationItems(ctx, id)` can be used to do batched pagination
items, err := client.ListByReplicationMigrationItemsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
