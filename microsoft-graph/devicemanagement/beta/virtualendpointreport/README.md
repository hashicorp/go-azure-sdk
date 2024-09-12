
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


### Example Usage: `VirtualEndpointReportClient.CreateVirtualEndpointReportRetrieveCrossRegionDisasterRecoveryReport`

```go
ctx := context.TODO()

payload := virtualendpointreport.CreateVirtualEndpointReportRetrieveCrossRegionDisasterRecoveryReportRequest{
	// ...
}


read, err := client.CreateVirtualEndpointReportRetrieveCrossRegionDisasterRecoveryReport(ctx, payload)
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


### Example Usage: `VirtualEndpointReportClient.GetVirtualEndpointReportsActionStatusReport`

```go
ctx := context.TODO()

payload := virtualendpointreport.GetVirtualEndpointReportsActionStatusReportRequest{
	// ...
}


read, err := client.GetVirtualEndpointReportsActionStatusReport(ctx, payload)
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


read, err := client.GetVirtualEndpointReportsCloudPCPerformanceReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointReportClient.GetVirtualEndpointReportsCloudPCRecommendationReport`

```go
ctx := context.TODO()

payload := virtualendpointreport.GetVirtualEndpointReportsCloudPCRecommendationReportRequest{
	// ...
}


read, err := client.GetVirtualEndpointReportsCloudPCRecommendationReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointReportClient.GetVirtualEndpointReportsConnectionQualityReport`

```go
ctx := context.TODO()

payload := virtualendpointreport.GetVirtualEndpointReportsConnectionQualityReportRequest{
	// ...
}


read, err := client.GetVirtualEndpointReportsConnectionQualityReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointReportClient.GetVirtualEndpointReportsDailyAggregatedRemoteConnectionReport`

```go
ctx := context.TODO()

payload := virtualendpointreport.GetVirtualEndpointReportsDailyAggregatedRemoteConnectionReportRequest{
	// ...
}


read, err := client.GetVirtualEndpointReportsDailyAggregatedRemoteConnectionReport(ctx, payload)
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


read, err := client.GetVirtualEndpointReportsFrontlineReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointReportClient.GetVirtualEndpointReportsInaccessibleCloudPCReport`

```go
ctx := context.TODO()

payload := virtualendpointreport.GetVirtualEndpointReportsInaccessibleCloudPCReportRequest{
	// ...
}


read, err := client.GetVirtualEndpointReportsInaccessibleCloudPCReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointReportClient.GetVirtualEndpointReportsRawRemoteConnectionReport`

```go
ctx := context.TODO()

payload := virtualendpointreport.GetVirtualEndpointReportsRawRemoteConnectionReportRequest{
	// ...
}


read, err := client.GetVirtualEndpointReportsRawRemoteConnectionReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointReportClient.GetVirtualEndpointReportsRemoteConnectionHistoricalReport`

```go
ctx := context.TODO()

payload := virtualendpointreport.GetVirtualEndpointReportsRemoteConnectionHistoricalReportRequest{
	// ...
}


read, err := client.GetVirtualEndpointReportsRemoteConnectionHistoricalReport(ctx, payload)
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


read, err := client.GetVirtualEndpointReportsSharedUseLicenseUsageReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointReportClient.GetVirtualEndpointReportsTotalAggregatedRemoteConnectionReport`

```go
ctx := context.TODO()

payload := virtualendpointreport.GetVirtualEndpointReportsTotalAggregatedRemoteConnectionReportRequest{
	// ...
}


read, err := client.GetVirtualEndpointReportsTotalAggregatedRemoteConnectionReport(ctx, payload)
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
