
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2023-01-15/recoverypointscrr` Documentation

The `recoverypointscrr` SDK allows for interaction with the Azure Resource Manager Service `recoveryservicesbackup` (API Version `2023-01-15`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2023-01-15/recoverypointscrr"
```


### Client Initialization

```go
client := recoverypointscrr.NewRecoveryPointsCrrClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `RecoveryPointsCrrClient.RecoveryPointsCrrGet`

```go
ctx := context.TODO()
id := recoverypointscrr.NewRecoveryPointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "backupFabricValue", "protectionContainerValue", "protectedItemValue", "recoveryPointIdValue")

read, err := client.RecoveryPointsCrrGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RecoveryPointsCrrClient.RecoveryPointsCrrList`

```go
ctx := context.TODO()
id := recoverypointscrr.NewProtectedItemID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "backupFabricValue", "protectionContainerValue", "protectedItemValue")

// alternatively `client.RecoveryPointsCrrList(ctx, id, recoverypointscrr.DefaultRecoveryPointsCrrListOperationOptions())` can be used to do batched pagination
items, err := client.RecoveryPointsCrrListComplete(ctx, id, recoverypointscrr.DefaultRecoveryPointsCrrListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
