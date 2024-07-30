
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/report` Documentation

The `report` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/report"
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


### Example Usage: `ReportClient.GetDeviceManagementReportsActiveMalwareReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsActiveMalwareReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsActiveMalwareReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsActiveMalwareSummaryReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsActiveMalwareSummaryReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsActiveMalwareSummaryReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsAllCertificatesReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsAllCertificatesReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsAllCertificatesReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsAppStatusOverviewReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsAppStatusOverviewReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsAppStatusOverviewReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsAppsInstallSummaryReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsAppsInstallSummaryReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsAppsInstallSummaryReport(ctx, payload)
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


### Example Usage: `ReportClient.GetDeviceManagementReportsCertificatesReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsCertificatesReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsCertificatesReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsCompliancePoliciesReportForDevice`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsCompliancePoliciesReportForDeviceRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsCompliancePoliciesReportForDevice(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsCompliancePolicyDeviceSummaryReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsCompliancePolicyDeviceSummaryReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsCompliancePolicyDeviceSummaryReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsCompliancePolicyDevicesReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsCompliancePolicyDevicesReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsCompliancePolicyDevicesReport(ctx, payload)
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


### Example Usage: `ReportClient.GetDeviceManagementReportsComplianceSettingDetailsReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsComplianceSettingDetailsReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsComplianceSettingDetailsReport(ctx, payload)
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


### Example Usage: `ReportClient.GetDeviceManagementReportsComplianceSettingsReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsComplianceSettingsReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsComplianceSettingsReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsConfigManagerDevicePolicyStatusReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsConfigManagerDevicePolicyStatusReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsConfigManagerDevicePolicyStatusReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsConfigurationPoliciesReportForDevice`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsConfigurationPoliciesReportForDeviceRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsConfigurationPoliciesReportForDevice(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsConfigurationPolicyDeviceSummaryReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsConfigurationPolicyDeviceSummaryReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsConfigurationPolicyDeviceSummaryReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsConfigurationPolicyDevicesReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsConfigurationPolicyDevicesReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsConfigurationPolicyDevicesReport(ctx, payload)
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


### Example Usage: `ReportClient.GetDeviceManagementReportsConfigurationPolicySettingsDeviceSummaryReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsConfigurationPolicySettingsDeviceSummaryReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsConfigurationPolicySettingsDeviceSummaryReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsConfigurationSettingDetailsReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsConfigurationSettingDetailsReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsConfigurationSettingDetailsReport(ctx, payload)
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


### Example Usage: `ReportClient.GetDeviceManagementReportsConfigurationSettingsReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsConfigurationSettingsReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsConfigurationSettingsReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsDeviceConfigurationPolicySettingsSummaryReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsDeviceConfigurationPolicySettingsSummaryReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsDeviceConfigurationPolicySettingsSummaryReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsDeviceConfigurationPolicyStatusSummary`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsDeviceConfigurationPolicyStatusSummaryRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsDeviceConfigurationPolicyStatusSummary(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsDeviceInstallStatusReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsDeviceInstallStatusReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsDeviceInstallStatusReport(ctx, payload)
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


### Example Usage: `ReportClient.GetDeviceManagementReportsDevicePoliciesComplianceReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsDevicePoliciesComplianceReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsDevicePoliciesComplianceReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsDevicePolicySettingsComplianceReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsDevicePolicySettingsComplianceReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsDevicePolicySettingsComplianceReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsDeviceStatusByCompliacePolicyReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsDeviceStatusByCompliacePolicyReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsDeviceStatusByCompliacePolicyReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsDeviceStatusByCompliancePolicySettingReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsDeviceStatusByCompliancePolicySettingReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsDeviceStatusByCompliancePolicySettingReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsDeviceStatusSummaryByCompliacePolicyReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsDeviceStatusSummaryByCompliacePolicyReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsDeviceStatusSummaryByCompliacePolicyReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsDeviceStatusSummaryByCompliancePolicySettingsReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsDeviceStatusSummaryByCompliancePolicySettingsReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsDeviceStatusSummaryByCompliancePolicySettingsReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsDevicesStatusByPolicyPlatformComplianceReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsDevicesStatusByPolicyPlatformComplianceReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsDevicesStatusByPolicyPlatformComplianceReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsDevicesStatusBySettingReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsDevicesStatusBySettingReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsDevicesStatusBySettingReport(ctx, payload)
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


### Example Usage: `ReportClient.GetDeviceManagementReportsEncryptionReportForDevice`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsEncryptionReportForDeviceRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsEncryptionReportForDevice(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsEnrollmentConfigurationPoliciesByDevice`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsEnrollmentConfigurationPoliciesByDeviceRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsEnrollmentConfigurationPoliciesByDevice(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsFailedMobileAppsReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsFailedMobileAppsReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsFailedMobileAppsReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsFailedMobileAppsSummaryReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsFailedMobileAppsSummaryReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsFailedMobileAppsSummaryReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsGroupPolicySettingsDeviceSettingsReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsGroupPolicySettingsDeviceSettingsReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsGroupPolicySettingsDeviceSettingsReport(ctx, payload)
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


### Example Usage: `ReportClient.GetDeviceManagementReportsMalwareSummaryReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsMalwareSummaryReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsMalwareSummaryReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsMobileApplicationManagementAppConfigurationReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsMobileApplicationManagementAppConfigurationReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsMobileApplicationManagementAppConfigurationReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsMobileApplicationManagementAppRegistrationSummaryReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsMobileApplicationManagementAppRegistrationSummaryReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsMobileApplicationManagementAppRegistrationSummaryReport(ctx, payload)
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


### Example Usage: `ReportClient.GetDeviceManagementReportsQuietTimePolicyUserSummaryReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsQuietTimePolicyUserSummaryReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsQuietTimePolicyUserSummaryReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsQuietTimePolicyUsersReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsQuietTimePolicyUsersReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsQuietTimePolicyUsersReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsRelatedAppsStatusReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsRelatedAppsStatusReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsRelatedAppsStatusReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsRemoteAssistanceSessionsReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsRemoteAssistanceSessionsReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsRemoteAssistanceSessionsReport(ctx, payload)
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


### Example Usage: `ReportClient.GetDeviceManagementReportsUnhealthyDefenderAgentsReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsUnhealthyDefenderAgentsReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsUnhealthyDefenderAgentsReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsUnhealthyFirewallReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsUnhealthyFirewallReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsUnhealthyFirewallReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsUnhealthyFirewallSummaryReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsUnhealthyFirewallSummaryReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsUnhealthyFirewallSummaryReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsUserInstallStatusReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsUserInstallStatusReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsUserInstallStatusReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsWindowsDriverUpdateAlertSummaryReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsWindowsDriverUpdateAlertSummaryReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsWindowsDriverUpdateAlertSummaryReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsWindowsDriverUpdateAlertsPerPolicyPerDeviceReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsWindowsDriverUpdateAlertsPerPolicyPerDeviceReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsWindowsDriverUpdateAlertsPerPolicyPerDeviceReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsWindowsQualityUpdateAlertSummaryReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsWindowsQualityUpdateAlertSummaryReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsWindowsQualityUpdateAlertSummaryReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsWindowsQualityUpdateAlertsPerPolicyPerDeviceReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsWindowsQualityUpdateAlertsPerPolicyPerDeviceReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsWindowsUpdateAlertSummaryReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsWindowsUpdateAlertSummaryReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsWindowsUpdateAlertSummaryReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsWindowsUpdateAlertsPerPolicyPerDeviceReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsWindowsUpdateAlertsPerPolicyPerDeviceReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsWindowsUpdateAlertsPerPolicyPerDeviceReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetDeviceManagementReportsZebraFotaDeploymentReport`

```go
ctx := context.TODO()

payload := report.GetDeviceManagementReportsZebraFotaDeploymentReportRequest{
	// ...
}


read, err := client.GetDeviceManagementReportsZebraFotaDeploymentReport(ctx, payload)
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
