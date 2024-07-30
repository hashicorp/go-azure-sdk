package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/applepushnotificationcertificate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/auditevent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/compliancemanagementpartner"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/conditionalaccesssetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/detectedapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/detectedappmanageddevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/devicecategory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/devicecompliancepolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/devicecompliancepolicyassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/devicecompliancepolicydevicesettingstatesummary"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/devicecompliancepolicydevicestatesummary"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/devicecompliancepolicydevicestatus"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/devicecompliancepolicydevicestatusoverview"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/devicecompliancepolicyscheduledactionsforrule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/devicecompliancepolicyscheduledactionsforrulescheduledactionconfiguration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/devicecompliancepolicysettingstatesummary"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/devicecompliancepolicysettingstatesummarydevicecompliancesettingstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/devicecompliancepolicyuserstatus"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/devicecompliancepolicyuserstatusoverview"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/deviceconfiguration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/deviceconfigurationassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/deviceconfigurationdevicesettingstatesummary"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/deviceconfigurationdevicestatesummary"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/deviceconfigurationdevicestatus"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/deviceconfigurationdevicestatusoverview"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/deviceconfigurationuserstatus"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/deviceconfigurationuserstatusoverview"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/deviceenrollmentconfiguration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/deviceenrollmentconfigurationassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/devicemanagement"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/devicemanagementpartner"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/exchangeconnector"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/importedwindowsautopilotdeviceidentity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/iosupdatestatus"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/manageddevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/manageddevicedevicecategory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/manageddevicedevicecompliancepolicystate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/manageddevicedeviceconfigurationstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/manageddevicelogcollectionrequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/manageddeviceoverview"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/manageddeviceuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/manageddevicewindowsprotectionstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/manageddevicewindowsprotectionstatedetectedmalwarestate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/mobileapptroubleshootingevent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/mobileapptroubleshootingeventapplogcollectionrequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/mobilethreatdefenseconnector"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/notificationmessagetemplate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/notificationmessagetemplatelocalizednotificationmessage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/remoteassistancepartner"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/report"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/reportexportjob"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/resourceoperation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/roleassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/roleassignmentroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/roledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/roledefinitionroleassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/roledefinitionroleassignmentroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/softwareupdatestatussummary"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/telecomexpensemanagementpartner"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/termsandcondition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/termsandconditionacceptancestatus"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/termsandconditionacceptancestatustermsandcondition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/termsandconditionassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/troubleshootingevent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/userexperienceanalyticsapphealthapplicationperformance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/userexperienceanalyticsapphealthapplicationperformancebyappversiondetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/userexperienceanalyticsapphealthapplicationperformancebyappversiondeviceid"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/userexperienceanalyticsapphealthapplicationperformancebyosversion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/userexperienceanalyticsapphealthdevicemodelperformance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/userexperienceanalyticsapphealthdeviceperformance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/userexperienceanalyticsapphealthdeviceperformancedetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/userexperienceanalyticsapphealthosversionperformance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/userexperienceanalyticsapphealthoverview"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/userexperienceanalyticsapphealthoverviewmetricvalue"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/userexperienceanalyticsbaseline"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/userexperienceanalyticsbaselineapphealthmetric"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/userexperienceanalyticsbaselinebatteryhealthmetric"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/userexperienceanalyticsbaselinebestpracticesmetric"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/userexperienceanalyticsbaselinedevicebootperformancemetric"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/userexperienceanalyticsbaselinerebootanalyticsmetric"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/userexperienceanalyticsbaselineresourceperformancemetric"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/userexperienceanalyticsbaselineworkfromanywheremetric"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/userexperienceanalyticscategory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/userexperienceanalyticscategorymetricvalue"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/userexperienceanalyticsdeviceperformance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/userexperienceanalyticsdevicescore"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/userexperienceanalyticsdevicestartuphistory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/userexperienceanalyticsdevicestartupprocess"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/userexperienceanalyticsdevicestartupprocessperformance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/userexperienceanalyticsmetrichistory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/userexperienceanalyticsmodelscore"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/userexperienceanalyticsoverview"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/userexperienceanalyticsscorehistory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/userexperienceanalyticsworkfromanywherehardwarereadinessmetric"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/userexperienceanalyticsworkfromanywheremetric"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/userexperienceanalyticsworkfromanywheremetricmetricdevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/userexperienceanalyticsworkfromanywheremodelperformance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/virtualendpoint"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/windowsautopilotdeviceidentity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/windowsinformationprotectionapplearningsummary"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/windowsinformationprotectionnetworklearningsummary"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/windowsmalwareinformation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/windowsmalwareinformationdevicemalwarestate"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
	"github.com/manicminer/hamilton/msgraph"
)

type Client struct {
	ApplePushNotificationCertificate                                           *applepushnotificationcertificate.ApplePushNotificationCertificateClient
	AuditEvent                                                                 *auditevent.AuditEventClient
	ComplianceManagementPartner                                                *compliancemanagementpartner.ComplianceManagementPartnerClient
	ConditionalAccessSetting                                                   *conditionalaccesssetting.ConditionalAccessSettingClient
	DetectedApp                                                                *detectedapp.DetectedAppClient
	DetectedAppManagedDevice                                                   *detectedappmanageddevice.DetectedAppManagedDeviceClient
	DeviceCategory                                                             *devicecategory.DeviceCategoryClient
	DeviceCompliancePolicy                                                     *devicecompliancepolicy.DeviceCompliancePolicyClient
	DeviceCompliancePolicyAssignment                                           *devicecompliancepolicyassignment.DeviceCompliancePolicyAssignmentClient
	DeviceCompliancePolicyDeviceSettingStateSummary                            *devicecompliancepolicydevicesettingstatesummary.DeviceCompliancePolicyDeviceSettingStateSummaryClient
	DeviceCompliancePolicyDeviceStateSummary                                   *devicecompliancepolicydevicestatesummary.DeviceCompliancePolicyDeviceStateSummaryClient
	DeviceCompliancePolicyDeviceStatus                                         *devicecompliancepolicydevicestatus.DeviceCompliancePolicyDeviceStatusClient
	DeviceCompliancePolicyDeviceStatusOverview                                 *devicecompliancepolicydevicestatusoverview.DeviceCompliancePolicyDeviceStatusOverviewClient
	DeviceCompliancePolicyScheduledActionsForRule                              *devicecompliancepolicyscheduledactionsforrule.DeviceCompliancePolicyScheduledActionsForRuleClient
	DeviceCompliancePolicyScheduledActionsForRuleScheduledActionConfiguration  *devicecompliancepolicyscheduledactionsforrulescheduledactionconfiguration.DeviceCompliancePolicyScheduledActionsForRuleScheduledActionConfigurationClient
	DeviceCompliancePolicySettingStateSummary                                  *devicecompliancepolicysettingstatesummary.DeviceCompliancePolicySettingStateSummaryClient
	DeviceCompliancePolicySettingStateSummaryDeviceComplianceSettingState      *devicecompliancepolicysettingstatesummarydevicecompliancesettingstate.DeviceCompliancePolicySettingStateSummaryDeviceComplianceSettingStateClient
	DeviceCompliancePolicyUserStatus                                           *devicecompliancepolicyuserstatus.DeviceCompliancePolicyUserStatusClient
	DeviceCompliancePolicyUserStatusOverview                                   *devicecompliancepolicyuserstatusoverview.DeviceCompliancePolicyUserStatusOverviewClient
	DeviceConfiguration                                                        *deviceconfiguration.DeviceConfigurationClient
	DeviceConfigurationAssignment                                              *deviceconfigurationassignment.DeviceConfigurationAssignmentClient
	DeviceConfigurationDeviceSettingStateSummary                               *deviceconfigurationdevicesettingstatesummary.DeviceConfigurationDeviceSettingStateSummaryClient
	DeviceConfigurationDeviceStateSummary                                      *deviceconfigurationdevicestatesummary.DeviceConfigurationDeviceStateSummaryClient
	DeviceConfigurationDeviceStatus                                            *deviceconfigurationdevicestatus.DeviceConfigurationDeviceStatusClient
	DeviceConfigurationDeviceStatusOverview                                    *deviceconfigurationdevicestatusoverview.DeviceConfigurationDeviceStatusOverviewClient
	DeviceConfigurationUserStatus                                              *deviceconfigurationuserstatus.DeviceConfigurationUserStatusClient
	DeviceConfigurationUserStatusOverview                                      *deviceconfigurationuserstatusoverview.DeviceConfigurationUserStatusOverviewClient
	DeviceEnrollmentConfiguration                                              *deviceenrollmentconfiguration.DeviceEnrollmentConfigurationClient
	DeviceEnrollmentConfigurationAssignment                                    *deviceenrollmentconfigurationassignment.DeviceEnrollmentConfigurationAssignmentClient
	DeviceManagement                                                           *devicemanagement.DeviceManagementClient
	DeviceManagementPartner                                                    *devicemanagementpartner.DeviceManagementPartnerClient
	ExchangeConnector                                                          *exchangeconnector.ExchangeConnectorClient
	ImportedWindowsAutopilotDeviceIdentity                                     *importedwindowsautopilotdeviceidentity.ImportedWindowsAutopilotDeviceIdentityClient
	IosUpdateStatus                                                            *iosupdatestatus.IosUpdateStatusClient
	ManagedDevice                                                              *manageddevice.ManagedDeviceClient
	ManagedDeviceDeviceCategory                                                *manageddevicedevicecategory.ManagedDeviceDeviceCategoryClient
	ManagedDeviceDeviceCompliancePolicyState                                   *manageddevicedevicecompliancepolicystate.ManagedDeviceDeviceCompliancePolicyStateClient
	ManagedDeviceDeviceConfigurationState                                      *manageddevicedeviceconfigurationstate.ManagedDeviceDeviceConfigurationStateClient
	ManagedDeviceLogCollectionRequest                                          *manageddevicelogcollectionrequest.ManagedDeviceLogCollectionRequestClient
	ManagedDeviceOverview                                                      *manageddeviceoverview.ManagedDeviceOverviewClient
	ManagedDeviceUser                                                          *manageddeviceuser.ManagedDeviceUserClient
	ManagedDeviceWindowsProtectionState                                        *manageddevicewindowsprotectionstate.ManagedDeviceWindowsProtectionStateClient
	ManagedDeviceWindowsProtectionStateDetectedMalwareState                    *manageddevicewindowsprotectionstatedetectedmalwarestate.ManagedDeviceWindowsProtectionStateDetectedMalwareStateClient
	MobileAppTroubleshootingEvent                                              *mobileapptroubleshootingevent.MobileAppTroubleshootingEventClient
	MobileAppTroubleshootingEventAppLogCollectionRequest                       *mobileapptroubleshootingeventapplogcollectionrequest.MobileAppTroubleshootingEventAppLogCollectionRequestClient
	MobileThreatDefenseConnector                                               *mobilethreatdefenseconnector.MobileThreatDefenseConnectorClient
	NotificationMessageTemplate                                                *notificationmessagetemplate.NotificationMessageTemplateClient
	NotificationMessageTemplateLocalizedNotificationMessage                    *notificationmessagetemplatelocalizednotificationmessage.NotificationMessageTemplateLocalizedNotificationMessageClient
	RemoteAssistancePartner                                                    *remoteassistancepartner.RemoteAssistancePartnerClient
	Report                                                                     *report.ReportClient
	ReportExportJob                                                            *reportexportjob.ReportExportJobClient
	ResourceOperation                                                          *resourceoperation.ResourceOperationClient
	RoleAssignment                                                             *roleassignment.RoleAssignmentClient
	RoleAssignmentRoleDefinition                                               *roleassignmentroledefinition.RoleAssignmentRoleDefinitionClient
	RoleDefinition                                                             *roledefinition.RoleDefinitionClient
	RoleDefinitionRoleAssignment                                               *roledefinitionroleassignment.RoleDefinitionRoleAssignmentClient
	RoleDefinitionRoleAssignmentRoleDefinition                                 *roledefinitionroleassignmentroledefinition.RoleDefinitionRoleAssignmentRoleDefinitionClient
	SoftwareUpdateStatusSummary                                                *softwareupdatestatussummary.SoftwareUpdateStatusSummaryClient
	TelecomExpenseManagementPartner                                            *telecomexpensemanagementpartner.TelecomExpenseManagementPartnerClient
	TermsAndCondition                                                          *termsandcondition.TermsAndConditionClient
	TermsAndConditionAcceptanceStatus                                          *termsandconditionacceptancestatus.TermsAndConditionAcceptanceStatusClient
	TermsAndConditionAcceptanceStatusTermsAndCondition                         *termsandconditionacceptancestatustermsandcondition.TermsAndConditionAcceptanceStatusTermsAndConditionClient
	TermsAndConditionAssignment                                                *termsandconditionassignment.TermsAndConditionAssignmentClient
	TroubleshootingEvent                                                       *troubleshootingevent.TroubleshootingEventClient
	UserExperienceAnalyticsAppHealthApplicationPerformance                     *userexperienceanalyticsapphealthapplicationperformance.UserExperienceAnalyticsAppHealthApplicationPerformanceClient
	UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetail   *userexperienceanalyticsapphealthapplicationperformancebyappversiondetail.UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailClient
	UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId *userexperienceanalyticsapphealthapplicationperformancebyappversiondeviceid.UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdClient
	UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion          *userexperienceanalyticsapphealthapplicationperformancebyosversion.UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionClient
	UserExperienceAnalyticsAppHealthDeviceModelPerformance                     *userexperienceanalyticsapphealthdevicemodelperformance.UserExperienceAnalyticsAppHealthDeviceModelPerformanceClient
	UserExperienceAnalyticsAppHealthDevicePerformance                          *userexperienceanalyticsapphealthdeviceperformance.UserExperienceAnalyticsAppHealthDevicePerformanceClient
	UserExperienceAnalyticsAppHealthDevicePerformanceDetail                    *userexperienceanalyticsapphealthdeviceperformancedetail.UserExperienceAnalyticsAppHealthDevicePerformanceDetailClient
	UserExperienceAnalyticsAppHealthOSVersionPerformance                       *userexperienceanalyticsapphealthosversionperformance.UserExperienceAnalyticsAppHealthOSVersionPerformanceClient
	UserExperienceAnalyticsAppHealthOverview                                   *userexperienceanalyticsapphealthoverview.UserExperienceAnalyticsAppHealthOverviewClient
	UserExperienceAnalyticsAppHealthOverviewMetricValue                        *userexperienceanalyticsapphealthoverviewmetricvalue.UserExperienceAnalyticsAppHealthOverviewMetricValueClient
	UserExperienceAnalyticsBaseline                                            *userexperienceanalyticsbaseline.UserExperienceAnalyticsBaselineClient
	UserExperienceAnalyticsBaselineAppHealthMetric                             *userexperienceanalyticsbaselineapphealthmetric.UserExperienceAnalyticsBaselineAppHealthMetricClient
	UserExperienceAnalyticsBaselineBatteryHealthMetric                         *userexperienceanalyticsbaselinebatteryhealthmetric.UserExperienceAnalyticsBaselineBatteryHealthMetricClient
	UserExperienceAnalyticsBaselineBestPracticesMetric                         *userexperienceanalyticsbaselinebestpracticesmetric.UserExperienceAnalyticsBaselineBestPracticesMetricClient
	UserExperienceAnalyticsBaselineDeviceBootPerformanceMetric                 *userexperienceanalyticsbaselinedevicebootperformancemetric.UserExperienceAnalyticsBaselineDeviceBootPerformanceMetricClient
	UserExperienceAnalyticsBaselineRebootAnalyticsMetric                       *userexperienceanalyticsbaselinerebootanalyticsmetric.UserExperienceAnalyticsBaselineRebootAnalyticsMetricClient
	UserExperienceAnalyticsBaselineResourcePerformanceMetric                   *userexperienceanalyticsbaselineresourceperformancemetric.UserExperienceAnalyticsBaselineResourcePerformanceMetricClient
	UserExperienceAnalyticsBaselineWorkFromAnywhereMetric                      *userexperienceanalyticsbaselineworkfromanywheremetric.UserExperienceAnalyticsBaselineWorkFromAnywhereMetricClient
	UserExperienceAnalyticsCategory                                            *userexperienceanalyticscategory.UserExperienceAnalyticsCategoryClient
	UserExperienceAnalyticsCategoryMetricValue                                 *userexperienceanalyticscategorymetricvalue.UserExperienceAnalyticsCategoryMetricValueClient
	UserExperienceAnalyticsDevicePerformance                                   *userexperienceanalyticsdeviceperformance.UserExperienceAnalyticsDevicePerformanceClient
	UserExperienceAnalyticsDeviceScore                                         *userexperienceanalyticsdevicescore.UserExperienceAnalyticsDeviceScoreClient
	UserExperienceAnalyticsDeviceStartupHistory                                *userexperienceanalyticsdevicestartuphistory.UserExperienceAnalyticsDeviceStartupHistoryClient
	UserExperienceAnalyticsDeviceStartupProcess                                *userexperienceanalyticsdevicestartupprocess.UserExperienceAnalyticsDeviceStartupProcessClient
	UserExperienceAnalyticsDeviceStartupProcessPerformance                     *userexperienceanalyticsdevicestartupprocessperformance.UserExperienceAnalyticsDeviceStartupProcessPerformanceClient
	UserExperienceAnalyticsMetricHistory                                       *userexperienceanalyticsmetrichistory.UserExperienceAnalyticsMetricHistoryClient
	UserExperienceAnalyticsModelScore                                          *userexperienceanalyticsmodelscore.UserExperienceAnalyticsModelScoreClient
	UserExperienceAnalyticsOverview                                            *userexperienceanalyticsoverview.UserExperienceAnalyticsOverviewClient
	UserExperienceAnalyticsScoreHistory                                        *userexperienceanalyticsscorehistory.UserExperienceAnalyticsScoreHistoryClient
	UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric             *userexperienceanalyticsworkfromanywherehardwarereadinessmetric.UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricClient
	UserExperienceAnalyticsWorkFromAnywhereMetric                              *userexperienceanalyticsworkfromanywheremetric.UserExperienceAnalyticsWorkFromAnywhereMetricClient
	UserExperienceAnalyticsWorkFromAnywhereMetricMetricDevice                  *userexperienceanalyticsworkfromanywheremetricmetricdevice.UserExperienceAnalyticsWorkFromAnywhereMetricMetricDeviceClient
	UserExperienceAnalyticsWorkFromAnywhereModelPerformance                    *userexperienceanalyticsworkfromanywheremodelperformance.UserExperienceAnalyticsWorkFromAnywhereModelPerformanceClient
	VirtualEndpoint                                                            *virtualendpoint.VirtualEndpointClient
	WindowsAutopilotDeviceIdentity                                             *windowsautopilotdeviceidentity.WindowsAutopilotDeviceIdentityClient
	WindowsInformationProtectionAppLearningSummary                             *windowsinformationprotectionapplearningsummary.WindowsInformationProtectionAppLearningSummaryClient
	WindowsInformationProtectionNetworkLearningSummary                         *windowsinformationprotectionnetworklearningsummary.WindowsInformationProtectionNetworkLearningSummaryClient
	WindowsMalwareInformation                                                  *windowsmalwareinformation.WindowsMalwareInformationClient
	WindowsMalwareInformationDeviceMalwareState                                *windowsmalwareinformationdevicemalwarestate.WindowsMalwareInformationDeviceMalwareStateClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *msgraph.Client)) (*Client, error) {
	applePushNotificationCertificateClient, err := applepushnotificationcertificate.NewApplePushNotificationCertificateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ApplePushNotificationCertificate client: %+v", err)
	}
	configureFunc(applePushNotificationCertificateClient.Client)

	auditEventClient, err := auditevent.NewAuditEventClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AuditEvent client: %+v", err)
	}
	configureFunc(auditEventClient.Client)

	complianceManagementPartnerClient, err := compliancemanagementpartner.NewComplianceManagementPartnerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ComplianceManagementPartner client: %+v", err)
	}
	configureFunc(complianceManagementPartnerClient.Client)

	conditionalAccessSettingClient, err := conditionalaccesssetting.NewConditionalAccessSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConditionalAccessSetting client: %+v", err)
	}
	configureFunc(conditionalAccessSettingClient.Client)

	detectedAppClient, err := detectedapp.NewDetectedAppClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DetectedApp client: %+v", err)
	}
	configureFunc(detectedAppClient.Client)

	detectedAppManagedDeviceClient, err := detectedappmanageddevice.NewDetectedAppManagedDeviceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DetectedAppManagedDevice client: %+v", err)
	}
	configureFunc(detectedAppManagedDeviceClient.Client)

	deviceCategoryClient, err := devicecategory.NewDeviceCategoryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceCategory client: %+v", err)
	}
	configureFunc(deviceCategoryClient.Client)

	deviceCompliancePolicyAssignmentClient, err := devicecompliancepolicyassignment.NewDeviceCompliancePolicyAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceCompliancePolicyAssignment client: %+v", err)
	}
	configureFunc(deviceCompliancePolicyAssignmentClient.Client)

	deviceCompliancePolicyClient, err := devicecompliancepolicy.NewDeviceCompliancePolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceCompliancePolicy client: %+v", err)
	}
	configureFunc(deviceCompliancePolicyClient.Client)

	deviceCompliancePolicyDeviceSettingStateSummaryClient, err := devicecompliancepolicydevicesettingstatesummary.NewDeviceCompliancePolicyDeviceSettingStateSummaryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceCompliancePolicyDeviceSettingStateSummary client: %+v", err)
	}
	configureFunc(deviceCompliancePolicyDeviceSettingStateSummaryClient.Client)

	deviceCompliancePolicyDeviceStateSummaryClient, err := devicecompliancepolicydevicestatesummary.NewDeviceCompliancePolicyDeviceStateSummaryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceCompliancePolicyDeviceStateSummary client: %+v", err)
	}
	configureFunc(deviceCompliancePolicyDeviceStateSummaryClient.Client)

	deviceCompliancePolicyDeviceStatusClient, err := devicecompliancepolicydevicestatus.NewDeviceCompliancePolicyDeviceStatusClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceCompliancePolicyDeviceStatus client: %+v", err)
	}
	configureFunc(deviceCompliancePolicyDeviceStatusClient.Client)

	deviceCompliancePolicyDeviceStatusOverviewClient, err := devicecompliancepolicydevicestatusoverview.NewDeviceCompliancePolicyDeviceStatusOverviewClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceCompliancePolicyDeviceStatusOverview client: %+v", err)
	}
	configureFunc(deviceCompliancePolicyDeviceStatusOverviewClient.Client)

	deviceCompliancePolicyScheduledActionsForRuleClient, err := devicecompliancepolicyscheduledactionsforrule.NewDeviceCompliancePolicyScheduledActionsForRuleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceCompliancePolicyScheduledActionsForRule client: %+v", err)
	}
	configureFunc(deviceCompliancePolicyScheduledActionsForRuleClient.Client)

	deviceCompliancePolicyScheduledActionsForRuleScheduledActionConfigurationClient, err := devicecompliancepolicyscheduledactionsforrulescheduledactionconfiguration.NewDeviceCompliancePolicyScheduledActionsForRuleScheduledActionConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceCompliancePolicyScheduledActionsForRuleScheduledActionConfiguration client: %+v", err)
	}
	configureFunc(deviceCompliancePolicyScheduledActionsForRuleScheduledActionConfigurationClient.Client)

	deviceCompliancePolicySettingStateSummaryClient, err := devicecompliancepolicysettingstatesummary.NewDeviceCompliancePolicySettingStateSummaryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceCompliancePolicySettingStateSummary client: %+v", err)
	}
	configureFunc(deviceCompliancePolicySettingStateSummaryClient.Client)

	deviceCompliancePolicySettingStateSummaryDeviceComplianceSettingStateClient, err := devicecompliancepolicysettingstatesummarydevicecompliancesettingstate.NewDeviceCompliancePolicySettingStateSummaryDeviceComplianceSettingStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceCompliancePolicySettingStateSummaryDeviceComplianceSettingState client: %+v", err)
	}
	configureFunc(deviceCompliancePolicySettingStateSummaryDeviceComplianceSettingStateClient.Client)

	deviceCompliancePolicyUserStatusClient, err := devicecompliancepolicyuserstatus.NewDeviceCompliancePolicyUserStatusClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceCompliancePolicyUserStatus client: %+v", err)
	}
	configureFunc(deviceCompliancePolicyUserStatusClient.Client)

	deviceCompliancePolicyUserStatusOverviewClient, err := devicecompliancepolicyuserstatusoverview.NewDeviceCompliancePolicyUserStatusOverviewClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceCompliancePolicyUserStatusOverview client: %+v", err)
	}
	configureFunc(deviceCompliancePolicyUserStatusOverviewClient.Client)

	deviceConfigurationAssignmentClient, err := deviceconfigurationassignment.NewDeviceConfigurationAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceConfigurationAssignment client: %+v", err)
	}
	configureFunc(deviceConfigurationAssignmentClient.Client)

	deviceConfigurationClient, err := deviceconfiguration.NewDeviceConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceConfiguration client: %+v", err)
	}
	configureFunc(deviceConfigurationClient.Client)

	deviceConfigurationDeviceSettingStateSummaryClient, err := deviceconfigurationdevicesettingstatesummary.NewDeviceConfigurationDeviceSettingStateSummaryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceConfigurationDeviceSettingStateSummary client: %+v", err)
	}
	configureFunc(deviceConfigurationDeviceSettingStateSummaryClient.Client)

	deviceConfigurationDeviceStateSummaryClient, err := deviceconfigurationdevicestatesummary.NewDeviceConfigurationDeviceStateSummaryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceConfigurationDeviceStateSummary client: %+v", err)
	}
	configureFunc(deviceConfigurationDeviceStateSummaryClient.Client)

	deviceConfigurationDeviceStatusClient, err := deviceconfigurationdevicestatus.NewDeviceConfigurationDeviceStatusClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceConfigurationDeviceStatus client: %+v", err)
	}
	configureFunc(deviceConfigurationDeviceStatusClient.Client)

	deviceConfigurationDeviceStatusOverviewClient, err := deviceconfigurationdevicestatusoverview.NewDeviceConfigurationDeviceStatusOverviewClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceConfigurationDeviceStatusOverview client: %+v", err)
	}
	configureFunc(deviceConfigurationDeviceStatusOverviewClient.Client)

	deviceConfigurationUserStatusClient, err := deviceconfigurationuserstatus.NewDeviceConfigurationUserStatusClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceConfigurationUserStatus client: %+v", err)
	}
	configureFunc(deviceConfigurationUserStatusClient.Client)

	deviceConfigurationUserStatusOverviewClient, err := deviceconfigurationuserstatusoverview.NewDeviceConfigurationUserStatusOverviewClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceConfigurationUserStatusOverview client: %+v", err)
	}
	configureFunc(deviceConfigurationUserStatusOverviewClient.Client)

	deviceEnrollmentConfigurationAssignmentClient, err := deviceenrollmentconfigurationassignment.NewDeviceEnrollmentConfigurationAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceEnrollmentConfigurationAssignment client: %+v", err)
	}
	configureFunc(deviceEnrollmentConfigurationAssignmentClient.Client)

	deviceEnrollmentConfigurationClient, err := deviceenrollmentconfiguration.NewDeviceEnrollmentConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceEnrollmentConfiguration client: %+v", err)
	}
	configureFunc(deviceEnrollmentConfigurationClient.Client)

	deviceManagementClient, err := devicemanagement.NewDeviceManagementClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceManagement client: %+v", err)
	}
	configureFunc(deviceManagementClient.Client)

	deviceManagementPartnerClient, err := devicemanagementpartner.NewDeviceManagementPartnerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceManagementPartner client: %+v", err)
	}
	configureFunc(deviceManagementPartnerClient.Client)

	exchangeConnectorClient, err := exchangeconnector.NewExchangeConnectorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExchangeConnector client: %+v", err)
	}
	configureFunc(exchangeConnectorClient.Client)

	importedWindowsAutopilotDeviceIdentityClient, err := importedwindowsautopilotdeviceidentity.NewImportedWindowsAutopilotDeviceIdentityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ImportedWindowsAutopilotDeviceIdentity client: %+v", err)
	}
	configureFunc(importedWindowsAutopilotDeviceIdentityClient.Client)

	iosUpdateStatusClient, err := iosupdatestatus.NewIosUpdateStatusClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IosUpdateStatus client: %+v", err)
	}
	configureFunc(iosUpdateStatusClient.Client)

	managedDeviceClient, err := manageddevice.NewManagedDeviceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDevice client: %+v", err)
	}
	configureFunc(managedDeviceClient.Client)

	managedDeviceDeviceCategoryClient, err := manageddevicedevicecategory.NewManagedDeviceDeviceCategoryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDeviceDeviceCategory client: %+v", err)
	}
	configureFunc(managedDeviceDeviceCategoryClient.Client)

	managedDeviceDeviceCompliancePolicyStateClient, err := manageddevicedevicecompliancepolicystate.NewManagedDeviceDeviceCompliancePolicyStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDeviceDeviceCompliancePolicyState client: %+v", err)
	}
	configureFunc(managedDeviceDeviceCompliancePolicyStateClient.Client)

	managedDeviceDeviceConfigurationStateClient, err := manageddevicedeviceconfigurationstate.NewManagedDeviceDeviceConfigurationStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDeviceDeviceConfigurationState client: %+v", err)
	}
	configureFunc(managedDeviceDeviceConfigurationStateClient.Client)

	managedDeviceLogCollectionRequestClient, err := manageddevicelogcollectionrequest.NewManagedDeviceLogCollectionRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDeviceLogCollectionRequest client: %+v", err)
	}
	configureFunc(managedDeviceLogCollectionRequestClient.Client)

	managedDeviceOverviewClient, err := manageddeviceoverview.NewManagedDeviceOverviewClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDeviceOverview client: %+v", err)
	}
	configureFunc(managedDeviceOverviewClient.Client)

	managedDeviceUserClient, err := manageddeviceuser.NewManagedDeviceUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDeviceUser client: %+v", err)
	}
	configureFunc(managedDeviceUserClient.Client)

	managedDeviceWindowsProtectionStateClient, err := manageddevicewindowsprotectionstate.NewManagedDeviceWindowsProtectionStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDeviceWindowsProtectionState client: %+v", err)
	}
	configureFunc(managedDeviceWindowsProtectionStateClient.Client)

	managedDeviceWindowsProtectionStateDetectedMalwareStateClient, err := manageddevicewindowsprotectionstatedetectedmalwarestate.NewManagedDeviceWindowsProtectionStateDetectedMalwareStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDeviceWindowsProtectionStateDetectedMalwareState client: %+v", err)
	}
	configureFunc(managedDeviceWindowsProtectionStateDetectedMalwareStateClient.Client)

	mobileAppTroubleshootingEventAppLogCollectionRequestClient, err := mobileapptroubleshootingeventapplogcollectionrequest.NewMobileAppTroubleshootingEventAppLogCollectionRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MobileAppTroubleshootingEventAppLogCollectionRequest client: %+v", err)
	}
	configureFunc(mobileAppTroubleshootingEventAppLogCollectionRequestClient.Client)

	mobileAppTroubleshootingEventClient, err := mobileapptroubleshootingevent.NewMobileAppTroubleshootingEventClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MobileAppTroubleshootingEvent client: %+v", err)
	}
	configureFunc(mobileAppTroubleshootingEventClient.Client)

	mobileThreatDefenseConnectorClient, err := mobilethreatdefenseconnector.NewMobileThreatDefenseConnectorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MobileThreatDefenseConnector client: %+v", err)
	}
	configureFunc(mobileThreatDefenseConnectorClient.Client)

	notificationMessageTemplateClient, err := notificationmessagetemplate.NewNotificationMessageTemplateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NotificationMessageTemplate client: %+v", err)
	}
	configureFunc(notificationMessageTemplateClient.Client)

	notificationMessageTemplateLocalizedNotificationMessageClient, err := notificationmessagetemplatelocalizednotificationmessage.NewNotificationMessageTemplateLocalizedNotificationMessageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NotificationMessageTemplateLocalizedNotificationMessage client: %+v", err)
	}
	configureFunc(notificationMessageTemplateLocalizedNotificationMessageClient.Client)

	remoteAssistancePartnerClient, err := remoteassistancepartner.NewRemoteAssistancePartnerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RemoteAssistancePartner client: %+v", err)
	}
	configureFunc(remoteAssistancePartnerClient.Client)

	reportClient, err := report.NewReportClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Report client: %+v", err)
	}
	configureFunc(reportClient.Client)

	reportExportJobClient, err := reportexportjob.NewReportExportJobClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ReportExportJob client: %+v", err)
	}
	configureFunc(reportExportJobClient.Client)

	resourceOperationClient, err := resourceoperation.NewResourceOperationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ResourceOperation client: %+v", err)
	}
	configureFunc(resourceOperationClient.Client)

	roleAssignmentClient, err := roleassignment.NewRoleAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleAssignment client: %+v", err)
	}
	configureFunc(roleAssignmentClient.Client)

	roleAssignmentRoleDefinitionClient, err := roleassignmentroledefinition.NewRoleAssignmentRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleAssignmentRoleDefinition client: %+v", err)
	}
	configureFunc(roleAssignmentRoleDefinitionClient.Client)

	roleDefinitionClient, err := roledefinition.NewRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleDefinition client: %+v", err)
	}
	configureFunc(roleDefinitionClient.Client)

	roleDefinitionRoleAssignmentClient, err := roledefinitionroleassignment.NewRoleDefinitionRoleAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleDefinitionRoleAssignment client: %+v", err)
	}
	configureFunc(roleDefinitionRoleAssignmentClient.Client)

	roleDefinitionRoleAssignmentRoleDefinitionClient, err := roledefinitionroleassignmentroledefinition.NewRoleDefinitionRoleAssignmentRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleDefinitionRoleAssignmentRoleDefinition client: %+v", err)
	}
	configureFunc(roleDefinitionRoleAssignmentRoleDefinitionClient.Client)

	softwareUpdateStatusSummaryClient, err := softwareupdatestatussummary.NewSoftwareUpdateStatusSummaryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SoftwareUpdateStatusSummary client: %+v", err)
	}
	configureFunc(softwareUpdateStatusSummaryClient.Client)

	telecomExpenseManagementPartnerClient, err := telecomexpensemanagementpartner.NewTelecomExpenseManagementPartnerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TelecomExpenseManagementPartner client: %+v", err)
	}
	configureFunc(telecomExpenseManagementPartnerClient.Client)

	termsAndConditionAcceptanceStatusClient, err := termsandconditionacceptancestatus.NewTermsAndConditionAcceptanceStatusClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TermsAndConditionAcceptanceStatus client: %+v", err)
	}
	configureFunc(termsAndConditionAcceptanceStatusClient.Client)

	termsAndConditionAcceptanceStatusTermsAndConditionClient, err := termsandconditionacceptancestatustermsandcondition.NewTermsAndConditionAcceptanceStatusTermsAndConditionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TermsAndConditionAcceptanceStatusTermsAndCondition client: %+v", err)
	}
	configureFunc(termsAndConditionAcceptanceStatusTermsAndConditionClient.Client)

	termsAndConditionAssignmentClient, err := termsandconditionassignment.NewTermsAndConditionAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TermsAndConditionAssignment client: %+v", err)
	}
	configureFunc(termsAndConditionAssignmentClient.Client)

	termsAndConditionClient, err := termsandcondition.NewTermsAndConditionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TermsAndCondition client: %+v", err)
	}
	configureFunc(termsAndConditionClient.Client)

	troubleshootingEventClient, err := troubleshootingevent.NewTroubleshootingEventClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TroubleshootingEvent client: %+v", err)
	}
	configureFunc(troubleshootingEventClient.Client)

	userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailClient, err := userexperienceanalyticsapphealthapplicationperformancebyappversiondetail.NewUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetail client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailClient.Client)

	userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdClient, err := userexperienceanalyticsapphealthapplicationperformancebyappversiondeviceid.NewUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdClient.Client)

	userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionClient, err := userexperienceanalyticsapphealthapplicationperformancebyosversion.NewUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionClient.Client)

	userExperienceAnalyticsAppHealthApplicationPerformanceClient, err := userexperienceanalyticsapphealthapplicationperformance.NewUserExperienceAnalyticsAppHealthApplicationPerformanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsAppHealthApplicationPerformance client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsAppHealthApplicationPerformanceClient.Client)

	userExperienceAnalyticsAppHealthDeviceModelPerformanceClient, err := userexperienceanalyticsapphealthdevicemodelperformance.NewUserExperienceAnalyticsAppHealthDeviceModelPerformanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsAppHealthDeviceModelPerformance client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsAppHealthDeviceModelPerformanceClient.Client)

	userExperienceAnalyticsAppHealthDevicePerformanceClient, err := userexperienceanalyticsapphealthdeviceperformance.NewUserExperienceAnalyticsAppHealthDevicePerformanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsAppHealthDevicePerformance client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsAppHealthDevicePerformanceClient.Client)

	userExperienceAnalyticsAppHealthDevicePerformanceDetailClient, err := userexperienceanalyticsapphealthdeviceperformancedetail.NewUserExperienceAnalyticsAppHealthDevicePerformanceDetailClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsAppHealthDevicePerformanceDetail client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsAppHealthDevicePerformanceDetailClient.Client)

	userExperienceAnalyticsAppHealthOSVersionPerformanceClient, err := userexperienceanalyticsapphealthosversionperformance.NewUserExperienceAnalyticsAppHealthOSVersionPerformanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsAppHealthOSVersionPerformance client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsAppHealthOSVersionPerformanceClient.Client)

	userExperienceAnalyticsAppHealthOverviewClient, err := userexperienceanalyticsapphealthoverview.NewUserExperienceAnalyticsAppHealthOverviewClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsAppHealthOverview client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsAppHealthOverviewClient.Client)

	userExperienceAnalyticsAppHealthOverviewMetricValueClient, err := userexperienceanalyticsapphealthoverviewmetricvalue.NewUserExperienceAnalyticsAppHealthOverviewMetricValueClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsAppHealthOverviewMetricValue client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsAppHealthOverviewMetricValueClient.Client)

	userExperienceAnalyticsBaselineAppHealthMetricClient, err := userexperienceanalyticsbaselineapphealthmetric.NewUserExperienceAnalyticsBaselineAppHealthMetricClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsBaselineAppHealthMetric client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsBaselineAppHealthMetricClient.Client)

	userExperienceAnalyticsBaselineBatteryHealthMetricClient, err := userexperienceanalyticsbaselinebatteryhealthmetric.NewUserExperienceAnalyticsBaselineBatteryHealthMetricClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsBaselineBatteryHealthMetric client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsBaselineBatteryHealthMetricClient.Client)

	userExperienceAnalyticsBaselineBestPracticesMetricClient, err := userexperienceanalyticsbaselinebestpracticesmetric.NewUserExperienceAnalyticsBaselineBestPracticesMetricClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsBaselineBestPracticesMetric client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsBaselineBestPracticesMetricClient.Client)

	userExperienceAnalyticsBaselineClient, err := userexperienceanalyticsbaseline.NewUserExperienceAnalyticsBaselineClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsBaseline client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsBaselineClient.Client)

	userExperienceAnalyticsBaselineDeviceBootPerformanceMetricClient, err := userexperienceanalyticsbaselinedevicebootperformancemetric.NewUserExperienceAnalyticsBaselineDeviceBootPerformanceMetricClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsBaselineDeviceBootPerformanceMetric client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsBaselineDeviceBootPerformanceMetricClient.Client)

	userExperienceAnalyticsBaselineRebootAnalyticsMetricClient, err := userexperienceanalyticsbaselinerebootanalyticsmetric.NewUserExperienceAnalyticsBaselineRebootAnalyticsMetricClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsBaselineRebootAnalyticsMetric client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsBaselineRebootAnalyticsMetricClient.Client)

	userExperienceAnalyticsBaselineResourcePerformanceMetricClient, err := userexperienceanalyticsbaselineresourceperformancemetric.NewUserExperienceAnalyticsBaselineResourcePerformanceMetricClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsBaselineResourcePerformanceMetric client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsBaselineResourcePerformanceMetricClient.Client)

	userExperienceAnalyticsBaselineWorkFromAnywhereMetricClient, err := userexperienceanalyticsbaselineworkfromanywheremetric.NewUserExperienceAnalyticsBaselineWorkFromAnywhereMetricClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsBaselineWorkFromAnywhereMetric client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsBaselineWorkFromAnywhereMetricClient.Client)

	userExperienceAnalyticsCategoryClient, err := userexperienceanalyticscategory.NewUserExperienceAnalyticsCategoryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsCategory client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsCategoryClient.Client)

	userExperienceAnalyticsCategoryMetricValueClient, err := userexperienceanalyticscategorymetricvalue.NewUserExperienceAnalyticsCategoryMetricValueClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsCategoryMetricValue client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsCategoryMetricValueClient.Client)

	userExperienceAnalyticsDevicePerformanceClient, err := userexperienceanalyticsdeviceperformance.NewUserExperienceAnalyticsDevicePerformanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsDevicePerformance client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsDevicePerformanceClient.Client)

	userExperienceAnalyticsDeviceScoreClient, err := userexperienceanalyticsdevicescore.NewUserExperienceAnalyticsDeviceScoreClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsDeviceScore client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsDeviceScoreClient.Client)

	userExperienceAnalyticsDeviceStartupHistoryClient, err := userexperienceanalyticsdevicestartuphistory.NewUserExperienceAnalyticsDeviceStartupHistoryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsDeviceStartupHistory client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsDeviceStartupHistoryClient.Client)

	userExperienceAnalyticsDeviceStartupProcessClient, err := userexperienceanalyticsdevicestartupprocess.NewUserExperienceAnalyticsDeviceStartupProcessClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsDeviceStartupProcess client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsDeviceStartupProcessClient.Client)

	userExperienceAnalyticsDeviceStartupProcessPerformanceClient, err := userexperienceanalyticsdevicestartupprocessperformance.NewUserExperienceAnalyticsDeviceStartupProcessPerformanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsDeviceStartupProcessPerformance client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsDeviceStartupProcessPerformanceClient.Client)

	userExperienceAnalyticsMetricHistoryClient, err := userexperienceanalyticsmetrichistory.NewUserExperienceAnalyticsMetricHistoryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsMetricHistory client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsMetricHistoryClient.Client)

	userExperienceAnalyticsModelScoreClient, err := userexperienceanalyticsmodelscore.NewUserExperienceAnalyticsModelScoreClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsModelScore client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsModelScoreClient.Client)

	userExperienceAnalyticsOverviewClient, err := userexperienceanalyticsoverview.NewUserExperienceAnalyticsOverviewClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsOverview client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsOverviewClient.Client)

	userExperienceAnalyticsScoreHistoryClient, err := userexperienceanalyticsscorehistory.NewUserExperienceAnalyticsScoreHistoryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsScoreHistory client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsScoreHistoryClient.Client)

	userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricClient, err := userexperienceanalyticsworkfromanywherehardwarereadinessmetric.NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricClient.Client)

	userExperienceAnalyticsWorkFromAnywhereMetricClient, err := userexperienceanalyticsworkfromanywheremetric.NewUserExperienceAnalyticsWorkFromAnywhereMetricClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsWorkFromAnywhereMetric client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsWorkFromAnywhereMetricClient.Client)

	userExperienceAnalyticsWorkFromAnywhereMetricMetricDeviceClient, err := userexperienceanalyticsworkfromanywheremetricmetricdevice.NewUserExperienceAnalyticsWorkFromAnywhereMetricMetricDeviceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsWorkFromAnywhereMetricMetricDevice client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsWorkFromAnywhereMetricMetricDeviceClient.Client)

	userExperienceAnalyticsWorkFromAnywhereModelPerformanceClient, err := userexperienceanalyticsworkfromanywheremodelperformance.NewUserExperienceAnalyticsWorkFromAnywhereModelPerformanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExperienceAnalyticsWorkFromAnywhereModelPerformance client: %+v", err)
	}
	configureFunc(userExperienceAnalyticsWorkFromAnywhereModelPerformanceClient.Client)

	virtualEndpointClient, err := virtualendpoint.NewVirtualEndpointClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualEndpoint client: %+v", err)
	}
	configureFunc(virtualEndpointClient.Client)

	windowsAutopilotDeviceIdentityClient, err := windowsautopilotdeviceidentity.NewWindowsAutopilotDeviceIdentityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WindowsAutopilotDeviceIdentity client: %+v", err)
	}
	configureFunc(windowsAutopilotDeviceIdentityClient.Client)

	windowsInformationProtectionAppLearningSummaryClient, err := windowsinformationprotectionapplearningsummary.NewWindowsInformationProtectionAppLearningSummaryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WindowsInformationProtectionAppLearningSummary client: %+v", err)
	}
	configureFunc(windowsInformationProtectionAppLearningSummaryClient.Client)

	windowsInformationProtectionNetworkLearningSummaryClient, err := windowsinformationprotectionnetworklearningsummary.NewWindowsInformationProtectionNetworkLearningSummaryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WindowsInformationProtectionNetworkLearningSummary client: %+v", err)
	}
	configureFunc(windowsInformationProtectionNetworkLearningSummaryClient.Client)

	windowsMalwareInformationClient, err := windowsmalwareinformation.NewWindowsMalwareInformationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WindowsMalwareInformation client: %+v", err)
	}
	configureFunc(windowsMalwareInformationClient.Client)

	windowsMalwareInformationDeviceMalwareStateClient, err := windowsmalwareinformationdevicemalwarestate.NewWindowsMalwareInformationDeviceMalwareStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WindowsMalwareInformationDeviceMalwareState client: %+v", err)
	}
	configureFunc(windowsMalwareInformationDeviceMalwareStateClient.Client)

	return &Client{
		ApplePushNotificationCertificate:                applePushNotificationCertificateClient,
		AuditEvent:                                      auditEventClient,
		ComplianceManagementPartner:                     complianceManagementPartnerClient,
		ConditionalAccessSetting:                        conditionalAccessSettingClient,
		DetectedApp:                                     detectedAppClient,
		DetectedAppManagedDevice:                        detectedAppManagedDeviceClient,
		DeviceCategory:                                  deviceCategoryClient,
		DeviceCompliancePolicy:                          deviceCompliancePolicyClient,
		DeviceCompliancePolicyAssignment:                deviceCompliancePolicyAssignmentClient,
		DeviceCompliancePolicyDeviceSettingStateSummary: deviceCompliancePolicyDeviceSettingStateSummaryClient,
		DeviceCompliancePolicyDeviceStateSummary:        deviceCompliancePolicyDeviceStateSummaryClient,
		DeviceCompliancePolicyDeviceStatus:              deviceCompliancePolicyDeviceStatusClient,
		DeviceCompliancePolicyDeviceStatusOverview:      deviceCompliancePolicyDeviceStatusOverviewClient,
		DeviceCompliancePolicyScheduledActionsForRule:   deviceCompliancePolicyScheduledActionsForRuleClient,
		DeviceCompliancePolicyScheduledActionsForRuleScheduledActionConfiguration: deviceCompliancePolicyScheduledActionsForRuleScheduledActionConfigurationClient,
		DeviceCompliancePolicySettingStateSummary:                                 deviceCompliancePolicySettingStateSummaryClient,
		DeviceCompliancePolicySettingStateSummaryDeviceComplianceSettingState:     deviceCompliancePolicySettingStateSummaryDeviceComplianceSettingStateClient,
		DeviceCompliancePolicyUserStatus:                                          deviceCompliancePolicyUserStatusClient,
		DeviceCompliancePolicyUserStatusOverview:                                  deviceCompliancePolicyUserStatusOverviewClient,
		DeviceConfiguration:                                                       deviceConfigurationClient,
		DeviceConfigurationAssignment:                                             deviceConfigurationAssignmentClient,
		DeviceConfigurationDeviceSettingStateSummary:                              deviceConfigurationDeviceSettingStateSummaryClient,
		DeviceConfigurationDeviceStateSummary:                                     deviceConfigurationDeviceStateSummaryClient,
		DeviceConfigurationDeviceStatus:                                           deviceConfigurationDeviceStatusClient,
		DeviceConfigurationDeviceStatusOverview:                                   deviceConfigurationDeviceStatusOverviewClient,
		DeviceConfigurationUserStatus:                                             deviceConfigurationUserStatusClient,
		DeviceConfigurationUserStatusOverview:                                     deviceConfigurationUserStatusOverviewClient,
		DeviceEnrollmentConfiguration:                                             deviceEnrollmentConfigurationClient,
		DeviceEnrollmentConfigurationAssignment:                                   deviceEnrollmentConfigurationAssignmentClient,
		DeviceManagement:                                                          deviceManagementClient,
		DeviceManagementPartner:                                                   deviceManagementPartnerClient,
		ExchangeConnector:                                                         exchangeConnectorClient,
		ImportedWindowsAutopilotDeviceIdentity:                                    importedWindowsAutopilotDeviceIdentityClient,
		IosUpdateStatus:                                                           iosUpdateStatusClient,
		ManagedDevice:                                                             managedDeviceClient,
		ManagedDeviceDeviceCategory:                                               managedDeviceDeviceCategoryClient,
		ManagedDeviceDeviceCompliancePolicyState:                                  managedDeviceDeviceCompliancePolicyStateClient,
		ManagedDeviceDeviceConfigurationState:                                     managedDeviceDeviceConfigurationStateClient,
		ManagedDeviceLogCollectionRequest:                                         managedDeviceLogCollectionRequestClient,
		ManagedDeviceOverview:                                                     managedDeviceOverviewClient,
		ManagedDeviceUser:                                                         managedDeviceUserClient,
		ManagedDeviceWindowsProtectionState:                                       managedDeviceWindowsProtectionStateClient,
		ManagedDeviceWindowsProtectionStateDetectedMalwareState:                   managedDeviceWindowsProtectionStateDetectedMalwareStateClient,
		MobileAppTroubleshootingEvent:                                             mobileAppTroubleshootingEventClient,
		MobileAppTroubleshootingEventAppLogCollectionRequest:                      mobileAppTroubleshootingEventAppLogCollectionRequestClient,
		MobileThreatDefenseConnector:                                              mobileThreatDefenseConnectorClient,
		NotificationMessageTemplate:                                               notificationMessageTemplateClient,
		NotificationMessageTemplateLocalizedNotificationMessage:                   notificationMessageTemplateLocalizedNotificationMessageClient,
		RemoteAssistancePartner:                                                   remoteAssistancePartnerClient,
		Report:                                                                    reportClient,
		ReportExportJob:                                                           reportExportJobClient,
		ResourceOperation:                                                         resourceOperationClient,
		RoleAssignment:                                                            roleAssignmentClient,
		RoleAssignmentRoleDefinition:                                              roleAssignmentRoleDefinitionClient,
		RoleDefinition:                                                            roleDefinitionClient,
		RoleDefinitionRoleAssignment:                                              roleDefinitionRoleAssignmentClient,
		RoleDefinitionRoleAssignmentRoleDefinition:                                roleDefinitionRoleAssignmentRoleDefinitionClient,
		SoftwareUpdateStatusSummary:                                               softwareUpdateStatusSummaryClient,
		TelecomExpenseManagementPartner:                                           telecomExpenseManagementPartnerClient,
		TermsAndCondition:                                                         termsAndConditionClient,
		TermsAndConditionAcceptanceStatus:                                         termsAndConditionAcceptanceStatusClient,
		TermsAndConditionAcceptanceStatusTermsAndCondition:                        termsAndConditionAcceptanceStatusTermsAndConditionClient,
		TermsAndConditionAssignment:                                               termsAndConditionAssignmentClient,
		TroubleshootingEvent:                                                      troubleshootingEventClient,
		UserExperienceAnalyticsAppHealthApplicationPerformance:                    userExperienceAnalyticsAppHealthApplicationPerformanceClient,
		UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetail:   userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailClient,
		UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId: userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdClient,
		UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion:          userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionClient,
		UserExperienceAnalyticsAppHealthDeviceModelPerformance:                     userExperienceAnalyticsAppHealthDeviceModelPerformanceClient,
		UserExperienceAnalyticsAppHealthDevicePerformance:                          userExperienceAnalyticsAppHealthDevicePerformanceClient,
		UserExperienceAnalyticsAppHealthDevicePerformanceDetail:                    userExperienceAnalyticsAppHealthDevicePerformanceDetailClient,
		UserExperienceAnalyticsAppHealthOSVersionPerformance:                       userExperienceAnalyticsAppHealthOSVersionPerformanceClient,
		UserExperienceAnalyticsAppHealthOverview:                                   userExperienceAnalyticsAppHealthOverviewClient,
		UserExperienceAnalyticsAppHealthOverviewMetricValue:                        userExperienceAnalyticsAppHealthOverviewMetricValueClient,
		UserExperienceAnalyticsBaseline:                                            userExperienceAnalyticsBaselineClient,
		UserExperienceAnalyticsBaselineAppHealthMetric:                             userExperienceAnalyticsBaselineAppHealthMetricClient,
		UserExperienceAnalyticsBaselineBatteryHealthMetric:                         userExperienceAnalyticsBaselineBatteryHealthMetricClient,
		UserExperienceAnalyticsBaselineBestPracticesMetric:                         userExperienceAnalyticsBaselineBestPracticesMetricClient,
		UserExperienceAnalyticsBaselineDeviceBootPerformanceMetric:                 userExperienceAnalyticsBaselineDeviceBootPerformanceMetricClient,
		UserExperienceAnalyticsBaselineRebootAnalyticsMetric:                       userExperienceAnalyticsBaselineRebootAnalyticsMetricClient,
		UserExperienceAnalyticsBaselineResourcePerformanceMetric:                   userExperienceAnalyticsBaselineResourcePerformanceMetricClient,
		UserExperienceAnalyticsBaselineWorkFromAnywhereMetric:                      userExperienceAnalyticsBaselineWorkFromAnywhereMetricClient,
		UserExperienceAnalyticsCategory:                                            userExperienceAnalyticsCategoryClient,
		UserExperienceAnalyticsCategoryMetricValue:                                 userExperienceAnalyticsCategoryMetricValueClient,
		UserExperienceAnalyticsDevicePerformance:                                   userExperienceAnalyticsDevicePerformanceClient,
		UserExperienceAnalyticsDeviceScore:                                         userExperienceAnalyticsDeviceScoreClient,
		UserExperienceAnalyticsDeviceStartupHistory:                                userExperienceAnalyticsDeviceStartupHistoryClient,
		UserExperienceAnalyticsDeviceStartupProcess:                                userExperienceAnalyticsDeviceStartupProcessClient,
		UserExperienceAnalyticsDeviceStartupProcessPerformance:                     userExperienceAnalyticsDeviceStartupProcessPerformanceClient,
		UserExperienceAnalyticsMetricHistory:                                       userExperienceAnalyticsMetricHistoryClient,
		UserExperienceAnalyticsModelScore:                                          userExperienceAnalyticsModelScoreClient,
		UserExperienceAnalyticsOverview:                                            userExperienceAnalyticsOverviewClient,
		UserExperienceAnalyticsScoreHistory:                                        userExperienceAnalyticsScoreHistoryClient,
		UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric:             userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricClient,
		UserExperienceAnalyticsWorkFromAnywhereMetric:                              userExperienceAnalyticsWorkFromAnywhereMetricClient,
		UserExperienceAnalyticsWorkFromAnywhereMetricMetricDevice:                  userExperienceAnalyticsWorkFromAnywhereMetricMetricDeviceClient,
		UserExperienceAnalyticsWorkFromAnywhereModelPerformance:                    userExperienceAnalyticsWorkFromAnywhereModelPerformanceClient,
		VirtualEndpoint:                                    virtualEndpointClient,
		WindowsAutopilotDeviceIdentity:                     windowsAutopilotDeviceIdentityClient,
		WindowsInformationProtectionAppLearningSummary:     windowsInformationProtectionAppLearningSummaryClient,
		WindowsInformationProtectionNetworkLearningSummary: windowsInformationProtectionNetworkLearningSummaryClient,
		WindowsMalwareInformation:                          windowsMalwareInformationClient,
		WindowsMalwareInformationDeviceMalwareState:        windowsMalwareInformationDeviceMalwareStateClient,
	}, nil
}
