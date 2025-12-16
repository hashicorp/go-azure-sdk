
## `github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-12-01/databoxedgedevices` Documentation

The `databoxedgedevices` SDK allows for interaction with Azure Resource Manager `databoxedge` (API Version `2023-12-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-12-01/databoxedgedevices"
```


### Client Initialization

```go
client := databoxedgedevices.NewDataBoxEdgeDevicesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DataBoxEdgeDevicesClient.DeviceCapacityCheckCheckResourceCreationFeasibility`

```go
ctx := context.TODO()
id := databoxedgedevices.NewDataBoxEdgeDeviceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceName")

payload := databoxedgedevices.DeviceCapacityRequestInfo{
	// ...
}


if err := client.DeviceCapacityCheckCheckResourceCreationFeasibilityThenPoll(ctx, id, payload, databoxedgedevices.DefaultDeviceCapacityCheckCheckResourceCreationFeasibilityOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `DataBoxEdgeDevicesClient.DevicesCreateOrUpdate`

```go
ctx := context.TODO()
id := databoxedgedevices.NewDataBoxEdgeDeviceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceName")

payload := databoxedgedevices.DataBoxEdgeDevice{
	// ...
}


read, err := client.DevicesCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataBoxEdgeDevicesClient.DevicesCreateOrUpdateSecuritySettings`

```go
ctx := context.TODO()
id := databoxedgedevices.NewDataBoxEdgeDeviceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceName")

payload := databoxedgedevices.SecuritySettings{
	// ...
}


if err := client.DevicesCreateOrUpdateSecuritySettingsThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DataBoxEdgeDevicesClient.DevicesDelete`

```go
ctx := context.TODO()
id := databoxedgedevices.NewDataBoxEdgeDeviceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceName")

if err := client.DevicesDeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `DataBoxEdgeDevicesClient.DevicesDownloadUpdates`

```go
ctx := context.TODO()
id := databoxedgedevices.NewDataBoxEdgeDeviceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceName")

if err := client.DevicesDownloadUpdatesThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `DataBoxEdgeDevicesClient.DevicesGenerateCertificate`

```go
ctx := context.TODO()
id := databoxedgedevices.NewDataBoxEdgeDeviceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceName")

read, err := client.DevicesGenerateCertificate(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataBoxEdgeDevicesClient.DevicesGet`

```go
ctx := context.TODO()
id := databoxedgedevices.NewDataBoxEdgeDeviceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceName")

read, err := client.DevicesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataBoxEdgeDevicesClient.DevicesGetExtendedInformation`

```go
ctx := context.TODO()
id := databoxedgedevices.NewDataBoxEdgeDeviceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceName")

read, err := client.DevicesGetExtendedInformation(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataBoxEdgeDevicesClient.DevicesInstallUpdates`

```go
ctx := context.TODO()
id := databoxedgedevices.NewDataBoxEdgeDeviceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceName")

if err := client.DevicesInstallUpdatesThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `DataBoxEdgeDevicesClient.DevicesListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.DevicesListByResourceGroup(ctx, id, databoxedgedevices.DefaultDevicesListByResourceGroupOperationOptions())` can be used to do batched pagination
items, err := client.DevicesListByResourceGroupComplete(ctx, id, databoxedgedevices.DefaultDevicesListByResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DataBoxEdgeDevicesClient.DevicesListBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.DevicesListBySubscription(ctx, id, databoxedgedevices.DefaultDevicesListBySubscriptionOperationOptions())` can be used to do batched pagination
items, err := client.DevicesListBySubscriptionComplete(ctx, id, databoxedgedevices.DefaultDevicesListBySubscriptionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DataBoxEdgeDevicesClient.DevicesScanForUpdates`

```go
ctx := context.TODO()
id := databoxedgedevices.NewDataBoxEdgeDeviceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceName")

if err := client.DevicesScanForUpdatesThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `DataBoxEdgeDevicesClient.DevicesUpdate`

```go
ctx := context.TODO()
id := databoxedgedevices.NewDataBoxEdgeDeviceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceName")

payload := databoxedgedevices.DataBoxEdgeDevicePatch{
	// ...
}


read, err := client.DevicesUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataBoxEdgeDevicesClient.DevicesUpdateExtendedInformation`

```go
ctx := context.TODO()
id := databoxedgedevices.NewDataBoxEdgeDeviceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceName")

payload := databoxedgedevices.DataBoxEdgeDeviceExtendedInfoPatch{
	// ...
}


read, err := client.DevicesUpdateExtendedInformation(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataBoxEdgeDevicesClient.DevicesUploadCertificate`

```go
ctx := context.TODO()
id := databoxedgedevices.NewDataBoxEdgeDeviceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceName")

payload := databoxedgedevices.UploadCertificateRequest{
	// ...
}


read, err := client.DevicesUploadCertificate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataBoxEdgeDevicesClient.NodesListByDataBoxEdgeDevice`

```go
ctx := context.TODO()
id := databoxedgedevices.NewDataBoxEdgeDeviceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceName")

// alternatively `client.NodesListByDataBoxEdgeDevice(ctx, id)` can be used to do batched pagination
items, err := client.NodesListByDataBoxEdgeDeviceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DataBoxEdgeDevicesClient.SupportPackagesTriggerSupportPackage`

```go
ctx := context.TODO()
id := databoxedgedevices.NewDataBoxEdgeDeviceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceName")

payload := databoxedgedevices.TriggerSupportPackageRequest{
	// ...
}


if err := client.SupportPackagesTriggerSupportPackageThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
