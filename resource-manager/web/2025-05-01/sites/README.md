
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/sites` Documentation

The `sites` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/sites"
```


### Client Initialization

```go
client := sites.NewSitesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SitesClient.RecommendationsDisableAllForWebApp`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.RecommendationsDisableAllForWebApp(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.RecommendationsListHistoryForWebApp`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

// alternatively `client.RecommendationsListHistoryForWebApp(ctx, id, sites.DefaultRecommendationsListHistoryForWebAppOperationOptions())` can be used to do batched pagination
items, err := client.RecommendationsListHistoryForWebAppComplete(ctx, id, sites.DefaultRecommendationsListHistoryForWebAppOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SitesClient.RecommendationsListRecommendedRulesForWebApp`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

// alternatively `client.RecommendationsListRecommendedRulesForWebApp(ctx, id, sites.DefaultRecommendationsListRecommendedRulesForWebAppOperationOptions())` can be used to do batched pagination
items, err := client.RecommendationsListRecommendedRulesForWebAppComplete(ctx, id, sites.DefaultRecommendationsListRecommendedRulesForWebAppOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SitesClient.RecommendationsResetAllFiltersForWebApp`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.RecommendationsResetAllFiltersForWebApp(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsAnalyzeCustomHostname`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.WebAppsAnalyzeCustomHostname(ctx, id, sites.DefaultWebAppsAnalyzeCustomHostnameOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsApplySlotConfigToProduction`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

payload := sites.CsmSlotEntity{
	// ...
}


read, err := client.WebAppsApplySlotConfigToProduction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsBackup`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

payload := sites.BackupRequest{
	// ...
}


read, err := client.WebAppsBackup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsCreateOneDeployOperation`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.WebAppsCreateOneDeployOperation(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsCreateOrUpdate`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

payload := sites.Site{
	// ...
}


if err := client.WebAppsCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `SitesClient.WebAppsCreateOrUpdateHostSecret`

```go
ctx := context.TODO()
id := sites.NewDefaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "defaultName", "keyName")

payload := sites.KeyInfo{
	// ...
}


read, err := client.WebAppsCreateOrUpdateHostSecret(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsDelete`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.WebAppsDelete(ctx, id, sites.DefaultWebAppsDeleteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsDeleteBackupConfiguration`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.WebAppsDeleteBackupConfiguration(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsDeleteHostSecret`

```go
ctx := context.TODO()
id := sites.NewDefaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "defaultName", "keyName")

read, err := client.WebAppsDeleteHostSecret(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsDeployWorkflowArtifacts`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

payload := sites.WorkflowArtifacts{
	// ...
}


read, err := client.WebAppsDeployWorkflowArtifacts(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsDiscoverBackup`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

payload := sites.RestoreRequest{
	// ...
}


read, err := client.WebAppsDiscoverBackup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsGenerateNewSitePublishingPassword`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.WebAppsGenerateNewSitePublishingPassword(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsGet`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.WebAppsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsGetAuthSettings`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.WebAppsGetAuthSettings(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsGetBackupConfiguration`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.WebAppsGetBackupConfiguration(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsGetContainerLogsZip`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.WebAppsGetContainerLogsZip(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsGetFunctionsAdminToken`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.WebAppsGetFunctionsAdminToken(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsGetNetworkTraces`

```go
ctx := context.TODO()
id := sites.NewNetworkTraceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "operationId")

read, err := client.WebAppsGetNetworkTraces(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsGetNetworkTracesV2`

```go
ctx := context.TODO()
id := sites.NewSiteNetworkTraceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "operationId")

read, err := client.WebAppsGetNetworkTracesV2(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsGetOneDeployStatus`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.WebAppsGetOneDeployStatus(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsGetPrivateLinkResources`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.WebAppsGetPrivateLinkResources(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsGetSitePhpErrorLogFlag`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.WebAppsGetSitePhpErrorLogFlag(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsGetWebSiteContainerLogs`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.WebAppsGetWebSiteContainerLogs(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsIsCloneable`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.WebAppsIsCloneable(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsList`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.WebAppsList(ctx, id)` can be used to do batched pagination
items, err := client.WebAppsListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SitesClient.WebAppsListApplicationSettings`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.WebAppsListApplicationSettings(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsListAzureStorageAccounts`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.WebAppsListAzureStorageAccounts(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsListConnectionStrings`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.WebAppsListConnectionStrings(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsListHostKeys`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.WebAppsListHostKeys(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsListHybridConnections`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.WebAppsListHybridConnections(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsListMetadata`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.WebAppsListMetadata(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsListPerfMonCounters`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

// alternatively `client.WebAppsListPerfMonCounters(ctx, id, sites.DefaultWebAppsListPerfMonCountersOperationOptions())` can be used to do batched pagination
items, err := client.WebAppsListPerfMonCountersComplete(ctx, id, sites.DefaultWebAppsListPerfMonCountersOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SitesClient.WebAppsListPremierAddOns`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.WebAppsListPremierAddOns(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsListPublishingCredentials`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

if err := client.WebAppsListPublishingCredentialsThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `SitesClient.WebAppsListPublishingProfileXmlWithSecrets`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

payload := sites.CsmPublishingProfileOptions{
	// ...
}


read, err := client.WebAppsListPublishingProfileXmlWithSecrets(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsListRelayServiceConnections`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.WebAppsListRelayServiceConnections(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsListSiteBackups`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

// alternatively `client.WebAppsListSiteBackups(ctx, id)` can be used to do batched pagination
items, err := client.WebAppsListSiteBackupsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SitesClient.WebAppsListSitePushSettings`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.WebAppsListSitePushSettings(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsListSlotDifferencesFromProduction`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

payload := sites.CsmSlotEntity{
	// ...
}


// alternatively `client.WebAppsListSlotDifferencesFromProduction(ctx, id, payload)` can be used to do batched pagination
items, err := client.WebAppsListSlotDifferencesFromProductionComplete(ctx, id, payload)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SitesClient.WebAppsListSnapshots`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

// alternatively `client.WebAppsListSnapshots(ctx, id)` can be used to do batched pagination
items, err := client.WebAppsListSnapshotsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SitesClient.WebAppsListSnapshotsFromDRSecondary`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

// alternatively `client.WebAppsListSnapshotsFromDRSecondary(ctx, id)` can be used to do batched pagination
items, err := client.WebAppsListSnapshotsFromDRSecondaryComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SitesClient.WebAppsListSyncFunctionTriggers`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.WebAppsListSyncFunctionTriggers(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsListSyncStatus`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.WebAppsListSyncStatus(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsListUsages`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

// alternatively `client.WebAppsListUsages(ctx, id, sites.DefaultWebAppsListUsagesOperationOptions())` can be used to do batched pagination
items, err := client.WebAppsListUsagesComplete(ctx, id, sites.DefaultWebAppsListUsagesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SitesClient.WebAppsListWorkflowsConnections`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.WebAppsListWorkflowsConnections(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsMigrateMySql`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

payload := sites.MigrateMySqlRequest{
	// ...
}


if err := client.WebAppsMigrateMySqlThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `SitesClient.WebAppsMigrateStorage`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

payload := sites.StorageMigrationOptions{
	// ...
}


if err := client.WebAppsMigrateStorageThenPoll(ctx, id, payload, sites.DefaultWebAppsMigrateStorageOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `SitesClient.WebAppsResetProductionSlotConfig`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.WebAppsResetProductionSlotConfig(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsRestart`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.WebAppsRestart(ctx, id, sites.DefaultWebAppsRestartOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsRestoreFromBackupBlob`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

payload := sites.RestoreRequest{
	// ...
}


if err := client.WebAppsRestoreFromBackupBlobThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `SitesClient.WebAppsRestoreFromDeletedApp`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

payload := sites.DeletedAppRestoreRequest{
	// ...
}


if err := client.WebAppsRestoreFromDeletedAppThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `SitesClient.WebAppsRestoreSnapshot`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

payload := sites.SnapshotRestoreRequest{
	// ...
}


if err := client.WebAppsRestoreSnapshotThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `SitesClient.WebAppsStart`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.WebAppsStart(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsStartNetworkTrace`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

if err := client.WebAppsStartNetworkTraceThenPoll(ctx, id, sites.DefaultWebAppsStartNetworkTraceOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `SitesClient.WebAppsStartWebSiteNetworkTrace`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.WebAppsStartWebSiteNetworkTrace(ctx, id, sites.DefaultWebAppsStartWebSiteNetworkTraceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsStartWebSiteNetworkTraceOperation`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

if err := client.WebAppsStartWebSiteNetworkTraceOperationThenPoll(ctx, id, sites.DefaultWebAppsStartWebSiteNetworkTraceOperationOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `SitesClient.WebAppsStop`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.WebAppsStop(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsStopNetworkTrace`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.WebAppsStopNetworkTrace(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsStopWebSiteNetworkTrace`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.WebAppsStopWebSiteNetworkTrace(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsSwapSlotWithProduction`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

payload := sites.CsmSlotEntity{
	// ...
}


if err := client.WebAppsSwapSlotWithProductionThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `SitesClient.WebAppsSyncFunctionTriggers`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.WebAppsSyncFunctionTriggers(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsSyncFunctions`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.WebAppsSyncFunctions(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsSyncRepository`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.WebAppsSyncRepository(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsUpdate`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

payload := sites.SitePatchResource{
	// ...
}


read, err := client.WebAppsUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsUpdateApplicationSettings`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

payload := sites.StringDictionary{
	// ...
}


read, err := client.WebAppsUpdateApplicationSettings(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsUpdateAuthSettings`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

payload := sites.SiteAuthSettings{
	// ...
}


read, err := client.WebAppsUpdateAuthSettings(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsUpdateAzureStorageAccounts`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

payload := sites.AzureStoragePropertyDictionaryResource{
	// ...
}


read, err := client.WebAppsUpdateAzureStorageAccounts(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsUpdateBackupConfiguration`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

payload := sites.BackupRequest{
	// ...
}


read, err := client.WebAppsUpdateBackupConfiguration(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsUpdateConnectionStrings`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

payload := sites.ConnectionStringDictionary{
	// ...
}


read, err := client.WebAppsUpdateConnectionStrings(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsUpdateMachineKey`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.WebAppsUpdateMachineKey(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsUpdateMetadata`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

payload := sites.StringDictionary{
	// ...
}


read, err := client.WebAppsUpdateMetadata(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WebAppsUpdateSitePushSettings`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

payload := sites.PushSettings{
	// ...
}


read, err := client.WebAppsUpdateSitePushSettings(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WorkflowsRegenerateAccessKey`

```go
ctx := context.TODO()
id := sites.NewManagementWorkflowID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "workflowName")

payload := sites.RegenerateActionParameter{
	// ...
}


read, err := client.WorkflowsRegenerateAccessKey(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.WorkflowsValidate`

```go
ctx := context.TODO()
id := sites.NewManagementWorkflowID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "workflowName")

payload := sites.Workflow{
	// ...
}


read, err := client.WorkflowsValidate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
