package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementIntentOperationPredicate struct {
	Description                      *string
	DisplayName                      *string
	Id                               *string
	IsAssigned                       *bool
	IsMigratingToConfigurationPolicy *bool
	LastModifiedDateTime             *string
	ODataType                        *string
	TemplateId                       *string
}

func (p DeviceManagementIntentOperationPredicate) Matches(input DeviceManagementIntent) bool {

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsAssigned != nil && (input.IsAssigned == nil || *p.IsAssigned != *input.IsAssigned) {
		return false
	}

	if p.IsMigratingToConfigurationPolicy != nil && (input.IsMigratingToConfigurationPolicy == nil || *p.IsMigratingToConfigurationPolicy != *input.IsMigratingToConfigurationPolicy) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.TemplateId != nil && (input.TemplateId == nil || *p.TemplateId != *input.TemplateId) {
		return false
	}

	return true
}
