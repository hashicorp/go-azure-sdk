package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSPrivacyAccessControlItemOperationPredicate struct {
	BlockCamera          *bool
	BlockListenEvent     *bool
	BlockMicrophone      *bool
	BlockScreenCapture   *bool
	CodeRequirement      *string
	DisplayName          *string
	Identifier           *string
	ODataType            *string
	StaticCodeValidation *bool
}

func (p MacOSPrivacyAccessControlItemOperationPredicate) Matches(input MacOSPrivacyAccessControlItem) bool {

	if p.BlockCamera != nil && (input.BlockCamera == nil || *p.BlockCamera != *input.BlockCamera) {
		return false
	}

	if p.BlockListenEvent != nil && (input.BlockListenEvent == nil || *p.BlockListenEvent != *input.BlockListenEvent) {
		return false
	}

	if p.BlockMicrophone != nil && (input.BlockMicrophone == nil || *p.BlockMicrophone != *input.BlockMicrophone) {
		return false
	}

	if p.BlockScreenCapture != nil && (input.BlockScreenCapture == nil || *p.BlockScreenCapture != *input.BlockScreenCapture) {
		return false
	}

	if p.CodeRequirement != nil && (input.CodeRequirement == nil || *p.CodeRequirement != *input.CodeRequirement) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Identifier != nil && (input.Identifier == nil || *p.Identifier != *input.Identifier) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.StaticCodeValidation != nil && (input.StaticCodeValidation == nil || *p.StaticCodeValidation != *input.StaticCodeValidation) {
		return false
	}

	return true
}
