package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetedManagedAppConfiguration struct {
	AppGroupType                *TargetedManagedAppConfigurationAppGroupType                `json:"appGroupType,omitempty"`
	Apps                        *[]ManagedMobileApp                                         `json:"apps,omitempty"`
	Assignments                 *[]TargetedManagedAppPolicyAssignment                       `json:"assignments,omitempty"`
	CreatedDateTime             *string                                                     `json:"createdDateTime,omitempty"`
	CustomSettings              *[]KeyValuePair                                             `json:"customSettings,omitempty"`
	DeployedAppCount            *int64                                                      `json:"deployedAppCount,omitempty"`
	DeploymentSummary           *ManagedAppPolicyDeploymentSummary                          `json:"deploymentSummary,omitempty"`
	Description                 *string                                                     `json:"description,omitempty"`
	DisplayName                 *string                                                     `json:"displayName,omitempty"`
	Id                          *string                                                     `json:"id,omitempty"`
	IsAssigned                  *bool                                                       `json:"isAssigned,omitempty"`
	LastModifiedDateTime        *string                                                     `json:"lastModifiedDateTime,omitempty"`
	ODataType                   *string                                                     `json:"@odata.type,omitempty"`
	RoleScopeTagIds             *[]string                                                   `json:"roleScopeTagIds,omitempty"`
	Settings                    *[]DeviceManagementConfigurationSetting                     `json:"settings,omitempty"`
	TargetedAppManagementLevels *TargetedManagedAppConfigurationTargetedAppManagementLevels `json:"targetedAppManagementLevels,omitempty"`
	Version                     *string                                                     `json:"version,omitempty"`
}
