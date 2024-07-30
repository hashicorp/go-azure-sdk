package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileEasEmailProfileBase struct {
	Assignments                                 *[]DeviceConfigurationAssignment                            `json:"assignments,omitempty"`
	AuthenticationMethod                        *AndroidWorkProfileEasEmailProfileBaseAuthenticationMethod  `json:"authenticationMethod,omitempty"`
	CreatedDateTime                             *string                                                     `json:"createdDateTime,omitempty"`
	Description                                 *string                                                     `json:"description,omitempty"`
	DeviceManagementApplicabilityRuleDeviceMode *DeviceManagementApplicabilityRuleDeviceMode                `json:"deviceManagementApplicabilityRuleDeviceMode,omitempty"`
	DeviceManagementApplicabilityRuleOsEdition  *DeviceManagementApplicabilityRuleOsEdition                 `json:"deviceManagementApplicabilityRuleOsEdition,omitempty"`
	DeviceManagementApplicabilityRuleOsVersion  *DeviceManagementApplicabilityRuleOsVersion                 `json:"deviceManagementApplicabilityRuleOsVersion,omitempty"`
	DeviceSettingStateSummaries                 *[]SettingStateDeviceSummary                                `json:"deviceSettingStateSummaries,omitempty"`
	DeviceStatusOverview                        *DeviceConfigurationDeviceOverview                          `json:"deviceStatusOverview,omitempty"`
	DeviceStatuses                              *[]DeviceConfigurationDeviceStatus                          `json:"deviceStatuses,omitempty"`
	DisplayName                                 *string                                                     `json:"displayName,omitempty"`
	DurationOfEmailToSync                       *AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSync `json:"durationOfEmailToSync,omitempty"`
	EmailAddressSource                          *AndroidWorkProfileEasEmailProfileBaseEmailAddressSource    `json:"emailAddressSource,omitempty"`
	GroupAssignments                            *[]DeviceConfigurationGroupAssignment                       `json:"groupAssignments,omitempty"`
	HostName                                    *string                                                     `json:"hostName,omitempty"`
	Id                                          *string                                                     `json:"id,omitempty"`
	IdentityCertificate                         *AndroidWorkProfileCertificateProfileBase                   `json:"identityCertificate,omitempty"`
	LastModifiedDateTime                        *string                                                     `json:"lastModifiedDateTime,omitempty"`
	ODataType                                   *string                                                     `json:"@odata.type,omitempty"`
	RequireSsl                                  *bool                                                       `json:"requireSsl,omitempty"`
	RoleScopeTagIds                             *[]string                                                   `json:"roleScopeTagIds,omitempty"`
	SupportsScopeTags                           *bool                                                       `json:"supportsScopeTags,omitempty"`
	UserStatusOverview                          *DeviceConfigurationUserOverview                            `json:"userStatusOverview,omitempty"`
	UserStatuses                                *[]DeviceConfigurationUserStatus                            `json:"userStatuses,omitempty"`
	UsernameSource                              *AndroidWorkProfileEasEmailProfileBaseUsernameSource        `json:"usernameSource,omitempty"`
	Version                                     *int64                                                      `json:"version,omitempty"`
}
