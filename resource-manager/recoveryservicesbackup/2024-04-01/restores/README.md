
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2024-04-01/restores` Documentation

The `restores` SDK allows for interaction with Azure Resource Manager `recoveryservicesbackup` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2024-04-01/restores"
```


### Client Initialization

```go
client := restores.NewRestoresClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `RestoresClient.Trigger`

```go
ctx := context.TODO()
id := restores.NewRecoveryPointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName", "backupFabricName", "protectionContainerName", "protectedItemName", "recoveryPointId")

payload := restores.RestoreRequestResource{
	// ...
}


if err := client.TriggerThenPoll(ctx, id, payload, restores.DefaultTriggerOperationOptions()); err != nil {
	// handle the error
}
```
