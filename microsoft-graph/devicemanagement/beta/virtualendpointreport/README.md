
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/virtualendpointreport` Documentation

The `virtualendpointreport` SDK allows for interaction with Microsoft Graph `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/virtualendpointreport"
```


### Client Initialization

```go
client := virtualendpointreport.NewVirtualEndpointReportClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `VirtualEndpointReportClient.CreateVirtualEndpointReportRetrieveCrossRegionDisasterRecoveryReport`

```go
ctx := context.TODO()

payload := virtualendpointreport.CreateVirtualEndpointReportRetrieveCrossRegionDisasterRecoveryReportRequest{
	// ...
}


read, err := client.CreateVirtualEndpointReportRetrieveCrossRegionDisasterRecoveryReport(ctx, payload, virtualendpointreport.DefaultCreateVirtualEndpointReportRetrieveCrossRegionDisasterRecoveryReportOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointReportClient.DeleteVirtualEndpointReport`

```go
ctx := context.TODO()


read, err := client.DeleteVirtualEndpointReport(ctx, virtualendpointreport.DefaultDeleteVirtualEndpointReportOperationOptions())
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


read, err := client.GetVirtualEndpointReport(ctx, virtualendpointreport.DefaultGetVirtualEndpointReportOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointReportClient.GetVirtualEndpointReportsActionStatusReports`

```go
ctx := context.TODO()

payload := virtualendpointreport.GetVirtualEndpointReportsActionStatusReportsRequest{
	// ...
}


read, err := client.GetVirtualEndpointReportsActionStatusReports(ctx, payload, virtualendpointreport.DefaultGetVirtualEndpointReportsActionStatusReportsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointReportClient.GetVirtualEndpointReportsCloudPCPerformanceReport`

```go
ctx := context.TODO()

payload := virtualendpointreport.GetVirtualEndpointReportsCloudPCPerformanceReportRequest{
	// ...
}


read, err := client.GetVirtualEndpointReportsCloudPCPerformanceReport(ctx, payload, virtualendpointreport.DefaultGetVirtualEndpointReportsCloudPCPerformanceReportOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointReportClient.GetVirtualEndpointReportsCloudPCRecommendationReports`

```go
ctx := context.TODO()

payload := virtualendpointreport.GetVirtualEndpointReportsCloudPCRecommendationReportsRequest{
	// ...
}


read, err := client.GetVirtualEndpointReportsCloudPCRecommendationReports(ctx, payload, virtualendpointreport.DefaultGetVirtualEndpointReportsCloudPCRecommendationReportsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointReportClient.GetVirtualEndpointReportsConnectionQualityReports`

```go
ctx := context.TODO()

payload := virtualendpointreport.GetVirtualEndpointReportsConnectionQualityReportsRequest{
	// ...
}


read, err := client.GetVirtualEndpointReportsConnectionQualityReports(ctx, payload, virtualendpointreport.DefaultGetVirtualEndpointReportsConnectionQualityReportsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointReportClient.GetVirtualEndpointReportsDailyAggregatedRemoteConnectionReports`

```go
ctx := context.TODO()

payload := virtualendpointreport.GetVirtualEndpointReportsDailyAggregatedRemoteConnectionReportsRequest{
	// ...
}


read, err := client.GetVirtualEndpointReportsDailyAggregatedRemoteConnectionReports(ctx, payload, virtualendpointreport.DefaultGetVirtualEndpointReportsDailyAggregatedRemoteConnectionReportsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointReportClient.GetVirtualEndpointReportsFrontlineReport`

```go
ctx := context.TODO()

payload := virtualendpointreport.GetVirtualEndpointReportsFrontlineReportRequest{
	// ...
}


read, err := client.GetVirtualEndpointReportsFrontlineReport(ctx, payload, virtualendpointreport.DefaultGetVirtualEndpointReportsFrontlineReportOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointReportClient.GetVirtualEndpointReportsInaccessibleCloudPCReports`

```go
ctx := context.TODO()

payload := virtualendpointreport.GetVirtualEndpointReportsInaccessibleCloudPCReportsRequest{
	// ...
}


read, err := client.GetVirtualEndpointReportsInaccessibleCloudPCReports(ctx, payload, virtualendpointreport.DefaultGetVirtualEndpointReportsInaccessibleCloudPCReportsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointReportClient.GetVirtualEndpointReportsRawRemoteConnectionReports`

```go
ctx := context.TODO()

payload := virtualendpointreport.GetVirtualEndpointReportsRawRemoteConnectionReportsRequest{
	// ...
}


read, err := client.GetVirtualEndpointReportsRawRemoteConnectionReports(ctx, payload, virtualendpointreport.DefaultGetVirtualEndpointReportsRawRemoteConnectionReportsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointReportClient.GetVirtualEndpointReportsRemoteConnectionHistoricalReports`

```go
ctx := context.TODO()

payload := virtualendpointreport.GetVirtualEndpointReportsRemoteConnectionHistoricalReportsRequest{
	// ...
}


read, err := client.GetVirtualEndpointReportsRemoteConnectionHistoricalReports(ctx, payload, virtualendpointreport.DefaultGetVirtualEndpointReportsRemoteConnectionHistoricalReportsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointReportClient.GetVirtualEndpointReportsSharedUseLicenseUsageReport`

```go
ctx := context.TODO()

payload := virtualendpointreport.GetVirtualEndpointReportsSharedUseLicenseUsageReportRequest{
	// ...
}


read, err := client.GetVirtualEndpointReportsSharedUseLicenseUsageReport(ctx, payload, virtualendpointreport.DefaultGetVirtualEndpointReportsSharedUseLicenseUsageReportOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointReportClient.GetVirtualEndpointReportsTotalAggregatedRemoteConnectionReports`

```go
ctx := context.TODO()

payload := virtualendpointreport.GetVirtualEndpointReportsTotalAggregatedRemoteConnectionReportsRequest{
	// ...
}


read, err := client.GetVirtualEndpointReportsTotalAggregatedRemoteConnectionReports(ctx, payload, virtualendpointreport.DefaultGetVirtualEndpointReportsTotalAggregatedRemoteConnectionReportsOperationOptions())
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


read, err := client.UpdateVirtualEndpointReport(ctx, payload, virtualendpointreport.DefaultUpdateVirtualEndpointReportOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
