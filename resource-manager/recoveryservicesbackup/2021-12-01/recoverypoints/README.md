
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2021-12-01/recoverypoints` Documentation

The `recoverypoints` SDK allows for interaction with the Azure Resource Manager Service `recoveryservicesbackup` (API Version `2021-12-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2021-12-01/recoverypoints"
```


### Client Initialization

```go
client := recoverypoints.NewRecoveryPointsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `RecoveryPointsClient.Get`

```go
ctx := context.TODO()
id := recoverypoints.NewRecoveryPointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "backupFabricValue", "protectionContainerValue", "protectedItemValue", "recoveryPointIdValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RecoveryPointsClient.List`

```go
ctx := context.TODO()
id := recoverypoints.NewProtectedItemID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "backupFabricValue", "protectionContainerValue", "protectedItemValue")

// alternatively `client.List(ctx, id, recoverypoints.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, recoverypoints.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
