package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementComplexSettingInstanceOperationPredicate struct {
	DefinitionId *string
	Id           *string
	ODataType    *string
	ValueJson    *string
}

func (p DeviceManagementComplexSettingInstanceOperationPredicate) Matches(input DeviceManagementComplexSettingInstance) bool {

	if p.DefinitionId != nil && (input.DefinitionId == nil || *p.DefinitionId != *input.DefinitionId) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ValueJson != nil && (input.ValueJson == nil || *p.ValueJson != *input.ValueJson) {
		return false
	}

	return true
}
