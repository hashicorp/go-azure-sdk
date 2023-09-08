package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkDisplayScreenConfigurationOperationPredicate struct {
	BacklightBrightness   *int64
	BacklightTimeout      *string
	IsHighContrastEnabled *bool
	IsScreensaverEnabled  *bool
	ODataType             *string
	ScreensaverTimeout    *string
}

func (p TeamworkDisplayScreenConfigurationOperationPredicate) Matches(input TeamworkDisplayScreenConfiguration) bool {

	if p.BacklightBrightness != nil && (input.BacklightBrightness == nil || *p.BacklightBrightness != *input.BacklightBrightness) {
		return false
	}

	if p.BacklightTimeout != nil && (input.BacklightTimeout == nil || *p.BacklightTimeout != *input.BacklightTimeout) {
		return false
	}

	if p.IsHighContrastEnabled != nil && (input.IsHighContrastEnabled == nil || *p.IsHighContrastEnabled != *input.IsHighContrastEnabled) {
		return false
	}

	if p.IsScreensaverEnabled != nil && (input.IsScreensaverEnabled == nil || *p.IsScreensaverEnabled != *input.IsScreensaverEnabled) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ScreensaverTimeout != nil && (input.ScreensaverTimeout == nil || *p.ScreensaverTimeout != *input.ScreensaverTimeout) {
		return false
	}

	return true
}
