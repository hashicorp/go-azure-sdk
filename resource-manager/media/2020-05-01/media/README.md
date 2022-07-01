
## `github.com/hashicorp/go-azure-sdk/resource-manager/media/2020-05-01/media` Documentation

The `media` SDK allows for interaction with the Azure Resource Manager Service `media` (API Version `2020-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/media/2020-05-01/media"
```


### Client Initialization

```go
client := media.NewMediaClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MediaClient.AccountFiltersCreateOrUpdate`

```go
ctx := context.TODO()
id := media.NewAccountFilterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "filterValue")

payload := media.AccountFilter{
	// ...
}


read, err := client.AccountFiltersCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.AccountFiltersDelete`

```go
ctx := context.TODO()
id := media.NewAccountFilterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "filterValue")

read, err := client.AccountFiltersDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.AccountFiltersGet`

```go
ctx := context.TODO()
id := media.NewAccountFilterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "filterValue")

read, err := client.AccountFiltersGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.AccountFiltersList`

```go
ctx := context.TODO()
id := media.NewMediaServiceID("12345678-1234-9876-4563-123456789012", "accountValue")

// alternatively `client.AccountFiltersList(ctx, id)` can be used to do batched pagination
items, err := client.AccountFiltersListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MediaClient.AccountFiltersUpdate`

```go
ctx := context.TODO()
id := media.NewAccountFilterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "filterValue")

payload := media.AccountFilter{
	// ...
}


read, err := client.AccountFiltersUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.AssetFiltersCreateOrUpdate`

```go
ctx := context.TODO()
id := media.NewAssetFilterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "assetValue", "filterValue")

payload := media.AssetFilter{
	// ...
}


read, err := client.AssetFiltersCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.AssetFiltersDelete`

```go
ctx := context.TODO()
id := media.NewAssetFilterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "assetValue", "filterValue")

read, err := client.AssetFiltersDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.AssetFiltersGet`

```go
ctx := context.TODO()
id := media.NewAssetFilterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "assetValue", "filterValue")

read, err := client.AssetFiltersGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.AssetFiltersList`

```go
ctx := context.TODO()
id := media.NewAssetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "assetValue")

// alternatively `client.AssetFiltersList(ctx, id)` can be used to do batched pagination
items, err := client.AssetFiltersListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MediaClient.AssetFiltersUpdate`

```go
ctx := context.TODO()
id := media.NewAssetFilterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "assetValue", "filterValue")

payload := media.AssetFilter{
	// ...
}


read, err := client.AssetFiltersUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.AssetsCreateOrUpdate`

```go
ctx := context.TODO()
id := media.NewAssetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "assetValue")

payload := media.Asset{
	// ...
}


read, err := client.AssetsCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.AssetsDelete`

```go
ctx := context.TODO()
id := media.NewAssetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "assetValue")

read, err := client.AssetsDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.AssetsGet`

```go
ctx := context.TODO()
id := media.NewAssetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "assetValue")

read, err := client.AssetsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.AssetsGetEncryptionKey`

```go
ctx := context.TODO()
id := media.NewAssetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "assetValue")

read, err := client.AssetsGetEncryptionKey(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.AssetsList`

```go
ctx := context.TODO()
id := media.NewMediaServiceID("12345678-1234-9876-4563-123456789012", "accountValue")

// alternatively `client.AssetsList(ctx, id, media.DefaultAssetsListOperationOptions())` can be used to do batched pagination
items, err := client.AssetsListComplete(ctx, id, media.DefaultAssetsListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MediaClient.AssetsListContainerSas`

```go
ctx := context.TODO()
id := media.NewAssetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "assetValue")

payload := media.ListContainerSasInput{
	// ...
}


read, err := client.AssetsListContainerSas(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.AssetsListStreamingLocators`

```go
ctx := context.TODO()
id := media.NewAssetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "assetValue")

read, err := client.AssetsListStreamingLocators(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.AssetsUpdate`

```go
ctx := context.TODO()
id := media.NewAssetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "assetValue")

payload := media.Asset{
	// ...
}


read, err := client.AssetsUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.ContentKeyPoliciesCreateOrUpdate`

```go
ctx := context.TODO()
id := media.NewContentKeyPoliciesID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "contentKeyPolicyValue")

payload := media.ContentKeyPolicy{
	// ...
}


read, err := client.ContentKeyPoliciesCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.ContentKeyPoliciesDelete`

```go
ctx := context.TODO()
id := media.NewContentKeyPoliciesID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "contentKeyPolicyValue")

read, err := client.ContentKeyPoliciesDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.ContentKeyPoliciesGet`

```go
ctx := context.TODO()
id := media.NewContentKeyPoliciesID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "contentKeyPolicyValue")

read, err := client.ContentKeyPoliciesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.ContentKeyPoliciesGetPolicyPropertiesWithSecrets`

```go
ctx := context.TODO()
id := media.NewContentKeyPoliciesID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "contentKeyPolicyValue")

read, err := client.ContentKeyPoliciesGetPolicyPropertiesWithSecrets(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.ContentKeyPoliciesList`

```go
ctx := context.TODO()
id := media.NewMediaServiceID("12345678-1234-9876-4563-123456789012", "accountValue")

// alternatively `client.ContentKeyPoliciesList(ctx, id, media.DefaultContentKeyPoliciesListOperationOptions())` can be used to do batched pagination
items, err := client.ContentKeyPoliciesListComplete(ctx, id, media.DefaultContentKeyPoliciesListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MediaClient.ContentKeyPoliciesUpdate`

```go
ctx := context.TODO()
id := media.NewContentKeyPoliciesID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "contentKeyPolicyValue")

payload := media.ContentKeyPolicy{
	// ...
}


read, err := client.ContentKeyPoliciesUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.JobsCancelJob`

```go
ctx := context.TODO()
id := media.NewJobID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "transformValue", "jobValue")

read, err := client.JobsCancelJob(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.JobsCreate`

```go
ctx := context.TODO()
id := media.NewJobID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "transformValue", "jobValue")

payload := media.Job{
	// ...
}


read, err := client.JobsCreate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.JobsDelete`

```go
ctx := context.TODO()
id := media.NewJobID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "transformValue", "jobValue")

read, err := client.JobsDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.JobsGet`

```go
ctx := context.TODO()
id := media.NewJobID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "transformValue", "jobValue")

read, err := client.JobsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.JobsList`

```go
ctx := context.TODO()
id := media.NewTransformID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "transformValue")

// alternatively `client.JobsList(ctx, id, media.DefaultJobsListOperationOptions())` can be used to do batched pagination
items, err := client.JobsListComplete(ctx, id, media.DefaultJobsListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MediaClient.JobsUpdate`

```go
ctx := context.TODO()
id := media.NewJobID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "transformValue", "jobValue")

payload := media.Job{
	// ...
}


read, err := client.JobsUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.LocationsCheckNameAvailability`

```go
ctx := context.TODO()
id := media.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

payload := media.CheckNameAvailabilityInput{
	// ...
}


read, err := client.LocationsCheckNameAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.MediaservicesCreateOrUpdate`

```go
ctx := context.TODO()
id := media.NewProviderMediaServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")

payload := media.MediaService{
	// ...
}


read, err := client.MediaservicesCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.MediaservicesDelete`

```go
ctx := context.TODO()
id := media.NewProviderMediaServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")

read, err := client.MediaservicesDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.MediaservicesGet`

```go
ctx := context.TODO()
id := media.NewProviderMediaServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")

read, err := client.MediaservicesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.MediaservicesGetBySubscription`

```go
ctx := context.TODO()
id := media.NewMediaServiceID("12345678-1234-9876-4563-123456789012", "accountValue")

read, err := client.MediaservicesGetBySubscription(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.MediaservicesList`

```go
ctx := context.TODO()
id := media.NewResourceGroupID()

// alternatively `client.MediaservicesList(ctx, id)` can be used to do batched pagination
items, err := client.MediaservicesListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MediaClient.MediaservicesListBySubscription`

```go
ctx := context.TODO()
id := media.NewSubscriptionID()

// alternatively `client.MediaservicesListBySubscription(ctx, id)` can be used to do batched pagination
items, err := client.MediaservicesListBySubscriptionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MediaClient.MediaservicesListEdgePolicies`

```go
ctx := context.TODO()
id := media.NewProviderMediaServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")

payload := media.ListEdgePoliciesInput{
	// ...
}


read, err := client.MediaservicesListEdgePolicies(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.MediaservicesSyncStorageKeys`

```go
ctx := context.TODO()
id := media.NewProviderMediaServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")

payload := media.SyncStorageKeysInput{
	// ...
}


read, err := client.MediaservicesSyncStorageKeys(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.MediaservicesUpdate`

```go
ctx := context.TODO()
id := media.NewProviderMediaServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")

payload := media.MediaService{
	// ...
}


read, err := client.MediaservicesUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.PrivateEndpointConnectionsCreateOrUpdate`

```go
ctx := context.TODO()
id := media.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "nameValue")

payload := media.PrivateEndpointConnection{
	// ...
}


read, err := client.PrivateEndpointConnectionsCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.PrivateEndpointConnectionsDelete`

```go
ctx := context.TODO()
id := media.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "nameValue")

read, err := client.PrivateEndpointConnectionsDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.PrivateEndpointConnectionsGet`

```go
ctx := context.TODO()
id := media.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "nameValue")

read, err := client.PrivateEndpointConnectionsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.PrivateEndpointConnectionsList`

```go
ctx := context.TODO()
id := media.NewProviderMediaServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")

read, err := client.PrivateEndpointConnectionsList(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.PrivateLinkResourcesGet`

```go
ctx := context.TODO()
id := media.NewPrivateLinkResourceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "nameValue")

read, err := client.PrivateLinkResourcesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.PrivateLinkResourcesList`

```go
ctx := context.TODO()
id := media.NewProviderMediaServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")

read, err := client.PrivateLinkResourcesList(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.StreamingLocatorsCreate`

```go
ctx := context.TODO()
id := media.NewStreamingLocatorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "streamingLocatorValue")

payload := media.StreamingLocator{
	// ...
}


read, err := client.StreamingLocatorsCreate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.StreamingLocatorsDelete`

```go
ctx := context.TODO()
id := media.NewStreamingLocatorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "streamingLocatorValue")

read, err := client.StreamingLocatorsDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.StreamingLocatorsGet`

```go
ctx := context.TODO()
id := media.NewStreamingLocatorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "streamingLocatorValue")

read, err := client.StreamingLocatorsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.StreamingLocatorsList`

```go
ctx := context.TODO()
id := media.NewMediaServiceID("12345678-1234-9876-4563-123456789012", "accountValue")

// alternatively `client.StreamingLocatorsList(ctx, id, media.DefaultStreamingLocatorsListOperationOptions())` can be used to do batched pagination
items, err := client.StreamingLocatorsListComplete(ctx, id, media.DefaultStreamingLocatorsListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MediaClient.StreamingLocatorsListContentKeys`

```go
ctx := context.TODO()
id := media.NewStreamingLocatorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "streamingLocatorValue")

read, err := client.StreamingLocatorsListContentKeys(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.StreamingLocatorsListPaths`

```go
ctx := context.TODO()
id := media.NewStreamingLocatorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "streamingLocatorValue")

read, err := client.StreamingLocatorsListPaths(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.StreamingPoliciesCreate`

```go
ctx := context.TODO()
id := media.NewStreamingPoliciesID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "streamingPolicyValue")

payload := media.StreamingPolicy{
	// ...
}


read, err := client.StreamingPoliciesCreate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.StreamingPoliciesDelete`

```go
ctx := context.TODO()
id := media.NewStreamingPoliciesID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "streamingPolicyValue")

read, err := client.StreamingPoliciesDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.StreamingPoliciesGet`

```go
ctx := context.TODO()
id := media.NewStreamingPoliciesID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "streamingPolicyValue")

read, err := client.StreamingPoliciesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.StreamingPoliciesList`

```go
ctx := context.TODO()
id := media.NewMediaServiceID("12345678-1234-9876-4563-123456789012", "accountValue")

// alternatively `client.StreamingPoliciesList(ctx, id, media.DefaultStreamingPoliciesListOperationOptions())` can be used to do batched pagination
items, err := client.StreamingPoliciesListComplete(ctx, id, media.DefaultStreamingPoliciesListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MediaClient.TransformsCreateOrUpdate`

```go
ctx := context.TODO()
id := media.NewTransformID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "transformValue")

payload := media.Transform{
	// ...
}


read, err := client.TransformsCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.TransformsDelete`

```go
ctx := context.TODO()
id := media.NewTransformID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "transformValue")

read, err := client.TransformsDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.TransformsGet`

```go
ctx := context.TODO()
id := media.NewTransformID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "transformValue")

read, err := client.TransformsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.TransformsList`

```go
ctx := context.TODO()
id := media.NewMediaServiceID("12345678-1234-9876-4563-123456789012", "accountValue")

// alternatively `client.TransformsList(ctx, id, media.DefaultTransformsListOperationOptions())` can be used to do batched pagination
items, err := client.TransformsListComplete(ctx, id, media.DefaultTransformsListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MediaClient.TransformsUpdate`

```go
ctx := context.TODO()
id := media.NewTransformID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "transformValue")

payload := media.Transform{
	// ...
}


read, err := client.TransformsUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
