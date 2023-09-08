package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagement struct {
	ApplePushNotificationCertificate                                           *ApplePushNotificationCertificate                                     `json:"applePushNotificationCertificate,omitempty"`
	AuditEvents                                                                *[]AuditEvent                                                         `json:"auditEvents,omitempty"`
	ComplianceManagementPartners                                               *[]ComplianceManagementPartner                                        `json:"complianceManagementPartners,omitempty"`
	ConditionalAccessSettings                                                  *OnPremisesConditionalAccessSettings                                  `json:"conditionalAccessSettings,omitempty"`
	DetectedApps                                                               *[]DetectedApp                                                        `json:"detectedApps,omitempty"`
	DeviceCategories                                                           *[]DeviceCategory                                                     `json:"deviceCategories,omitempty"`
	DeviceCompliancePolicies                                                   *[]DeviceCompliancePolicy                                             `json:"deviceCompliancePolicies,omitempty"`
	DeviceCompliancePolicyDeviceStateSummary                                   *DeviceCompliancePolicyDeviceStateSummary                             `json:"deviceCompliancePolicyDeviceStateSummary,omitempty"`
	DeviceCompliancePolicySettingStateSummaries                                *[]DeviceCompliancePolicySettingStateSummary                          `json:"deviceCompliancePolicySettingStateSummaries,omitempty"`
	DeviceConfigurationDeviceStateSummaries                                    *DeviceConfigurationDeviceStateSummary                                `json:"deviceConfigurationDeviceStateSummaries,omitempty"`
	DeviceConfigurations                                                       *[]DeviceConfiguration                                                `json:"deviceConfigurations,omitempty"`
	DeviceEnrollmentConfigurations                                             *[]DeviceEnrollmentConfiguration                                      `json:"deviceEnrollmentConfigurations,omitempty"`
	DeviceManagementPartners                                                   *[]DeviceManagementPartner                                            `json:"deviceManagementPartners,omitempty"`
	DeviceProtectionOverview                                                   *DeviceProtectionOverview                                             `json:"deviceProtectionOverview,omitempty"`
	ExchangeConnectors                                                         *[]DeviceManagementExchangeConnector                                  `json:"exchangeConnectors,omitempty"`
	Id                                                                         *string                                                               `json:"id,omitempty"`
	ImportedWindowsAutopilotDeviceIdentities                                   *[]ImportedWindowsAutopilotDeviceIdentity                             `json:"importedWindowsAutopilotDeviceIdentities,omitempty"`
	IntuneAccountId                                                            *string                                                               `json:"intuneAccountId,omitempty"`
	IntuneBrand                                                                *IntuneBrand                                                          `json:"intuneBrand,omitempty"`
	IosUpdateStatuses                                                          *[]IosUpdateDeviceStatus                                              `json:"iosUpdateStatuses,omitempty"`
	ManagedDeviceOverview                                                      *ManagedDeviceOverview                                                `json:"managedDeviceOverview,omitempty"`
	ManagedDevices                                                             *[]ManagedDevice                                                      `json:"managedDevices,omitempty"`
	MobileAppTroubleshootingEvents                                             *[]MobileAppTroubleshootingEvent                                      `json:"mobileAppTroubleshootingEvents,omitempty"`
	MobileThreatDefenseConnectors                                              *[]MobileThreatDefenseConnector                                       `json:"mobileThreatDefenseConnectors,omitempty"`
	NotificationMessageTemplates                                               *[]NotificationMessageTemplate                                        `json:"notificationMessageTemplates,omitempty"`
	ODataType                                                                  *string                                                               `json:"@odata.type,omitempty"`
	RemoteAssistancePartners                                                   *[]RemoteAssistancePartner                                            `json:"remoteAssistancePartners,omitempty"`
	Reports                                                                    *DeviceManagementReports                                              `json:"reports,omitempty"`
	ResourceOperations                                                         *[]ResourceOperation                                                  `json:"resourceOperations,omitempty"`
	RoleAssignments                                                            *[]DeviceAndAppManagementRoleAssignment                               `json:"roleAssignments,omitempty"`
	RoleDefinitions                                                            *[]RoleDefinition                                                     `json:"roleDefinitions,omitempty"`
	Settings                                                                   *DeviceManagementSettings                                             `json:"settings,omitempty"`
	SoftwareUpdateStatusSummary                                                *SoftwareUpdateStatusSummary                                          `json:"softwareUpdateStatusSummary,omitempty"`
	SubscriptionState                                                          *DeviceManagementSubscriptionState                                    `json:"subscriptionState,omitempty"`
	TelecomExpenseManagementPartners                                           *[]TelecomExpenseManagementPartner                                    `json:"telecomExpenseManagementPartners,omitempty"`
	TermsAndConditions                                                         *[]TermsAndConditions                                                 `json:"termsAndConditions,omitempty"`
	TroubleshootingEvents                                                      *[]DeviceManagementTroubleshootingEvent                               `json:"troubleshootingEvents,omitempty"`
	UserExperienceAnalyticsAppHealthApplicationPerformance                     *[]UserExperienceAnalyticsAppHealthApplicationPerformance             `json:"userExperienceAnalyticsAppHealthApplicationPerformance,omitempty"`
	UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails  *[]UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails  `json:"userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails,omitempty"`
	UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId *[]UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId `json:"userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId,omitempty"`
	UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion          *[]UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion          `json:"userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion,omitempty"`
	UserExperienceAnalyticsAppHealthDeviceModelPerformance                     *[]UserExperienceAnalyticsAppHealthDeviceModelPerformance             `json:"userExperienceAnalyticsAppHealthDeviceModelPerformance,omitempty"`
	UserExperienceAnalyticsAppHealthDevicePerformance                          *[]UserExperienceAnalyticsAppHealthDevicePerformance                  `json:"userExperienceAnalyticsAppHealthDevicePerformance,omitempty"`
	UserExperienceAnalyticsAppHealthDevicePerformanceDetails                   *[]UserExperienceAnalyticsAppHealthDevicePerformanceDetails           `json:"userExperienceAnalyticsAppHealthDevicePerformanceDetails,omitempty"`
	UserExperienceAnalyticsAppHealthOSVersionPerformance                       *[]UserExperienceAnalyticsAppHealthOSVersionPerformance               `json:"userExperienceAnalyticsAppHealthOSVersionPerformance,omitempty"`
	UserExperienceAnalyticsAppHealthOverview                                   *UserExperienceAnalyticsCategory                                      `json:"userExperienceAnalyticsAppHealthOverview,omitempty"`
	UserExperienceAnalyticsBaselines                                           *[]UserExperienceAnalyticsBaseline                                    `json:"userExperienceAnalyticsBaselines,omitempty"`
	UserExperienceAnalyticsCategories                                          *[]UserExperienceAnalyticsCategory                                    `json:"userExperienceAnalyticsCategories,omitempty"`
	UserExperienceAnalyticsDevicePerformance                                   *[]UserExperienceAnalyticsDevicePerformance                           `json:"userExperienceAnalyticsDevicePerformance,omitempty"`
	UserExperienceAnalyticsDeviceScores                                        *[]UserExperienceAnalyticsDeviceScores                                `json:"userExperienceAnalyticsDeviceScores,omitempty"`
	UserExperienceAnalyticsDeviceStartupHistory                                *[]UserExperienceAnalyticsDeviceStartupHistory                        `json:"userExperienceAnalyticsDeviceStartupHistory,omitempty"`
	UserExperienceAnalyticsDeviceStartupProcesses                              *[]UserExperienceAnalyticsDeviceStartupProcess                        `json:"userExperienceAnalyticsDeviceStartupProcesses,omitempty"`
	UserExperienceAnalyticsMetricHistory                                       *[]UserExperienceAnalyticsMetricHistory                               `json:"userExperienceAnalyticsMetricHistory,omitempty"`
	UserExperienceAnalyticsModelScores                                         *[]UserExperienceAnalyticsModelScores                                 `json:"userExperienceAnalyticsModelScores,omitempty"`
	UserExperienceAnalyticsOverview                                            *UserExperienceAnalyticsOverview                                      `json:"userExperienceAnalyticsOverview,omitempty"`
	UserExperienceAnalyticsScoreHistory                                        *[]UserExperienceAnalyticsScoreHistory                                `json:"userExperienceAnalyticsScoreHistory,omitempty"`
	UserExperienceAnalyticsSettings                                            *UserExperienceAnalyticsSettings                                      `json:"userExperienceAnalyticsSettings,omitempty"`
	UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric             *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric       `json:"userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric,omitempty"`
	UserExperienceAnalyticsWorkFromAnywhereMetrics                             *[]UserExperienceAnalyticsWorkFromAnywhereMetric                      `json:"userExperienceAnalyticsWorkFromAnywhereMetrics,omitempty"`
	UserExperienceAnalyticsWorkFromAnywhereModelPerformance                    *[]UserExperienceAnalyticsWorkFromAnywhereModelPerformance            `json:"userExperienceAnalyticsWorkFromAnywhereModelPerformance,omitempty"`
	WindowsAutopilotDeviceIdentities                                           *[]WindowsAutopilotDeviceIdentity                                     `json:"windowsAutopilotDeviceIdentities,omitempty"`
	WindowsInformationProtectionAppLearningSummaries                           *[]WindowsInformationProtectionAppLearningSummary                     `json:"windowsInformationProtectionAppLearningSummaries,omitempty"`
	WindowsInformationProtectionNetworkLearningSummaries                       *[]WindowsInformationProtectionNetworkLearningSummary                 `json:"windowsInformationProtectionNetworkLearningSummaries,omitempty"`
	WindowsMalwareInformation                                                  *[]WindowsMalwareInformation                                          `json:"windowsMalwareInformation,omitempty"`
	WindowsMalwareOverview                                                     *WindowsMalwareOverview                                               `json:"windowsMalwareOverview,omitempty"`
}
