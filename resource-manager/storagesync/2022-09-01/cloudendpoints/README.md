
## `github.com/hashicorp/go-azure-sdk/resource-manager/storagesync/2022-09-01/cloudendpoints` Documentation

The `cloudendpoints` SDK allows for interaction with Azure Resource Manager `storagesync` (API Version `2022-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/storagesync/2022-09-01/cloudendpoints"
```


### Client Initialization

```go
client := cloudendpoints.NewCloudEndpointsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CloudEndpointsClient.AfsShareMetadataCertificatePublicKeys`

```go
ctx := context.TODO()
id := cloudendpoints.NewCloudEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceName", "syncGroupName", "cloudEndpointName")

read, err := client.AfsShareMetadataCertificatePublicKeys(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudEndpointsClient.Create`

```go
ctx := context.TODO()
id := cloudendpoints.NewCloudEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceName", "syncGroupName", "cloudEndpointName")

payload := cloudendpoints.CloudEndpointCreateParameters{
	// ...
}


if err := client.CreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `CloudEndpointsClient.Delete`

```go
ctx := context.TODO()
id := cloudendpoints.NewCloudEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceName", "syncGroupName", "cloudEndpointName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `CloudEndpointsClient.Get`

```go
ctx := context.TODO()
id := cloudendpoints.NewCloudEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceName", "syncGroupName", "cloudEndpointName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudEndpointsClient.ListBySyncGroup`

```go
ctx := context.TODO()
id := cloudendpoints.NewSyncGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceName", "syncGroupName")

// alternatively `client.ListBySyncGroup(ctx, id)` can be used to do batched pagination
items, err := client.ListBySyncGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CloudEndpointsClient.PostBackup`

```go
ctx := context.TODO()
id := cloudendpoints.NewCloudEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceName", "syncGroupName", "cloudEndpointName")

payload := cloudendpoints.BackupRequest{
	// ...
}


if err := client.PostBackupThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `CloudEndpointsClient.PostRestore`

```go
ctx := context.TODO()
id := cloudendpoints.NewCloudEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceName", "syncGroupName", "cloudEndpointName")

payload := cloudendpoints.PostRestoreRequest{
	// ...
}


if err := client.PostRestoreThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `CloudEndpointsClient.PreBackup`

```go
ctx := context.TODO()
id := cloudendpoints.NewCloudEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceName", "syncGroupName", "cloudEndpointName")

payload := cloudendpoints.BackupRequest{
	// ...
}


if err := client.PreBackupThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `CloudEndpointsClient.PreRestore`

```go
ctx := context.TODO()
id := cloudendpoints.NewCloudEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceName", "syncGroupName", "cloudEndpointName")

payload := cloudendpoints.PreRestoreRequest{
	// ...
}


if err := client.PreRestoreThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `CloudEndpointsClient.Restoreheartbeat`

```go
ctx := context.TODO()
id := cloudendpoints.NewCloudEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceName", "syncGroupName", "cloudEndpointName")

read, err := client.Restoreheartbeat(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudEndpointsClient.TriggerChangeDetection`

```go
ctx := context.TODO()
id := cloudendpoints.NewCloudEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceName", "syncGroupName", "cloudEndpointName")

payload := cloudendpoints.TriggerChangeDetectionParameters{
	// ...
}


if err := client.TriggerChangeDetectionThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
