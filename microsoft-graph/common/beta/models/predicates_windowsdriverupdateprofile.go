package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsDriverUpdateProfileOperationPredicate struct {
	CreatedDateTime          *string
	DeploymentDeferralInDays *int64
	Description              *string
	DeviceReporting          *int64
	DisplayName              *string
	Id                       *string
	LastModifiedDateTime     *string
	NewUpdates               *int64
	ODataType                *string
}

func (p WindowsDriverUpdateProfileOperationPredicate) Matches(input WindowsDriverUpdateProfile) bool {

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.DeploymentDeferralInDays != nil && (input.DeploymentDeferralInDays == nil || *p.DeploymentDeferralInDays != *input.DeploymentDeferralInDays) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DeviceReporting != nil && (input.DeviceReporting == nil || *p.DeviceReporting != *input.DeviceReporting) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.NewUpdates != nil && (input.NewUpdates == nil || *p.NewUpdates != *input.NewUpdates) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
