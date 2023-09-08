package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementPartner struct {
	DisplayName                                          *string                                `json:"displayName,omitempty"`
	GroupsRequiringPartnerEnrollment                     *[]DeviceManagementPartnerAssignment   `json:"groupsRequiringPartnerEnrollment,omitempty"`
	Id                                                   *string                                `json:"id,omitempty"`
	IsConfigured                                         *bool                                  `json:"isConfigured,omitempty"`
	LastHeartbeatDateTime                                *string                                `json:"lastHeartbeatDateTime,omitempty"`
	ODataType                                            *string                                `json:"@odata.type,omitempty"`
	PartnerAppType                                       *DeviceManagementPartnerPartnerAppType `json:"partnerAppType,omitempty"`
	PartnerState                                         *DeviceManagementPartnerPartnerState   `json:"partnerState,omitempty"`
	SingleTenantAppId                                    *string                                `json:"singleTenantAppId,omitempty"`
	WhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime *string                                `json:"whenPartnerDevicesWillBeMarkedAsNonCompliantDateTime,omitempty"`
	WhenPartnerDevicesWillBeRemovedDateTime              *string                                `json:"whenPartnerDevicesWillBeRemovedDateTime,omitempty"`
}
