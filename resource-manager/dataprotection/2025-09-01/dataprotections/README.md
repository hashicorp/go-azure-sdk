
## `github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2025-09-01/dataprotections` Documentation

The `dataprotections` SDK allows for interaction with Azure Resource Manager `dataprotection` (API Version `2025-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2025-09-01/dataprotections"
```


### Client Initialization

```go
client := dataprotections.NewDataprotectionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DataprotectionsClient.BackupInstancesExtensionRoutingList`

```go
ctx := context.TODO()
id := commonids.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

// alternatively `client.BackupInstancesExtensionRoutingList(ctx, id)` can be used to do batched pagination
items, err := client.BackupInstancesExtensionRoutingListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DataprotectionsClient.BackupInstancesTriggerCrossRegionRestore`

```go
ctx := context.TODO()
id := dataprotections.NewProviderLocationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "locationName")

payload := dataprotections.CrossRegionRestoreRequestObject{
	// ...
}


if err := client.BackupInstancesTriggerCrossRegionRestoreThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DataprotectionsClient.BackupInstancesValidateCrossRegionRestore`

```go
ctx := context.TODO()
id := dataprotections.NewProviderLocationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "locationName")

payload := dataprotections.ValidateCrossRegionRestoreRequestObject{
	// ...
}


if err := client.BackupInstancesValidateCrossRegionRestoreThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DataprotectionsClient.BackupVaultsCheckNameAvailability`

```go
ctx := context.TODO()
id := dataprotections.NewProviderLocationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "locationName")

payload := dataprotections.CheckNameAvailabilityRequest{
	// ...
}


read, err := client.BackupVaultsCheckNameAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataprotectionsClient.DataProtectionCheckFeatureSupport`

```go
ctx := context.TODO()
id := dataprotections.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

payload := dataprotections.FeatureValidationRequestBase{
	// ...
}


read, err := client.DataProtectionCheckFeatureSupport(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataprotectionsClient.FetchCrossRegionRestoreJobGet`

```go
ctx := context.TODO()
id := dataprotections.NewProviderLocationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "locationName")

payload := dataprotections.CrossRegionRestoreJobRequest{
	// ...
}


read, err := client.FetchCrossRegionRestoreJobGet(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataprotectionsClient.FetchCrossRegionRestoreJobsList`

```go
ctx := context.TODO()
id := dataprotections.NewProviderLocationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "locationName")

payload := dataprotections.CrossRegionRestoreJobsRequest{
	// ...
}


// alternatively `client.FetchCrossRegionRestoreJobsList(ctx, id, payload, dataprotections.DefaultFetchCrossRegionRestoreJobsListOperationOptions())` can be used to do batched pagination
items, err := client.FetchCrossRegionRestoreJobsListComplete(ctx, id, payload, dataprotections.DefaultFetchCrossRegionRestoreJobsListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DataprotectionsClient.FetchSecondaryRecoveryPointsList`

```go
ctx := context.TODO()
id := dataprotections.NewProviderLocationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "locationName")

payload := dataprotections.FetchSecondaryRPsRequestParameters{
	// ...
}


// alternatively `client.FetchSecondaryRecoveryPointsList(ctx, id, payload, dataprotections.DefaultFetchSecondaryRecoveryPointsListOperationOptions())` can be used to do batched pagination
items, err := client.FetchSecondaryRecoveryPointsListComplete(ctx, id, payload, dataprotections.DefaultFetchSecondaryRecoveryPointsListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
