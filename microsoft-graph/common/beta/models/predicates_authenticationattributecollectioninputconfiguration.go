package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationAttributeCollectionInputConfigurationOperationPredicate struct {
	Attribute        *string
	DefaultValue     *string
	Editable         *bool
	Hidden           *bool
	Label            *string
	ODataType        *string
	Required         *bool
	ValidationRegEx  *string
	WriteToDirectory *bool
}

func (p AuthenticationAttributeCollectionInputConfigurationOperationPredicate) Matches(input AuthenticationAttributeCollectionInputConfiguration) bool {

	if p.Attribute != nil && (input.Attribute == nil || *p.Attribute != *input.Attribute) {
		return false
	}

	if p.DefaultValue != nil && (input.DefaultValue == nil || *p.DefaultValue != *input.DefaultValue) {
		return false
	}

	if p.Editable != nil && (input.Editable == nil || *p.Editable != *input.Editable) {
		return false
	}

	if p.Hidden != nil && (input.Hidden == nil || *p.Hidden != *input.Hidden) {
		return false
	}

	if p.Label != nil && (input.Label == nil || *p.Label != *input.Label) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Required != nil && (input.Required == nil || *p.Required != *input.Required) {
		return false
	}

	if p.ValidationRegEx != nil && (input.ValidationRegEx == nil || *p.ValidationRegEx != *input.ValidationRegEx) {
		return false
	}

	if p.WriteToDirectory != nil && (input.WriteToDirectory == nil || *p.WriteToDirectory != *input.WriteToDirectory) {
		return false
	}

	return true
}
