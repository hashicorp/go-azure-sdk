
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/webapps` Documentation

The `webapps` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/webapps"
```


### Client Initialization

```go
client := webapps.NewWebAppsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WebAppsClient.AnalyzeCustomHostnameSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

read, err := client.AnalyzeCustomHostnameSlot(ctx, id, webapps.DefaultAnalyzeCustomHostnameSlotOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebAppsClient.ApplySlotConfigurationSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

payload := webapps.CsmSlotEntity{
	// ...
}


read, err := client.ApplySlotConfigurationSlot(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebAppsClient.BackupSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

payload := webapps.BackupRequest{
	// ...
}


read, err := client.BackupSlot(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebAppsClient.CreateOrUpdateHostSecretSlot`

```go
ctx := context.TODO()
id := webapps.NewHostDefaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "defaultName", "keyName")

payload := webapps.KeyInfo{
	// ...
}


read, err := client.CreateOrUpdateHostSecretSlot(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebAppsClient.CreateOrUpdateSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

payload := webapps.Site{
	// ...
}


if err := client.CreateOrUpdateSlotThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `WebAppsClient.DeleteBackupConfigurationSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

read, err := client.DeleteBackupConfigurationSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebAppsClient.DeleteHostSecretSlot`

```go
ctx := context.TODO()
id := webapps.NewHostDefaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "defaultName", "keyName")

read, err := client.DeleteHostSecretSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebAppsClient.DeleteSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

read, err := client.DeleteSlot(ctx, id, webapps.DefaultDeleteSlotOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebAppsClient.DeployWorkflowArtifactsSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

payload := webapps.WorkflowArtifacts{
	// ...
}


read, err := client.DeployWorkflowArtifactsSlot(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebAppsClient.DiscoverBackupSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

payload := webapps.RestoreRequest{
	// ...
}


read, err := client.DiscoverBackupSlot(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebAppsClient.GenerateNewSitePublishingPasswordSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

read, err := client.GenerateNewSitePublishingPasswordSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebAppsClient.GetAuthSettingsSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

read, err := client.GetAuthSettingsSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebAppsClient.GetBackupConfigurationSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

read, err := client.GetBackupConfigurationSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebAppsClient.GetContainerLogsZipSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

read, err := client.GetContainerLogsZipSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebAppsClient.GetFunctionsAdminTokenSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

read, err := client.GetFunctionsAdminTokenSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebAppsClient.GetNetworkTracesSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotNetworkTraceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "operationId")

read, err := client.GetNetworkTracesSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebAppsClient.GetNetworkTracesSlotV2`

```go
ctx := context.TODO()
id := webapps.NewSiteSlotNetworkTraceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "operationId")

read, err := client.GetNetworkTracesSlotV2(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebAppsClient.GetPrivateLinkResourcesSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

read, err := client.GetPrivateLinkResourcesSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebAppsClient.GetSitePhpErrorLogFlagSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

read, err := client.GetSitePhpErrorLogFlagSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebAppsClient.GetSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

read, err := client.GetSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebAppsClient.GetWebSiteContainerLogsSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

read, err := client.GetWebSiteContainerLogsSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebAppsClient.IsCloneableSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

read, err := client.IsCloneableSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebAppsClient.ListApplicationSettingsSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

read, err := client.ListApplicationSettingsSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebAppsClient.ListAzureStorageAccountsSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

read, err := client.ListAzureStorageAccountsSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebAppsClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListByResourceGroup(ctx, id, webapps.DefaultListByResourceGroupOperationOptions())` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id, webapps.DefaultListByResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WebAppsClient.ListConnectionStringsSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

read, err := client.ListConnectionStringsSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebAppsClient.ListHostKeysSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

read, err := client.ListHostKeysSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebAppsClient.ListHybridConnectionsSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

read, err := client.ListHybridConnectionsSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebAppsClient.ListMetadataSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

read, err := client.ListMetadataSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebAppsClient.ListPerfMonCountersSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

// alternatively `client.ListPerfMonCountersSlot(ctx, id, webapps.DefaultListPerfMonCountersSlotOperationOptions())` can be used to do batched pagination
items, err := client.ListPerfMonCountersSlotComplete(ctx, id, webapps.DefaultListPerfMonCountersSlotOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WebAppsClient.ListPremierAddOnsSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

read, err := client.ListPremierAddOnsSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebAppsClient.ListPublishingCredentialsSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

if err := client.ListPublishingCredentialsSlotThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `WebAppsClient.ListPublishingProfileXmlWithSecretsSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

payload := webapps.CsmPublishingProfileOptions{
	// ...
}


read, err := client.ListPublishingProfileXmlWithSecretsSlot(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebAppsClient.ListRelayServiceConnectionsSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

read, err := client.ListRelayServiceConnectionsSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebAppsClient.ListSiteBackupsSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

// alternatively `client.ListSiteBackupsSlot(ctx, id)` can be used to do batched pagination
items, err := client.ListSiteBackupsSlotComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WebAppsClient.ListSitePushSettingsSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

read, err := client.ListSitePushSettingsSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebAppsClient.ListSlotDifferencesSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

payload := webapps.CsmSlotEntity{
	// ...
}


// alternatively `client.ListSlotDifferencesSlot(ctx, id, payload)` can be used to do batched pagination
items, err := client.ListSlotDifferencesSlotComplete(ctx, id, payload)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WebAppsClient.ListSlots`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

// alternatively `client.ListSlots(ctx, id)` can be used to do batched pagination
items, err := client.ListSlotsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WebAppsClient.ListSnapshotsFromDRSecondarySlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

// alternatively `client.ListSnapshotsFromDRSecondarySlot(ctx, id)` can be used to do batched pagination
items, err := client.ListSnapshotsFromDRSecondarySlotComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WebAppsClient.ListSnapshotsSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

// alternatively `client.ListSnapshotsSlot(ctx, id)` can be used to do batched pagination
items, err := client.ListSnapshotsSlotComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WebAppsClient.ListSyncFunctionTriggersSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

read, err := client.ListSyncFunctionTriggersSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebAppsClient.ListSyncStatusSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

read, err := client.ListSyncStatusSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebAppsClient.ListUsagesSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

// alternatively `client.ListUsagesSlot(ctx, id, webapps.DefaultListUsagesSlotOperationOptions())` can be used to do batched pagination
items, err := client.ListUsagesSlotComplete(ctx, id, webapps.DefaultListUsagesSlotOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WebAppsClient.ListWorkflowsConnectionsSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

read, err := client.ListWorkflowsConnectionsSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebAppsClient.ResetSlotConfigurationSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

read, err := client.ResetSlotConfigurationSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebAppsClient.RestartSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

read, err := client.RestartSlot(ctx, id, webapps.DefaultRestartSlotOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebAppsClient.RestoreFromBackupBlobSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

payload := webapps.RestoreRequest{
	// ...
}


if err := client.RestoreFromBackupBlobSlotThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `WebAppsClient.RestoreFromDeletedAppSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

payload := webapps.DeletedAppRestoreRequest{
	// ...
}


if err := client.RestoreFromDeletedAppSlotThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `WebAppsClient.RestoreSnapshotSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

payload := webapps.SnapshotRestoreRequest{
	// ...
}


if err := client.RestoreSnapshotSlotThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `WebAppsClient.StartNetworkTraceSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

if err := client.StartNetworkTraceSlotThenPoll(ctx, id, webapps.DefaultStartNetworkTraceSlotOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `WebAppsClient.StartSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

read, err := client.StartSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebAppsClient.StartWebSiteNetworkTraceOperationSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

if err := client.StartWebSiteNetworkTraceOperationSlotThenPoll(ctx, id, webapps.DefaultStartWebSiteNetworkTraceOperationSlotOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `WebAppsClient.StartWebSiteNetworkTraceSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

read, err := client.StartWebSiteNetworkTraceSlot(ctx, id, webapps.DefaultStartWebSiteNetworkTraceSlotOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebAppsClient.StopNetworkTraceSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

read, err := client.StopNetworkTraceSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebAppsClient.StopSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

read, err := client.StopSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebAppsClient.StopWebSiteNetworkTraceSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

read, err := client.StopWebSiteNetworkTraceSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebAppsClient.SwapSlotSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

payload := webapps.CsmSlotEntity{
	// ...
}


if err := client.SwapSlotSlotThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `WebAppsClient.SyncFunctionTriggersSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

read, err := client.SyncFunctionTriggersSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebAppsClient.SyncFunctionsSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

read, err := client.SyncFunctionsSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebAppsClient.SyncRepositorySlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

read, err := client.SyncRepositorySlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebAppsClient.UpdateApplicationSettingsSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

payload := webapps.StringDictionary{
	// ...
}


read, err := client.UpdateApplicationSettingsSlot(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebAppsClient.UpdateAuthSettingsSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

payload := webapps.SiteAuthSettings{
	// ...
}


read, err := client.UpdateAuthSettingsSlot(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebAppsClient.UpdateAzureStorageAccountsSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

payload := webapps.AzureStoragePropertyDictionaryResource{
	// ...
}


read, err := client.UpdateAzureStorageAccountsSlot(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebAppsClient.UpdateBackupConfigurationSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

payload := webapps.BackupRequest{
	// ...
}


read, err := client.UpdateBackupConfigurationSlot(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebAppsClient.UpdateConnectionStringsSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

payload := webapps.ConnectionStringDictionary{
	// ...
}


read, err := client.UpdateConnectionStringsSlot(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebAppsClient.UpdateMetadataSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

payload := webapps.StringDictionary{
	// ...
}


read, err := client.UpdateMetadataSlot(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebAppsClient.UpdateSitePushSettingsSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

payload := webapps.PushSettings{
	// ...
}


read, err := client.UpdateSitePushSettingsSlot(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebAppsClient.UpdateSlot`

```go
ctx := context.TODO()
id := webapps.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

payload := webapps.SitePatchResource{
	// ...
}


read, err := client.UpdateSlot(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
