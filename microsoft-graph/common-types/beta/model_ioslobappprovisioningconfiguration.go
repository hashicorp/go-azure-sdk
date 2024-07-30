package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosLobAppProvisioningConfiguration struct {
	Assignments          *[]IosLobAppProvisioningConfigurationAssignment    `json:"assignments,omitempty"`
	CreatedDateTime      *string                                            `json:"createdDateTime,omitempty"`
	Description          *string                                            `json:"description,omitempty"`
	DeviceStatuses       *[]ManagedDeviceMobileAppConfigurationDeviceStatus `json:"deviceStatuses,omitempty"`
	DisplayName          *string                                            `json:"displayName,omitempty"`
	ExpirationDateTime   *string                                            `json:"expirationDateTime,omitempty"`
	GroupAssignments     *[]MobileAppProvisioningConfigGroupAssignment      `json:"groupAssignments,omitempty"`
	Id                   *string                                            `json:"id,omitempty"`
	LastModifiedDateTime *string                                            `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                                            `json:"@odata.type,omitempty"`
	Payload              *string                                            `json:"payload,omitempty"`
	PayloadFileName      *string                                            `json:"payloadFileName,omitempty"`
	RoleScopeTagIds      *[]string                                          `json:"roleScopeTagIds,omitempty"`
	UserStatuses         *[]ManagedDeviceMobileAppConfigurationUserStatus   `json:"userStatuses,omitempty"`
	Version              *int64                                             `json:"version,omitempty"`
}
