
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2024-01-01/recoverypoint` Documentation

The `recoverypoint` SDK allows for interaction with Azure Resource Manager `recoveryservicesbackup` (API Version `2024-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2024-01-01/recoverypoint"
```


### Client Initialization

```go
client := recoverypoint.NewRecoveryPointClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `RecoveryPointClient.MoveRecoveryPoint`

```go
ctx := context.TODO()
id := recoverypoint.NewRecoveryPointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName", "fabricName", "containerName", "protectedItemName", "recoveryPointId")

payload := recoverypoint.MoveRPAcrossTiersRequest{
	// ...
}


if err := client.MoveRecoveryPointThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
