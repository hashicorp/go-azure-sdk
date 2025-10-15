
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-08-01/protectionintentresources` Documentation

The `protectionintentresources` SDK allows for interaction with Azure Resource Manager `recoveryservicesbackup` (API Version `2025-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-08-01/protectionintentresources"
```


### Client Initialization

```go
client := protectionintentresources.NewProtectionIntentResourcesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ProtectionIntentResourcesClient.ProtectionIntentCreateOrUpdate`

```go
ctx := context.TODO()
id := protectionintentresources.NewBackupProtectionIntentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName", "backupFabricName", "backupProtectionIntentName")

payload := protectionintentresources.ProtectionIntentResource{
	// ...
}


read, err := client.ProtectionIntentCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ProtectionIntentResourcesClient.ProtectionIntentDelete`

```go
ctx := context.TODO()
id := protectionintentresources.NewBackupProtectionIntentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName", "backupFabricName", "backupProtectionIntentName")

read, err := client.ProtectionIntentDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ProtectionIntentResourcesClient.ProtectionIntentGet`

```go
ctx := context.TODO()
id := protectionintentresources.NewBackupProtectionIntentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName", "backupFabricName", "backupProtectionIntentName")

read, err := client.ProtectionIntentGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
