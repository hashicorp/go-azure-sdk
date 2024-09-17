
## `github.com/hashicorp/go-azure-sdk/resource-manager/storagesync/2022-09-01/actions` Documentation

The `actions` SDK allows for interaction with Azure Resource Manager `storagesync` (API Version `2022-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/storagesync/2022-09-01/actions"
```


### Client Initialization

```go
client := actions.NewActionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ActionsClient.CloudEndpointsAfsShareMetadataCertificatePublicKeys`

```go
ctx := context.TODO()
id := actions.NewCloudEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceValue", "syncGroupValue", "cloudEndpointValue")

read, err := client.CloudEndpointsAfsShareMetadataCertificatePublicKeys(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ActionsClient.CloudEndpointsPostBackup`

```go
ctx := context.TODO()
id := actions.NewCloudEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceValue", "syncGroupValue", "cloudEndpointValue")

payload := actions.BackupRequest{
	// ...
}


if err := client.CloudEndpointsPostBackupThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ActionsClient.CloudEndpointsPostRestore`

```go
ctx := context.TODO()
id := actions.NewCloudEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceValue", "syncGroupValue", "cloudEndpointValue")

payload := actions.PostRestoreRequest{
	// ...
}


if err := client.CloudEndpointsPostRestoreThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ActionsClient.CloudEndpointsPreBackup`

```go
ctx := context.TODO()
id := actions.NewCloudEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceValue", "syncGroupValue", "cloudEndpointValue")

payload := actions.BackupRequest{
	// ...
}


if err := client.CloudEndpointsPreBackupThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ActionsClient.CloudEndpointsPreRestore`

```go
ctx := context.TODO()
id := actions.NewCloudEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceValue", "syncGroupValue", "cloudEndpointValue")

payload := actions.PreRestoreRequest{
	// ...
}


if err := client.CloudEndpointsPreRestoreThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ActionsClient.CloudEndpointsTriggerChangeDetection`

```go
ctx := context.TODO()
id := actions.NewCloudEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceValue", "syncGroupValue", "cloudEndpointValue")

payload := actions.TriggerChangeDetectionParameters{
	// ...
}


if err := client.CloudEndpointsTriggerChangeDetectionThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ActionsClient.CloudEndpointsrestoreheartbeat`

```go
ctx := context.TODO()
id := actions.NewCloudEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceValue", "syncGroupValue", "cloudEndpointValue")

read, err := client.CloudEndpointsrestoreheartbeat(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ActionsClient.RegisteredServerstriggerRollover`

```go
ctx := context.TODO()
id := actions.NewRegisteredServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceValue", "serverIdValue")

payload := actions.TriggerRolloverRequest{
	// ...
}


if err := client.RegisteredServerstriggerRolloverThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ActionsClient.ServerEndpointsrecallAction`

```go
ctx := context.TODO()
id := actions.NewServerEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceValue", "syncGroupValue", "serverEndpointValue")

payload := actions.RecallActionParameters{
	// ...
}


if err := client.ServerEndpointsrecallActionThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ActionsClient.WorkflowsAbort`

```go
ctx := context.TODO()
id := actions.NewWorkflowID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceValue", "workflowIdValue")

read, err := client.WorkflowsAbort(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
