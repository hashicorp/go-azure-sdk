package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSWiredNetworkConfiguration struct {
	Assignments                                 *[]DeviceConfigurationAssignment                                    `json:"assignments,omitempty"`
	AuthenticationMethod                        *MacOSWiredNetworkConfigurationAuthenticationMethod                 `json:"authenticationMethod,omitempty"`
	CreatedDateTime                             *string                                                             `json:"createdDateTime,omitempty"`
	Description                                 *string                                                             `json:"description,omitempty"`
	DeviceManagementApplicabilityRuleDeviceMode *DeviceManagementApplicabilityRuleDeviceMode                        `json:"deviceManagementApplicabilityRuleDeviceMode,omitempty"`
	DeviceManagementApplicabilityRuleOsEdition  *DeviceManagementApplicabilityRuleOsEdition                         `json:"deviceManagementApplicabilityRuleOsEdition,omitempty"`
	DeviceManagementApplicabilityRuleOsVersion  *DeviceManagementApplicabilityRuleOsVersion                         `json:"deviceManagementApplicabilityRuleOsVersion,omitempty"`
	DeviceSettingStateSummaries                 *[]SettingStateDeviceSummary                                        `json:"deviceSettingStateSummaries,omitempty"`
	DeviceStatusOverview                        *DeviceConfigurationDeviceOverview                                  `json:"deviceStatusOverview,omitempty"`
	DeviceStatuses                              *[]DeviceConfigurationDeviceStatus                                  `json:"deviceStatuses,omitempty"`
	DisplayName                                 *string                                                             `json:"displayName,omitempty"`
	EapFastConfiguration                        *MacOSWiredNetworkConfigurationEapFastConfiguration                 `json:"eapFastConfiguration,omitempty"`
	EapType                                     *MacOSWiredNetworkConfigurationEapType                              `json:"eapType,omitempty"`
	EnableOuterIdentityPrivacy                  *string                                                             `json:"enableOuterIdentityPrivacy,omitempty"`
	GroupAssignments                            *[]DeviceConfigurationGroupAssignment                               `json:"groupAssignments,omitempty"`
	Id                                          *string                                                             `json:"id,omitempty"`
	IdentityCertificateForClientAuthentication  *MacOSCertificateProfileBase                                        `json:"identityCertificateForClientAuthentication,omitempty"`
	LastModifiedDateTime                        *string                                                             `json:"lastModifiedDateTime,omitempty"`
	NetworkInterface                            *MacOSWiredNetworkConfigurationNetworkInterface                     `json:"networkInterface,omitempty"`
	NetworkName                                 *string                                                             `json:"networkName,omitempty"`
	NonEapAuthenticationMethodForEapTtls        *MacOSWiredNetworkConfigurationNonEapAuthenticationMethodForEapTtls `json:"nonEapAuthenticationMethodForEapTtls,omitempty"`
	ODataType                                   *string                                                             `json:"@odata.type,omitempty"`
	RoleScopeTagIds                             *[]string                                                           `json:"roleScopeTagIds,omitempty"`
	RootCertificateForServerValidation          *MacOSTrustedRootCertificate                                        `json:"rootCertificateForServerValidation,omitempty"`
	SupportsScopeTags                           *bool                                                               `json:"supportsScopeTags,omitempty"`
	TrustedServerCertificateNames               *[]string                                                           `json:"trustedServerCertificateNames,omitempty"`
	UserStatusOverview                          *DeviceConfigurationUserOverview                                    `json:"userStatusOverview,omitempty"`
	UserStatuses                                *[]DeviceConfigurationUserStatus                                    `json:"userStatuses,omitempty"`
	Version                                     *int64                                                              `json:"version,omitempty"`
}
