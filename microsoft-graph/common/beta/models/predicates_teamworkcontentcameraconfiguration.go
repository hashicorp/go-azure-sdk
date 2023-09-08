package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkContentCameraConfigurationOperationPredicate struct {
	IsContentCameraInverted     *bool
	IsContentCameraOptional     *bool
	IsContentEnhancementEnabled *bool
	ODataType                   *string
}

func (p TeamworkContentCameraConfigurationOperationPredicate) Matches(input TeamworkContentCameraConfiguration) bool {

	if p.IsContentCameraInverted != nil && (input.IsContentCameraInverted == nil || *p.IsContentCameraInverted != *input.IsContentCameraInverted) {
		return false
	}

	if p.IsContentCameraOptional != nil && (input.IsContentCameraOptional == nil || *p.IsContentCameraOptional != *input.IsContentCameraOptional) {
		return false
	}

	if p.IsContentEnhancementEnabled != nil && (input.IsContentEnhancementEnabled == nil || *p.IsContentEnhancementEnabled != *input.IsContentEnhancementEnabled) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
