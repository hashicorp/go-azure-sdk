package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementAlertRuleOperationPredicate struct {
	Description  *string
	DisplayName  *string
	Enabled      *bool
	Id           *string
	IsSystemRule *bool
	ODataType    *string
}

func (p DeviceManagementAlertRuleOperationPredicate) Matches(input DeviceManagementAlertRule) bool {

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Enabled != nil && (input.Enabled == nil || *p.Enabled != *input.Enabled) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsSystemRule != nil && (input.IsSystemRule == nil || *p.IsSystemRule != *input.IsSystemRule) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
