
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2022-10-01/recoverypoint` Documentation

The `recoverypoint` SDK allows for interaction with the Azure Resource Manager Service `recoveryservicesbackup` (API Version `2022-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2022-10-01/recoverypoint"
```


### Client Initialization

```go
client := recoverypoint.NewRecoveryPointClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `RecoveryPointClient.MoveRecoveryPoint`

```go
ctx := context.TODO()
id := recoverypoint.NewRecoveryPointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "backupFabricValue", "protectionContainerValue", "protectedItemValue", "recoveryPointIdValue")

payload := recoverypoint.MoveRPAcrossTiersRequest{
	// ...
}


if err := client.MoveRecoveryPointThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
