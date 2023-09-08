package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementPartnerOperationPredicate struct {
	DisplayName                                          *string
	Id                                                   *string
	IsConfigured                                         *bool
	LastHeartbeatDateTime                                *string
	ODataType                                            *string
	SingleTenantAppId                                    *string
	WhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime *string
	WhenPartnerDevicesWillBeRemovedDateTime              *string
}

func (p DeviceManagementPartnerOperationPredicate) Matches(input DeviceManagementPartner) bool {

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsConfigured != nil && (input.IsConfigured == nil || *p.IsConfigured != *input.IsConfigured) {
		return false
	}

	if p.LastHeartbeatDateTime != nil && (input.LastHeartbeatDateTime == nil || *p.LastHeartbeatDateTime != *input.LastHeartbeatDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SingleTenantAppId != nil && (input.SingleTenantAppId == nil || *p.SingleTenantAppId != *input.SingleTenantAppId) {
		return false
	}

	if p.WhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime != nil && (input.WhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime == nil || *p.WhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime != *input.WhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime) {
		return false
	}

	if p.WhenPartnerDevicesWillBeRemovedDateTime != nil && (input.WhenPartnerDevicesWillBeRemovedDateTime == nil || *p.WhenPartnerDevicesWillBeRemovedDateTime != *input.WhenPartnerDevicesWillBeRemovedDateTime) {
		return false
	}

	return true
}
