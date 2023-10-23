
## `github.com/hashicorp/go-azure-sdk/resource-manager/storagesync/2022-09-01/backuprestore` Documentation

The `backuprestore` SDK allows for interaction with the Azure Resource Manager Service `storagesync` (API Version `2022-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/storagesync/2022-09-01/backuprestore"
```


### Client Initialization

```go
client := backuprestore.NewBackupRestoreClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BackupRestoreClient.CloudEndpointsAfsShareMetadataCertificatePublicKeys`

```go
ctx := context.TODO()
id := backuprestore.NewCloudEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceValue", "syncGroupValue", "cloudEndpointValue")

read, err := client.CloudEndpointsAfsShareMetadataCertificatePublicKeys(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BackupRestoreClient.CloudEndpointsPostBackup`

```go
ctx := context.TODO()
id := backuprestore.NewCloudEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceValue", "syncGroupValue", "cloudEndpointValue")

payload := backuprestore.BackupRequest{
	// ...
}


if err := client.CloudEndpointsPostBackupThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BackupRestoreClient.CloudEndpointsPostRestore`

```go
ctx := context.TODO()
id := backuprestore.NewCloudEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceValue", "syncGroupValue", "cloudEndpointValue")

payload := backuprestore.PostRestoreRequest{
	// ...
}


if err := client.CloudEndpointsPostRestoreThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BackupRestoreClient.CloudEndpointsPreBackup`

```go
ctx := context.TODO()
id := backuprestore.NewCloudEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceValue", "syncGroupValue", "cloudEndpointValue")

payload := backuprestore.BackupRequest{
	// ...
}


if err := client.CloudEndpointsPreBackupThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BackupRestoreClient.CloudEndpointsPreRestore`

```go
ctx := context.TODO()
id := backuprestore.NewCloudEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceValue", "syncGroupValue", "cloudEndpointValue")

payload := backuprestore.PreRestoreRequest{
	// ...
}


if err := client.CloudEndpointsPreRestoreThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BackupRestoreClient.CloudEndpointsrestoreheartbeat`

```go
ctx := context.TODO()
id := backuprestore.NewCloudEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceValue", "syncGroupValue", "cloudEndpointValue")

read, err := client.CloudEndpointsrestoreheartbeat(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
