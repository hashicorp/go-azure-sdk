package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamFunSettingsOperationPredicate struct {
	AllowCustomMemes      *bool
	AllowGiphy            *bool
	AllowStickersAndMemes *bool
	ODataType             *string
}

func (p TeamFunSettingsOperationPredicate) Matches(input TeamFunSettings) bool {

	if p.AllowCustomMemes != nil && (input.AllowCustomMemes == nil || *p.AllowCustomMemes != *input.AllowCustomMemes) {
		return false
	}

	if p.AllowGiphy != nil && (input.AllowGiphy == nil || *p.AllowGiphy != *input.AllowGiphy) {
		return false
	}

	if p.AllowStickersAndMemes != nil && (input.AllowStickersAndMemes == nil || *p.AllowStickersAndMemes != *input.AllowStickersAndMemes) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
