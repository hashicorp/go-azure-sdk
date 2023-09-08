package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicySettingMappingOperationPredicate struct {
	AdmxSettingDefinitionId   *string
	Id                        *string
	IntuneSettingDefinitionId *string
	IsMdmSupported            *bool
	MdmCspName                *string
	MdmMinimumOSVersion       *int64
	MdmSettingUri             *string
	ODataType                 *string
	ParentId                  *string
	SettingCategory           *string
	SettingDisplayName        *string
	SettingDisplayValue       *string
	SettingDisplayValueType   *string
	SettingName               *string
	SettingValue              *string
	SettingValueDisplayUnits  *string
	SettingValueType          *string
}

func (p GroupPolicySettingMappingOperationPredicate) Matches(input GroupPolicySettingMapping) bool {

	if p.AdmxSettingDefinitionId != nil && (input.AdmxSettingDefinitionId == nil || *p.AdmxSettingDefinitionId != *input.AdmxSettingDefinitionId) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IntuneSettingDefinitionId != nil && (input.IntuneSettingDefinitionId == nil || *p.IntuneSettingDefinitionId != *input.IntuneSettingDefinitionId) {
		return false
	}

	if p.IsMdmSupported != nil && (input.IsMdmSupported == nil || *p.IsMdmSupported != *input.IsMdmSupported) {
		return false
	}

	if p.MdmCspName != nil && (input.MdmCspName == nil || *p.MdmCspName != *input.MdmCspName) {
		return false
	}

	if p.MdmMinimumOSVersion != nil && (input.MdmMinimumOSVersion == nil || *p.MdmMinimumOSVersion != *input.MdmMinimumOSVersion) {
		return false
	}

	if p.MdmSettingUri != nil && (input.MdmSettingUri == nil || *p.MdmSettingUri != *input.MdmSettingUri) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ParentId != nil && (input.ParentId == nil || *p.ParentId != *input.ParentId) {
		return false
	}

	if p.SettingCategory != nil && (input.SettingCategory == nil || *p.SettingCategory != *input.SettingCategory) {
		return false
	}

	if p.SettingDisplayName != nil && (input.SettingDisplayName == nil || *p.SettingDisplayName != *input.SettingDisplayName) {
		return false
	}

	if p.SettingDisplayValue != nil && (input.SettingDisplayValue == nil || *p.SettingDisplayValue != *input.SettingDisplayValue) {
		return false
	}

	if p.SettingDisplayValueType != nil && (input.SettingDisplayValueType == nil || *p.SettingDisplayValueType != *input.SettingDisplayValueType) {
		return false
	}

	if p.SettingName != nil && (input.SettingName == nil || *p.SettingName != *input.SettingName) {
		return false
	}

	if p.SettingValue != nil && (input.SettingValue == nil || *p.SettingValue != *input.SettingValue) {
		return false
	}

	if p.SettingValueDisplayUnits != nil && (input.SettingValueDisplayUnits == nil || *p.SettingValueDisplayUnits != *input.SettingValueDisplayUnits) {
		return false
	}

	if p.SettingValueType != nil && (input.SettingValueType == nil || *p.SettingValueType != *input.SettingValueType) {
		return false
	}

	return true
}
