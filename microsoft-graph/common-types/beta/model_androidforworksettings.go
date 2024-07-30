package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkSettings struct {
	BindStatus                   *AndroidForWorkSettingsBindStatus        `json:"bindStatus,omitempty"`
	DeviceOwnerManagementEnabled *bool                                    `json:"deviceOwnerManagementEnabled,omitempty"`
	EnrollmentTarget             *AndroidForWorkSettingsEnrollmentTarget  `json:"enrollmentTarget,omitempty"`
	Id                           *string                                  `json:"id,omitempty"`
	LastAppSyncDateTime          *string                                  `json:"lastAppSyncDateTime,omitempty"`
	LastAppSyncStatus            *AndroidForWorkSettingsLastAppSyncStatus `json:"lastAppSyncStatus,omitempty"`
	LastModifiedDateTime         *string                                  `json:"lastModifiedDateTime,omitempty"`
	ODataType                    *string                                  `json:"@odata.type,omitempty"`
	OwnerOrganizationName        *string                                  `json:"ownerOrganizationName,omitempty"`
	OwnerUserPrincipalName       *string                                  `json:"ownerUserPrincipalName,omitempty"`
	TargetGroupIds               *[]string                                `json:"targetGroupIds,omitempty"`
}
