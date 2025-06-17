
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/report` Documentation

The `report` SDK allows for interaction with Microsoft Graph `devicemanagement` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/report"
```


### Client Initialization

```go
client := report.NewReportClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ReportClient.CreateReportRetrieveDeviceAppInstallationStatusReport`

```go
ctx := context.TODO()

payload := report.CreateReportRetrieveDeviceAppInstallationStatusReportRequest{
	// ...
}


read, err := client.CreateReportRetrieveDeviceAppInstallationStatusReport(ctx, payload, report.DefaultCreateReportRetrieveDeviceAppInstallationStatusReportOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.DeleteReport`

```go
ctx := context.TODO()


read, err := client.DeleteReport(ctx, report.DefaultDeleteReportOperationOptions())
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


read, err := client.GetReport(ctx, report.DefaultGetReportOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsCachedReport`

```go
ctx := context.TODO()

payload := report.GetReportsCachedReportRequest{
	// ...
}


read, err := client.GetReportsCachedReport(ctx, payload, report.DefaultGetReportsCachedReportOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsCompliancePolicyNonComplianceReport`

```go
ctx := context.TODO()

payload := report.GetReportsCompliancePolicyNonComplianceReportRequest{
	// ...
}


read, err := client.GetReportsCompliancePolicyNonComplianceReport(ctx, payload, report.DefaultGetReportsCompliancePolicyNonComplianceReportOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsCompliancePolicyNonComplianceSummaryReport`

```go
ctx := context.TODO()

payload := report.GetReportsCompliancePolicyNonComplianceSummaryReportRequest{
	// ...
}


read, err := client.GetReportsCompliancePolicyNonComplianceSummaryReport(ctx, payload, report.DefaultGetReportsCompliancePolicyNonComplianceSummaryReportOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsComplianceSettingNonComplianceReport`

```go
ctx := context.TODO()

payload := report.GetReportsComplianceSettingNonComplianceReportRequest{
	// ...
}


read, err := client.GetReportsComplianceSettingNonComplianceReport(ctx, payload, report.DefaultGetReportsComplianceSettingNonComplianceReportOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsConfigurationPolicyNonComplianceReport`

```go
ctx := context.TODO()

payload := report.GetReportsConfigurationPolicyNonComplianceReportRequest{
	// ...
}


read, err := client.GetReportsConfigurationPolicyNonComplianceReport(ctx, payload, report.DefaultGetReportsConfigurationPolicyNonComplianceReportOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsConfigurationPolicyNonComplianceSummaryReport`

```go
ctx := context.TODO()

payload := report.GetReportsConfigurationPolicyNonComplianceSummaryReportRequest{
	// ...
}


read, err := client.GetReportsConfigurationPolicyNonComplianceSummaryReport(ctx, payload, report.DefaultGetReportsConfigurationPolicyNonComplianceSummaryReportOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsConfigurationSettingNonComplianceReport`

```go
ctx := context.TODO()

payload := report.GetReportsConfigurationSettingNonComplianceReportRequest{
	// ...
}


read, err := client.GetReportsConfigurationSettingNonComplianceReport(ctx, payload, report.DefaultGetReportsConfigurationSettingNonComplianceReportOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsDeviceManagementIntentPerSettingContributingProfiles`

```go
ctx := context.TODO()

payload := report.GetReportsDeviceManagementIntentPerSettingContributingProfilesRequest{
	// ...
}


read, err := client.GetReportsDeviceManagementIntentPerSettingContributingProfiles(ctx, payload, report.DefaultGetReportsDeviceManagementIntentPerSettingContributingProfilesOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsDeviceManagementIntentSettingsReport`

```go
ctx := context.TODO()

payload := report.GetReportsDeviceManagementIntentSettingsReportRequest{
	// ...
}


read, err := client.GetReportsDeviceManagementIntentSettingsReport(ctx, payload, report.DefaultGetReportsDeviceManagementIntentSettingsReportOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsDeviceNonComplianceReport`

```go
ctx := context.TODO()

payload := report.GetReportsDeviceNonComplianceReportRequest{
	// ...
}


read, err := client.GetReportsDeviceNonComplianceReport(ctx, payload, report.DefaultGetReportsDeviceNonComplianceReportOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsDevicesWithoutCompliancePolicyReport`

```go
ctx := context.TODO()

payload := report.GetReportsDevicesWithoutCompliancePolicyReportRequest{
	// ...
}


read, err := client.GetReportsDevicesWithoutCompliancePolicyReport(ctx, payload, report.DefaultGetReportsDevicesWithoutCompliancePolicyReportOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsFilters`

```go
ctx := context.TODO()

payload := report.GetReportsFiltersRequest{
	// ...
}


read, err := client.GetReportsFilters(ctx, payload, report.DefaultGetReportsFiltersOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsHistoricalReport`

```go
ctx := context.TODO()

payload := report.GetReportsHistoricalReportRequest{
	// ...
}


read, err := client.GetReportsHistoricalReport(ctx, payload, report.DefaultGetReportsHistoricalReportOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsNoncompliantDevicesAndSettingsReport`

```go
ctx := context.TODO()

payload := report.GetReportsNoncompliantDevicesAndSettingsReportRequest{
	// ...
}


read, err := client.GetReportsNoncompliantDevicesAndSettingsReport(ctx, payload, report.DefaultGetReportsNoncompliantDevicesAndSettingsReportOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsPolicyNonComplianceMetadata`

```go
ctx := context.TODO()

payload := report.GetReportsPolicyNonComplianceMetadataRequest{
	// ...
}


read, err := client.GetReportsPolicyNonComplianceMetadata(ctx, payload, report.DefaultGetReportsPolicyNonComplianceMetadataOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsPolicyNonComplianceReport`

```go
ctx := context.TODO()

payload := report.GetReportsPolicyNonComplianceReportRequest{
	// ...
}


read, err := client.GetReportsPolicyNonComplianceReport(ctx, payload, report.DefaultGetReportsPolicyNonComplianceReportOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsPolicyNonComplianceSummaryReport`

```go
ctx := context.TODO()

payload := report.GetReportsPolicyNonComplianceSummaryReportRequest{
	// ...
}


read, err := client.GetReportsPolicyNonComplianceSummaryReport(ctx, payload, report.DefaultGetReportsPolicyNonComplianceSummaryReportOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsSettingNonComplianceReport`

```go
ctx := context.TODO()

payload := report.GetReportsSettingNonComplianceReportRequest{
	// ...
}


read, err := client.GetReportsSettingNonComplianceReport(ctx, payload, report.DefaultGetReportsSettingNonComplianceReportOperationOptions())
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


read, err := client.UpdateReport(ctx, payload, report.DefaultUpdateReportOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
