
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/virtualendpointreport` Documentation

The `virtualendpointreport` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/virtualendpointreport"
```


### Client Initialization

```go
client := virtualendpointreport.NewVirtualEndpointReportClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `VirtualEndpointReportClient.DeleteVirtualEndpointReport`

```go
ctx := context.TODO()


read, err := client.DeleteVirtualEndpointReport(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointReportClient.GetDeviceManagementVirtualEndpointReportsActionStatusReport`

```go
ctx := context.TODO()

payload := virtualendpointreport.GetDeviceManagementVirtualEndpointReportsActionStatusReportRequest{
	// ...
}


read, err := client.GetDeviceManagementVirtualEndpointReportsActionStatusReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointReportClient.GetDeviceManagementVirtualEndpointReportsCloudPCRecommendationReport`

```go
ctx := context.TODO()

payload := virtualendpointreport.GetDeviceManagementVirtualEndpointReportsCloudPCRecommendationReportRequest{
	// ...
}


read, err := client.GetDeviceManagementVirtualEndpointReportsCloudPCRecommendationReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointReportClient.GetDeviceManagementVirtualEndpointReportsConnectionQualityReport`

```go
ctx := context.TODO()

payload := virtualendpointreport.GetDeviceManagementVirtualEndpointReportsConnectionQualityReportRequest{
	// ...
}


read, err := client.GetDeviceManagementVirtualEndpointReportsConnectionQualityReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointReportClient.GetDeviceManagementVirtualEndpointReportsDailyAggregatedRemoteConnectionReport`

```go
ctx := context.TODO()

payload := virtualendpointreport.GetDeviceManagementVirtualEndpointReportsDailyAggregatedRemoteConnectionReportRequest{
	// ...
}


read, err := client.GetDeviceManagementVirtualEndpointReportsDailyAggregatedRemoteConnectionReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointReportClient.GetDeviceManagementVirtualEndpointReportsFrontlineReport`

```go
ctx := context.TODO()

payload := virtualendpointreport.GetDeviceManagementVirtualEndpointReportsFrontlineReportRequest{
	// ...
}


read, err := client.GetDeviceManagementVirtualEndpointReportsFrontlineReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointReportClient.GetDeviceManagementVirtualEndpointReportsInaccessibleCloudPCReport`

```go
ctx := context.TODO()

payload := virtualendpointreport.GetDeviceManagementVirtualEndpointReportsInaccessibleCloudPCReportRequest{
	// ...
}


read, err := client.GetDeviceManagementVirtualEndpointReportsInaccessibleCloudPCReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointReportClient.GetDeviceManagementVirtualEndpointReportsRawRemoteConnectionReport`

```go
ctx := context.TODO()

payload := virtualendpointreport.GetDeviceManagementVirtualEndpointReportsRawRemoteConnectionReportRequest{
	// ...
}


read, err := client.GetDeviceManagementVirtualEndpointReportsRawRemoteConnectionReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointReportClient.GetDeviceManagementVirtualEndpointReportsRemoteConnectionHistoricalReport`

```go
ctx := context.TODO()

payload := virtualendpointreport.GetDeviceManagementVirtualEndpointReportsRemoteConnectionHistoricalReportRequest{
	// ...
}


read, err := client.GetDeviceManagementVirtualEndpointReportsRemoteConnectionHistoricalReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointReportClient.GetDeviceManagementVirtualEndpointReportsSharedUseLicenseUsageReport`

```go
ctx := context.TODO()

payload := virtualendpointreport.GetDeviceManagementVirtualEndpointReportsSharedUseLicenseUsageReportRequest{
	// ...
}


read, err := client.GetDeviceManagementVirtualEndpointReportsSharedUseLicenseUsageReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointReportClient.GetDeviceManagementVirtualEndpointReportsTotalAggregatedRemoteConnectionReport`

```go
ctx := context.TODO()

payload := virtualendpointreport.GetDeviceManagementVirtualEndpointReportsTotalAggregatedRemoteConnectionReportRequest{
	// ...
}


read, err := client.GetDeviceManagementVirtualEndpointReportsTotalAggregatedRemoteConnectionReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointReportClient.GetVirtualEndpointReport`

```go
ctx := context.TODO()


read, err := client.GetVirtualEndpointReport(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointReportClient.UpdateVirtualEndpointReport`

```go
ctx := context.TODO()

payload := virtualendpointreport.CloudPCReports{
	// ...
}


read, err := client.UpdateVirtualEndpointReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
