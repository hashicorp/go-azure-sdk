package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyMigrationReport struct {
	CreatedDateTime                  *string                                       `json:"createdDateTime,omitempty"`
	DisplayName                      *string                                       `json:"displayName,omitempty"`
	GroupPolicyCreatedDateTime       *string                                       `json:"groupPolicyCreatedDateTime,omitempty"`
	GroupPolicyLastModifiedDateTime  *string                                       `json:"groupPolicyLastModifiedDateTime,omitempty"`
	GroupPolicyObjectId              *string                                       `json:"groupPolicyObjectId,omitempty"`
	GroupPolicySettingMappings       *[]GroupPolicySettingMapping                  `json:"groupPolicySettingMappings,omitempty"`
	Id                               *string                                       `json:"id,omitempty"`
	LastModifiedDateTime             *string                                       `json:"lastModifiedDateTime,omitempty"`
	MigrationReadiness               *GroupPolicyMigrationReportMigrationReadiness `json:"migrationReadiness,omitempty"`
	ODataType                        *string                                       `json:"@odata.type,omitempty"`
	OuDistinguishedName              *string                                       `json:"ouDistinguishedName,omitempty"`
	RoleScopeTagIds                  *[]string                                     `json:"roleScopeTagIds,omitempty"`
	SupportedSettingsCount           *int64                                        `json:"supportedSettingsCount,omitempty"`
	SupportedSettingsPercent         *int64                                        `json:"supportedSettingsPercent,omitempty"`
	TargetedInActiveDirectory        *bool                                         `json:"targetedInActiveDirectory,omitempty"`
	TotalSettingsCount               *int64                                        `json:"totalSettingsCount,omitempty"`
	UnsupportedGroupPolicyExtensions *[]UnsupportedGroupPolicyExtension            `json:"unsupportedGroupPolicyExtensions,omitempty"`
}
