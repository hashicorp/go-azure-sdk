
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2023-02-01/itemlevelrecoveryconnections` Documentation

The `itemlevelrecoveryconnections` SDK allows for interaction with the Azure Resource Manager Service `recoveryservicesbackup` (API Version `2023-02-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2023-02-01/itemlevelrecoveryconnections"
```


### Client Initialization

```go
client := itemlevelrecoveryconnections.NewItemLevelRecoveryConnectionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ItemLevelRecoveryConnectionsClient.Provision`

```go
ctx := context.TODO()
id := itemlevelrecoveryconnections.NewRecoveryPointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "backupFabricValue", "protectionContainerValue", "protectedItemValue", "recoveryPointIdValue")

payload := itemlevelrecoveryconnections.ILRRequestResource{
	// ...
}


read, err := client.Provision(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ItemLevelRecoveryConnectionsClient.Revoke`

```go
ctx := context.TODO()
id := itemlevelrecoveryconnections.NewRecoveryPointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "backupFabricValue", "protectionContainerValue", "protectedItemValue", "recoveryPointIdValue")

read, err := client.Revoke(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
