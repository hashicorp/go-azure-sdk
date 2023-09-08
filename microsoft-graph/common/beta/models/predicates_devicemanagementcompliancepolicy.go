package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementCompliancePolicyOperationPredicate struct {
	CreatedDateTime      *string
	CreationSource       *string
	Description          *string
	Id                   *string
	IsAssigned           *bool
	LastModifiedDateTime *string
	Name                 *string
	ODataType            *string
	SettingCount         *int64
}

func (p DeviceManagementCompliancePolicyOperationPredicate) Matches(input DeviceManagementCompliancePolicy) bool {

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.CreationSource != nil && (input.CreationSource == nil || *p.CreationSource != *input.CreationSource) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsAssigned != nil && (input.IsAssigned == nil || *p.IsAssigned != *input.IsAssigned) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SettingCount != nil && (input.SettingCount == nil || *p.SettingCount != *input.SettingCount) {
		return false
	}

	return true
}
