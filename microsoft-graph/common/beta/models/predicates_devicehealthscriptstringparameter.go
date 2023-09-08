package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceHealthScriptStringParameterOperationPredicate struct {
	ApplyDefaultValueWhenNotAssigned *bool
	DefaultValue                     *string
	Description                      *string
	IsRequired                       *bool
	Name                             *string
	ODataType                        *string
}

func (p DeviceHealthScriptStringParameterOperationPredicate) Matches(input DeviceHealthScriptStringParameter) bool {

	if p.ApplyDefaultValueWhenNotAssigned != nil && (input.ApplyDefaultValueWhenNotAssigned == nil || *p.ApplyDefaultValueWhenNotAssigned != *input.ApplyDefaultValueWhenNotAssigned) {
		return false
	}

	if p.DefaultValue != nil && (input.DefaultValue == nil || *p.DefaultValue != *input.DefaultValue) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.IsRequired != nil && (input.IsRequired == nil || *p.IsRequired != *input.IsRequired) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
