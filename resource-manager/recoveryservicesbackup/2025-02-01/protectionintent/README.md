
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-02-01/protectionintent` Documentation

The `protectionintent` SDK allows for interaction with Azure Resource Manager `recoveryservicesbackup` (API Version `2025-02-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-02-01/protectionintent"
```


### Client Initialization

```go
client := protectionintent.NewProtectionIntentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ProtectionIntentClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := protectionintent.NewBackupProtectionIntentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName", "backupFabricName", "backupProtectionIntentName")

payload := protectionintent.ProtectionIntentResource{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ProtectionIntentClient.Delete`

```go
ctx := context.TODO()
id := protectionintent.NewBackupProtectionIntentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName", "backupFabricName", "backupProtectionIntentName")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ProtectionIntentClient.Get`

```go
ctx := context.TODO()
id := protectionintent.NewBackupProtectionIntentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName", "backupFabricName", "backupProtectionIntentName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ProtectionIntentClient.Validate`

```go
ctx := context.TODO()
id := protectionintent.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

payload := protectionintent.PreValidateEnableBackupRequest{
	// ...
}


read, err := client.Validate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
