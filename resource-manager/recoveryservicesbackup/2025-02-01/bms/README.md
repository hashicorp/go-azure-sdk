
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-02-01/bms` Documentation

The `bms` SDK allows for interaction with Azure Resource Manager `recoveryservicesbackup` (API Version `2025-02-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-02-01/bms"
```


### Client Initialization

```go
client := bms.NewBmsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BmsClient.BackupOperationStatusesGet`

```go
ctx := context.TODO()
id := bms.NewBackupOperationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName", "operationId")

read, err := client.BackupOperationStatusesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BmsClient.BackupProtectableItemsList`

```go
ctx := context.TODO()
id := bms.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName")

// alternatively `client.BackupProtectableItemsList(ctx, id, bms.DefaultBackupProtectableItemsListOperationOptions())` can be used to do batched pagination
items, err := client.BackupProtectableItemsListComplete(ctx, id, bms.DefaultBackupProtectableItemsListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BmsClient.BackupProtectedItemsList`

```go
ctx := context.TODO()
id := bms.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName")

// alternatively `client.BackupProtectedItemsList(ctx, id, bms.DefaultBackupProtectedItemsListOperationOptions())` can be used to do batched pagination
items, err := client.BackupProtectedItemsListComplete(ctx, id, bms.DefaultBackupProtectedItemsListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BmsClient.BackupProtectionContainersList`

```go
ctx := context.TODO()
id := bms.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName")

// alternatively `client.BackupProtectionContainersList(ctx, id, bms.DefaultBackupProtectionContainersListOperationOptions())` can be used to do batched pagination
items, err := client.BackupProtectionContainersListComplete(ctx, id, bms.DefaultBackupProtectionContainersListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BmsClient.BackupProtectionIntentList`

```go
ctx := context.TODO()
id := bms.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName")

// alternatively `client.BackupProtectionIntentList(ctx, id, bms.DefaultBackupProtectionIntentListOperationOptions())` can be used to do batched pagination
items, err := client.BackupProtectionIntentListComplete(ctx, id, bms.DefaultBackupProtectionIntentListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BmsClient.BackupStatusGet`

```go
ctx := context.TODO()
id := bms.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

payload := bms.BackupStatusRequest{
	// ...
}


read, err := client.BackupStatusGet(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BmsClient.BackupUsageSummariesList`

```go
ctx := context.TODO()
id := bms.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName")

// alternatively `client.BackupUsageSummariesList(ctx, id, bms.DefaultBackupUsageSummariesListOperationOptions())` can be used to do batched pagination
items, err := client.BackupUsageSummariesListComplete(ctx, id, bms.DefaultBackupUsageSummariesListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BmsClient.DeletedProtectionContainersList`

```go
ctx := context.TODO()
id := bms.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName")

// alternatively `client.DeletedProtectionContainersList(ctx, id, bms.DefaultDeletedProtectionContainersListOperationOptions())` can be used to do batched pagination
items, err := client.DeletedProtectionContainersListComplete(ctx, id, bms.DefaultDeletedProtectionContainersListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BmsClient.FeatureSupportValidate`

```go
ctx := context.TODO()
id := bms.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

payload := bms.FeatureSupportRequest{
	// ...
}


read, err := client.FeatureSupportValidate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BmsClient.FetchTieringCostPost`

```go
ctx := context.TODO()
id := bms.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName")

payload := bms.FetchTieringCostInfoRequest{
	// ...
}


if err := client.FetchTieringCostPostThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BmsClient.JobsExport`

```go
ctx := context.TODO()
id := bms.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName")

read, err := client.JobsExport(ctx, id, bms.DefaultJobsExportOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BmsClient.OperationValidate`

```go
ctx := context.TODO()
id := bms.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName")

payload := bms.ValidateOperationRequestResource{
	// ...
}


read, err := client.OperationValidate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BmsClient.ProtectableContainersList`

```go
ctx := context.TODO()
id := bms.NewBackupFabricID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName", "backupFabricName")

// alternatively `client.ProtectableContainersList(ctx, id, bms.DefaultProtectableContainersListOperationOptions())` can be used to do batched pagination
items, err := client.ProtectableContainersListComplete(ctx, id, bms.DefaultProtectableContainersListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BmsClient.ProtectionContainersRefresh`

```go
ctx := context.TODO()
id := bms.NewBackupFabricID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName", "backupFabricName")

read, err := client.ProtectionContainersRefresh(ctx, id, bms.DefaultProtectionContainersRefreshOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BmsClient.ProtectionIntentValidate`

```go
ctx := context.TODO()
id := bms.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

payload := bms.PreValidateEnableBackupRequest{
	// ...
}


read, err := client.ProtectionIntentValidate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BmsClient.SecurityPINsGet`

```go
ctx := context.TODO()
id := bms.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName")

payload := bms.SecurityPinBase{
	// ...
}


read, err := client.SecurityPINsGet(ctx, id, payload, bms.DefaultSecurityPINsGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BmsClient.ValidateOperationStatusesGet`

```go
ctx := context.TODO()
id := bms.NewBackupValidateOperationsStatusID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName", "operationId")

read, err := client.ValidateOperationStatusesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BmsClient.ValidateOperationTrigger`

```go
ctx := context.TODO()
id := bms.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName")

payload := bms.ValidateOperationRequestResource{
	// ...
}


if err := client.ValidateOperationTriggerThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
