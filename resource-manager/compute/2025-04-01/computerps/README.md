
## `github.com/hashicorp/go-azure-sdk/resource-manager/compute/2025-04-01/computerps` Documentation

The `computerps` SDK allows for interaction with Azure Resource Manager `compute` (API Version `2025-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/compute/2025-04-01/computerps"
```


### Client Initialization

```go
client := computerps.NewComputeRPSClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ComputeRPSClient.LogAnalyticsExportRequestRateByInterval`

```go
ctx := context.TODO()
id := computerps.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

payload := computerps.RequestRateByIntervalInput{
	// ...
}


if err := client.LogAnalyticsExportRequestRateByIntervalThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ComputeRPSClient.LogAnalyticsExportThrottledRequests`

```go
ctx := context.TODO()
id := computerps.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

payload := computerps.LogAnalyticsInputBase{
	// ...
}


if err := client.LogAnalyticsExportThrottledRequestsThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ComputeRPSClient.UsageList`

```go
ctx := context.TODO()
id := computerps.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

// alternatively `client.UsageList(ctx, id)` can be used to do batched pagination
items, err := client.UsageListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ComputeRPSClient.VirtualMachineExtensionImagesGet`

```go
ctx := context.TODO()
id := computerps.NewVersionID("12345678-1234-9876-4563-123456789012", "locationName", "publisherName", "typeName", "versionName")

read, err := client.VirtualMachineExtensionImagesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComputeRPSClient.VirtualMachineExtensionImagesListTypes`

```go
ctx := context.TODO()
id := computerps.NewPublisherID("12345678-1234-9876-4563-123456789012", "locationName", "publisherName")

read, err := client.VirtualMachineExtensionImagesListTypes(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComputeRPSClient.VirtualMachineExtensionImagesListVersions`

```go
ctx := context.TODO()
id := computerps.NewTypeID("12345678-1234-9876-4563-123456789012", "locationName", "publisherName", "typeName")

read, err := client.VirtualMachineExtensionImagesListVersions(ctx, id, computerps.DefaultVirtualMachineExtensionImagesListVersionsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComputeRPSClient.VirtualMachineImagesEdgeZoneGet`

```go
ctx := context.TODO()
id := computerps.NewOfferSkuVersionID("12345678-1234-9876-4563-123456789012", "locationName", "edgeZoneName", "publisherName", "offerName", "skuName", "versionName")

read, err := client.VirtualMachineImagesEdgeZoneGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComputeRPSClient.VirtualMachineImagesEdgeZoneList`

```go
ctx := context.TODO()
id := computerps.NewOfferSkuID("12345678-1234-9876-4563-123456789012", "locationName", "edgeZoneName", "publisherName", "offerName", "skuName")

read, err := client.VirtualMachineImagesEdgeZoneList(ctx, id, computerps.DefaultVirtualMachineImagesEdgeZoneListOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComputeRPSClient.VirtualMachineImagesEdgeZoneListOffers`

```go
ctx := context.TODO()
id := computerps.NewEdgeZonePublisherID("12345678-1234-9876-4563-123456789012", "locationName", "edgeZoneName", "publisherName")

read, err := client.VirtualMachineImagesEdgeZoneListOffers(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComputeRPSClient.VirtualMachineImagesEdgeZoneListPublishers`

```go
ctx := context.TODO()
id := computerps.NewEdgeZoneID("12345678-1234-9876-4563-123456789012", "locationName", "edgeZoneName")

read, err := client.VirtualMachineImagesEdgeZoneListPublishers(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComputeRPSClient.VirtualMachineImagesEdgeZoneListSkus`

```go
ctx := context.TODO()
id := computerps.NewVMImageOfferID("12345678-1234-9876-4563-123456789012", "locationName", "edgeZoneName", "publisherName", "offerName")

read, err := client.VirtualMachineImagesEdgeZoneListSkus(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComputeRPSClient.VirtualMachineImagesGet`

```go
ctx := context.TODO()
id := computerps.NewSkuVersionID("12345678-1234-9876-4563-123456789012", "locationName", "publisherName", "offerName", "skuName", "versionName")

read, err := client.VirtualMachineImagesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComputeRPSClient.VirtualMachineImagesList`

```go
ctx := context.TODO()
id := computerps.NewSkuID("12345678-1234-9876-4563-123456789012", "locationName", "publisherName", "offerName", "skuName")

read, err := client.VirtualMachineImagesList(ctx, id, computerps.DefaultVirtualMachineImagesListOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComputeRPSClient.VirtualMachineImagesListByEdgeZone`

```go
ctx := context.TODO()
id := computerps.NewEdgeZoneID("12345678-1234-9876-4563-123456789012", "locationName", "edgeZoneName")

// alternatively `client.VirtualMachineImagesListByEdgeZone(ctx, id)` can be used to do batched pagination
items, err := client.VirtualMachineImagesListByEdgeZoneComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ComputeRPSClient.VirtualMachineImagesListOffers`

```go
ctx := context.TODO()
id := computerps.NewPublisherID("12345678-1234-9876-4563-123456789012", "locationName", "publisherName")

read, err := client.VirtualMachineImagesListOffers(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComputeRPSClient.VirtualMachineImagesListPublishers`

```go
ctx := context.TODO()
id := computerps.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

read, err := client.VirtualMachineImagesListPublishers(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComputeRPSClient.VirtualMachineImagesListSkus`

```go
ctx := context.TODO()
id := computerps.NewOfferID("12345678-1234-9876-4563-123456789012", "locationName", "publisherName", "offerName")

read, err := client.VirtualMachineImagesListSkus(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComputeRPSClient.VirtualMachineRunCommandsGet`

```go
ctx := context.TODO()
id := computerps.NewRunCommandID("12345678-1234-9876-4563-123456789012", "locationName", "commandId")

read, err := client.VirtualMachineRunCommandsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComputeRPSClient.VirtualMachineRunCommandsList`

```go
ctx := context.TODO()
id := computerps.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

// alternatively `client.VirtualMachineRunCommandsList(ctx, id)` can be used to do batched pagination
items, err := client.VirtualMachineRunCommandsListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ComputeRPSClient.VirtualMachineScaleSetsListByLocation`

```go
ctx := context.TODO()
id := computerps.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

// alternatively `client.VirtualMachineScaleSetsListByLocation(ctx, id)` can be used to do batched pagination
items, err := client.VirtualMachineScaleSetsListByLocationComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ComputeRPSClient.VirtualMachineSizesList`

```go
ctx := context.TODO()
id := computerps.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

// alternatively `client.VirtualMachineSizesList(ctx, id)` can be used to do batched pagination
items, err := client.VirtualMachineSizesListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ComputeRPSClient.VirtualMachinesListByLocation`

```go
ctx := context.TODO()
id := computerps.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

// alternatively `client.VirtualMachinesListByLocation(ctx, id)` can be used to do batched pagination
items, err := client.VirtualMachinesListByLocationComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
