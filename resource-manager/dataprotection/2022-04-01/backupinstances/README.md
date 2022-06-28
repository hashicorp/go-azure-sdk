
## `github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2022-04-01/backupinstances` Documentation

The `backupinstances` SDK allows for interaction with the Azure Resource Manager Service `dataprotection` (API Version `2022-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2022-04-01/backupinstances"
```


### Client Initialization

```go
client := backupinstances.NewBackupInstancesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `BackupInstancesClient.AdhocBackup`

```go
ctx := context.TODO()
id := backupinstances.NewBackupInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "backupInstanceValue")

payload := backupinstances.TriggerBackupRequest{
	// ...
}

future, err := client.AdhocBackup(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `BackupInstancesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := backupinstances.NewBackupInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "backupInstanceValue")

payload := backupinstances.BackupInstanceResource{
	// ...
}

future, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `BackupInstancesClient.Delete`

```go
ctx := context.TODO()
id := backupinstances.NewBackupInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "backupInstanceValue")
future, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `BackupInstancesClient.Get`

```go
ctx := context.TODO()
id := backupinstances.NewBackupInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "backupInstanceValue")
read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BackupInstancesClient.List`

```go
ctx := context.TODO()
id := backupinstances.NewBackupVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue")
// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BackupInstancesClient.ResumeBackups`

```go
ctx := context.TODO()
id := backupinstances.NewBackupInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "backupInstanceValue")
future, err := client.ResumeBackups(ctx, id)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `BackupInstancesClient.ResumeProtection`

```go
ctx := context.TODO()
id := backupinstances.NewBackupInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "backupInstanceValue")
future, err := client.ResumeProtection(ctx, id)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `BackupInstancesClient.StopProtection`

```go
ctx := context.TODO()
id := backupinstances.NewBackupInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "backupInstanceValue")
future, err := client.StopProtection(ctx, id)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `BackupInstancesClient.SuspendBackups`

```go
ctx := context.TODO()
id := backupinstances.NewBackupInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "backupInstanceValue")
future, err := client.SuspendBackups(ctx, id)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `BackupInstancesClient.SyncBackupInstance`

```go
ctx := context.TODO()
id := backupinstances.NewBackupInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "backupInstanceValue")

payload := backupinstances.SyncBackupInstanceRequest{
	// ...
}

future, err := client.SyncBackupInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `BackupInstancesClient.TriggerRehydrate`

```go
ctx := context.TODO()
id := backupinstances.NewBackupInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "backupInstanceValue")

payload := backupinstances.AzureBackupRehydrationRequest{
	// ...
}

future, err := client.TriggerRehydrate(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `BackupInstancesClient.TriggerRestore`

```go
ctx := context.TODO()
id := backupinstances.NewBackupInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "backupInstanceValue")

payload := backupinstances.AzureBackupRestoreRequest{
	// ...
}

future, err := client.TriggerRestore(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `BackupInstancesClient.ValidateForBackup`

```go
ctx := context.TODO()
id := backupinstances.NewBackupVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue")

payload := backupinstances.ValidateForBackupRequest{
	// ...
}

future, err := client.ValidateForBackup(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `BackupInstancesClient.ValidateForRestore`

```go
ctx := context.TODO()
id := backupinstances.NewBackupInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "backupInstanceValue")

payload := backupinstances.ValidateRestoreRequestObject{
	// ...
}

future, err := client.ValidateForRestore(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```
