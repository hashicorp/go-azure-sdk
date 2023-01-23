
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2021-12-01/restores` Documentation

The `restores` SDK allows for interaction with the Azure Resource Manager Service `recoveryservicesbackup` (API Version `2021-12-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2021-12-01/restores"
```


### Client Initialization

```go
client := restores.NewRestoresClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `RestoresClient.Trigger`

```go
ctx := context.TODO()
id := restores.NewRecoveryPointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "backupFabricValue", "protectionContainerValue", "protectedItemValue", "recoveryPointIdValue")

payload := restores.RestoreRequestResource{
	// ...
}


if err := client.TriggerThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
