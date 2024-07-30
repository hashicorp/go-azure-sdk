
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/report` Documentation

The `report` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/report"
```


### Client Initialization

```go
client := report.NewReportClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ReportClient.DeleteReport`

```go
ctx := context.TODO()


read, err := client.DeleteReport(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsCachedReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsCachedReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsCachedReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsCompliancePolicyNonComplianceReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsCompliancePolicyNonComplianceReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsCompliancePolicyNonComplianceReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsCompliancePolicyNonComplianceSummaryReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsCompliancePolicyNonComplianceSummaryReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsCompliancePolicyNonComplianceSummaryReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsComplianceSettingNonComplianceReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsComplianceSettingNonComplianceReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsComplianceSettingNonComplianceReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsConfigurationPolicyNonComplianceReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsConfigurationPolicyNonComplianceReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsConfigurationPolicyNonComplianceReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsConfigurationPolicyNonComplianceSummaryReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsConfigurationPolicyNonComplianceSummaryReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsConfigurationPolicyNonComplianceSummaryReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsConfigurationSettingNonComplianceReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsConfigurationSettingNonComplianceReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsConfigurationSettingNonComplianceReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsDeviceManagementIntentPerSettingContributingProfile`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsDeviceManagementIntentPerSettingContributingProfileRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsDeviceManagementIntentPerSettingContributingProfile(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsDeviceManagementIntentSettingsReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsDeviceManagementIntentSettingsReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsDeviceManagementIntentSettingsReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsDeviceNonComplianceReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsDeviceNonComplianceReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsDeviceNonComplianceReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsDevicesWithoutCompliancePolicyReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsDevicesWithoutCompliancePolicyReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsDevicesWithoutCompliancePolicyReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsHistoricalReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsHistoricalReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsHistoricalReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsNoncompliantDevicesAndSettingsReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsNoncompliantDevicesAndSettingsReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsNoncompliantDevicesAndSettingsReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsPolicyNonComplianceMetadata`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsPolicyNonComplianceMetadataRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsPolicyNonComplianceMetadata(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsPolicyNonComplianceReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsPolicyNonComplianceReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsPolicyNonComplianceReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsPolicyNonComplianceSummaryReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsPolicyNonComplianceSummaryReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsPolicyNonComplianceSummaryReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsReportFilter`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsReportFilterRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsReportFilter(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsSettingNonComplianceReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsSettingNonComplianceReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsSettingNonComplianceReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReport`

```go
ctx := context.TODO()


read, err := client.GetReport(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.UpdateReport`

```go
ctx := context.TODO()

payload := report.DeviceManagementReports{
	// ...
}


read, err := client.UpdateReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
