
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


### Example Usage: `ReportClient.CreateReportRetrieveAssignedApplicationsReport`

```go
ctx := context.TODO()

payload := report.CreateReportRetrieveAssignedApplicationsReportRequest{
	// ...
}


read, err := client.CreateReportRetrieveAssignedApplicationsReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.CreateReportRetrieveWin32CatalogAppsUpdateReport`

```go
ctx := context.TODO()

payload := report.CreateReportRetrieveWin32CatalogAppsUpdateReportRequest{
	// ...
}


read, err := client.CreateReportRetrieveWin32CatalogAppsUpdateReport(ctx, payload)
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


### Example Usage: `ReportClient.GetReportsActiveMalwareReport`

```go
ctx := context.TODO()

payload := report.GetReportsActiveMalwareReportRequest{
	// ...
}


read, err := client.GetReportsActiveMalwareReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsActiveMalwareSummaryReport`

```go
ctx := context.TODO()

payload := report.GetReportsActiveMalwareSummaryReportRequest{
	// ...
}


read, err := client.GetReportsActiveMalwareSummaryReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsAllCertificatesReport`

```go
ctx := context.TODO()

payload := report.GetReportsAllCertificatesReportRequest{
	// ...
}


read, err := client.GetReportsAllCertificatesReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsAppStatusOverviewReport`

```go
ctx := context.TODO()

payload := report.GetReportsAppStatusOverviewReportRequest{
	// ...
}


read, err := client.GetReportsAppStatusOverviewReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsAppsInstallSummaryReport`

```go
ctx := context.TODO()

payload := report.GetReportsAppsInstallSummaryReportRequest{
	// ...
}


read, err := client.GetReportsAppsInstallSummaryReport(ctx, payload)
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


read, err := client.GetReportsCachedReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsCertificatesReport`

```go
ctx := context.TODO()

payload := report.GetReportsCertificatesReportRequest{
	// ...
}


read, err := client.GetReportsCertificatesReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsCompliancePoliciesReportForDevice`

```go
ctx := context.TODO()

payload := report.GetReportsCompliancePoliciesReportForDeviceRequest{
	// ...
}


read, err := client.GetReportsCompliancePoliciesReportForDevice(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsCompliancePolicyDeviceSummaryReport`

```go
ctx := context.TODO()

payload := report.GetReportsCompliancePolicyDeviceSummaryReportRequest{
	// ...
}


read, err := client.GetReportsCompliancePolicyDeviceSummaryReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsCompliancePolicyDevicesReport`

```go
ctx := context.TODO()

payload := report.GetReportsCompliancePolicyDevicesReportRequest{
	// ...
}


read, err := client.GetReportsCompliancePolicyDevicesReport(ctx, payload)
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


read, err := client.GetReportsCompliancePolicyNonComplianceReport(ctx, payload)
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


read, err := client.GetReportsCompliancePolicyNonComplianceSummaryReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsComplianceSettingDetailsReport`

```go
ctx := context.TODO()

payload := report.GetReportsComplianceSettingDetailsReportRequest{
	// ...
}


read, err := client.GetReportsComplianceSettingDetailsReport(ctx, payload)
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


read, err := client.GetReportsComplianceSettingNonComplianceReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsComplianceSettingsReport`

```go
ctx := context.TODO()

payload := report.GetReportsComplianceSettingsReportRequest{
	// ...
}


read, err := client.GetReportsComplianceSettingsReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsConfigManagerDevicePolicyStatusReport`

```go
ctx := context.TODO()

payload := report.GetReportsConfigManagerDevicePolicyStatusReportRequest{
	// ...
}


read, err := client.GetReportsConfigManagerDevicePolicyStatusReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsConfigurationPoliciesReportForDevice`

```go
ctx := context.TODO()

payload := report.GetReportsConfigurationPoliciesReportForDeviceRequest{
	// ...
}


read, err := client.GetReportsConfigurationPoliciesReportForDevice(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsConfigurationPolicyDeviceSummaryReport`

```go
ctx := context.TODO()

payload := report.GetReportsConfigurationPolicyDeviceSummaryReportRequest{
	// ...
}


read, err := client.GetReportsConfigurationPolicyDeviceSummaryReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsConfigurationPolicyDevicesReport`

```go
ctx := context.TODO()

payload := report.GetReportsConfigurationPolicyDevicesReportRequest{
	// ...
}


read, err := client.GetReportsConfigurationPolicyDevicesReport(ctx, payload)
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


read, err := client.GetReportsConfigurationPolicyNonComplianceReport(ctx, payload)
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


read, err := client.GetReportsConfigurationPolicyNonComplianceSummaryReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsConfigurationPolicySettingsDeviceSummaryReport`

```go
ctx := context.TODO()

payload := report.GetReportsConfigurationPolicySettingsDeviceSummaryReportRequest{
	// ...
}


read, err := client.GetReportsConfigurationPolicySettingsDeviceSummaryReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsConfigurationSettingDetailsReport`

```go
ctx := context.TODO()

payload := report.GetReportsConfigurationSettingDetailsReportRequest{
	// ...
}


read, err := client.GetReportsConfigurationSettingDetailsReport(ctx, payload)
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


read, err := client.GetReportsConfigurationSettingNonComplianceReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsConfigurationSettingsReport`

```go
ctx := context.TODO()

payload := report.GetReportsConfigurationSettingsReportRequest{
	// ...
}


read, err := client.GetReportsConfigurationSettingsReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsDeviceConfigurationPolicySettingsSummaryReport`

```go
ctx := context.TODO()

payload := report.GetReportsDeviceConfigurationPolicySettingsSummaryReportRequest{
	// ...
}


read, err := client.GetReportsDeviceConfigurationPolicySettingsSummaryReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsDeviceConfigurationPolicyStatusSummary`

```go
ctx := context.TODO()

payload := report.GetReportsDeviceConfigurationPolicyStatusSummaryRequest{
	// ...
}


read, err := client.GetReportsDeviceConfigurationPolicyStatusSummary(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsDeviceInstallStatusReport`

```go
ctx := context.TODO()

payload := report.GetReportsDeviceInstallStatusReportRequest{
	// ...
}


read, err := client.GetReportsDeviceInstallStatusReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsDeviceManagementIntentPerSettingContributingProfile`

```go
ctx := context.TODO()

payload := report.GetReportsDeviceManagementIntentPerSettingContributingProfileRequest{
	// ...
}


read, err := client.GetReportsDeviceManagementIntentPerSettingContributingProfile(ctx, payload)
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


read, err := client.GetReportsDeviceManagementIntentSettingsReport(ctx, payload)
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


read, err := client.GetReportsDeviceNonComplianceReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsDevicePoliciesComplianceReport`

```go
ctx := context.TODO()

payload := report.GetReportsDevicePoliciesComplianceReportRequest{
	// ...
}


read, err := client.GetReportsDevicePoliciesComplianceReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsDevicePolicySettingsComplianceReport`

```go
ctx := context.TODO()

payload := report.GetReportsDevicePolicySettingsComplianceReportRequest{
	// ...
}


read, err := client.GetReportsDevicePolicySettingsComplianceReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsDeviceStatusByCompliacePolicyReport`

```go
ctx := context.TODO()

payload := report.GetReportsDeviceStatusByCompliacePolicyReportRequest{
	// ...
}


read, err := client.GetReportsDeviceStatusByCompliacePolicyReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsDeviceStatusByCompliancePolicySettingReport`

```go
ctx := context.TODO()

payload := report.GetReportsDeviceStatusByCompliancePolicySettingReportRequest{
	// ...
}


read, err := client.GetReportsDeviceStatusByCompliancePolicySettingReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsDeviceStatusSummaryByCompliacePolicyReport`

```go
ctx := context.TODO()

payload := report.GetReportsDeviceStatusSummaryByCompliacePolicyReportRequest{
	// ...
}


read, err := client.GetReportsDeviceStatusSummaryByCompliacePolicyReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsDeviceStatusSummaryByCompliancePolicySettingsReport`

```go
ctx := context.TODO()

payload := report.GetReportsDeviceStatusSummaryByCompliancePolicySettingsReportRequest{
	// ...
}


read, err := client.GetReportsDeviceStatusSummaryByCompliancePolicySettingsReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsDevicesStatusByPolicyPlatformComplianceReport`

```go
ctx := context.TODO()

payload := report.GetReportsDevicesStatusByPolicyPlatformComplianceReportRequest{
	// ...
}


read, err := client.GetReportsDevicesStatusByPolicyPlatformComplianceReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsDevicesStatusBySettingReport`

```go
ctx := context.TODO()

payload := report.GetReportsDevicesStatusBySettingReportRequest{
	// ...
}


read, err := client.GetReportsDevicesStatusBySettingReport(ctx, payload)
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


read, err := client.GetReportsDevicesWithoutCompliancePolicyReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsEncryptionReportForDevice`

```go
ctx := context.TODO()

payload := report.GetReportsEncryptionReportForDeviceRequest{
	// ...
}


read, err := client.GetReportsEncryptionReportForDevice(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsEnrollmentConfigurationPoliciesByDevice`

```go
ctx := context.TODO()

payload := report.GetReportsEnrollmentConfigurationPoliciesByDeviceRequest{
	// ...
}


read, err := client.GetReportsEnrollmentConfigurationPoliciesByDevice(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsFailedMobileAppsReport`

```go
ctx := context.TODO()

payload := report.GetReportsFailedMobileAppsReportRequest{
	// ...
}


read, err := client.GetReportsFailedMobileAppsReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsFailedMobileAppsSummaryReport`

```go
ctx := context.TODO()

payload := report.GetReportsFailedMobileAppsSummaryReportRequest{
	// ...
}


read, err := client.GetReportsFailedMobileAppsSummaryReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsFilter`

```go
ctx := context.TODO()

payload := report.GetReportsFilterRequest{
	// ...
}


read, err := client.GetReportsFilter(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsGroupPolicySettingsDeviceSettingsReport`

```go
ctx := context.TODO()

payload := report.GetReportsGroupPolicySettingsDeviceSettingsReportRequest{
	// ...
}


read, err := client.GetReportsGroupPolicySettingsDeviceSettingsReport(ctx, payload)
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


read, err := client.GetReportsHistoricalReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsMalwareSummaryReport`

```go
ctx := context.TODO()

payload := report.GetReportsMalwareSummaryReportRequest{
	// ...
}


read, err := client.GetReportsMalwareSummaryReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsMobileApplicationManagementAppConfigurationReport`

```go
ctx := context.TODO()

payload := report.GetReportsMobileApplicationManagementAppConfigurationReportRequest{
	// ...
}


read, err := client.GetReportsMobileApplicationManagementAppConfigurationReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsMobileApplicationManagementAppRegistrationSummaryReport`

```go
ctx := context.TODO()

payload := report.GetReportsMobileApplicationManagementAppRegistrationSummaryReportRequest{
	// ...
}


read, err := client.GetReportsMobileApplicationManagementAppRegistrationSummaryReport(ctx, payload)
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


read, err := client.GetReportsNoncompliantDevicesAndSettingsReport(ctx, payload)
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


read, err := client.GetReportsPolicyNonComplianceMetadata(ctx, payload)
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


read, err := client.GetReportsPolicyNonComplianceReport(ctx, payload)
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


read, err := client.GetReportsPolicyNonComplianceSummaryReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsQuietTimePolicyUserSummaryReport`

```go
ctx := context.TODO()

payload := report.GetReportsQuietTimePolicyUserSummaryReportRequest{
	// ...
}


read, err := client.GetReportsQuietTimePolicyUserSummaryReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsQuietTimePolicyUsersReport`

```go
ctx := context.TODO()

payload := report.GetReportsQuietTimePolicyUsersReportRequest{
	// ...
}


read, err := client.GetReportsQuietTimePolicyUsersReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsRelatedAppsStatusReport`

```go
ctx := context.TODO()

payload := report.GetReportsRelatedAppsStatusReportRequest{
	// ...
}


read, err := client.GetReportsRelatedAppsStatusReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsRemoteAssistanceSessionsReport`

```go
ctx := context.TODO()

payload := report.GetReportsRemoteAssistanceSessionsReportRequest{
	// ...
}


read, err := client.GetReportsRemoteAssistanceSessionsReport(ctx, payload)
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


read, err := client.GetReportsSettingNonComplianceReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsUnhealthyDefenderAgentsReport`

```go
ctx := context.TODO()

payload := report.GetReportsUnhealthyDefenderAgentsReportRequest{
	// ...
}


read, err := client.GetReportsUnhealthyDefenderAgentsReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsUnhealthyFirewallReport`

```go
ctx := context.TODO()

payload := report.GetReportsUnhealthyFirewallReportRequest{
	// ...
}


read, err := client.GetReportsUnhealthyFirewallReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsUnhealthyFirewallSummaryReport`

```go
ctx := context.TODO()

payload := report.GetReportsUnhealthyFirewallSummaryReportRequest{
	// ...
}


read, err := client.GetReportsUnhealthyFirewallSummaryReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsUserInstallStatusReport`

```go
ctx := context.TODO()

payload := report.GetReportsUserInstallStatusReportRequest{
	// ...
}


read, err := client.GetReportsUserInstallStatusReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsWindowsDriverUpdateAlertSummaryReport`

```go
ctx := context.TODO()

payload := report.GetReportsWindowsDriverUpdateAlertSummaryReportRequest{
	// ...
}


read, err := client.GetReportsWindowsDriverUpdateAlertSummaryReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsWindowsDriverUpdateAlertsPerPolicyPerDeviceReport`

```go
ctx := context.TODO()

payload := report.GetReportsWindowsDriverUpdateAlertsPerPolicyPerDeviceReportRequest{
	// ...
}


read, err := client.GetReportsWindowsDriverUpdateAlertsPerPolicyPerDeviceReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsWindowsQualityUpdateAlertSummaryReport`

```go
ctx := context.TODO()

payload := report.GetReportsWindowsQualityUpdateAlertSummaryReportRequest{
	// ...
}


read, err := client.GetReportsWindowsQualityUpdateAlertSummaryReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsWindowsQualityUpdateAlertsPerPolicyPerDeviceReport`

```go
ctx := context.TODO()

payload := report.GetReportsWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequest{
	// ...
}


read, err := client.GetReportsWindowsQualityUpdateAlertsPerPolicyPerDeviceReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsWindowsUpdateAlertSummaryReport`

```go
ctx := context.TODO()

payload := report.GetReportsWindowsUpdateAlertSummaryReportRequest{
	// ...
}


read, err := client.GetReportsWindowsUpdateAlertSummaryReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsWindowsUpdateAlertsPerPolicyPerDeviceReport`

```go
ctx := context.TODO()

payload := report.GetReportsWindowsUpdateAlertsPerPolicyPerDeviceReportRequest{
	// ...
}


read, err := client.GetReportsWindowsUpdateAlertsPerPolicyPerDeviceReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReportClient.GetReportsZebraFotaDeploymentReport`

```go
ctx := context.TODO()

payload := report.GetReportsZebraFotaDeploymentReportRequest{
	// ...
}


read, err := client.GetReportsZebraFotaDeploymentReport(ctx, payload)
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
