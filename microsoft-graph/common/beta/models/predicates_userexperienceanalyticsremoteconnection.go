package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsRemoteConnectionOperationPredicate struct {
	DeviceCount       *int64
	DeviceId          *string
	DeviceName        *string
	Id                *string
	Manufacturer      *string
	Model             *string
	ODataType         *string
	UserPrincipalName *string
	VirtualNetwork    *string
}

func (p UserExperienceAnalyticsRemoteConnectionOperationPredicate) Matches(input UserExperienceAnalyticsRemoteConnection) bool {

	if p.DeviceCount != nil && (input.DeviceCount == nil || *p.DeviceCount != *input.DeviceCount) {
		return false
	}

	if p.DeviceId != nil && (input.DeviceId == nil || *p.DeviceId != *input.DeviceId) {
		return false
	}

	if p.DeviceName != nil && (input.DeviceName == nil || *p.DeviceName != *input.DeviceName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Manufacturer != nil && (input.Manufacturer == nil || *p.Manufacturer != *input.Manufacturer) {
		return false
	}

	if p.Model != nil && (input.Model == nil || *p.Model != *input.Model) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.UserPrincipalName != nil && (input.UserPrincipalName == nil || *p.UserPrincipalName != *input.UserPrincipalName) {
		return false
	}

	if p.VirtualNetwork != nil && (input.VirtualNetwork == nil || *p.VirtualNetwork != *input.VirtualNetwork) {
		return false
	}

	return true
}
