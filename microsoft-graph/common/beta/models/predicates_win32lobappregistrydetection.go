package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppRegistryDetectionOperationPredicate struct {
	Check32BitOn64System *bool
	DetectionValue       *string
	KeyPath              *string
	ODataType            *string
	ValueName            *string
}

func (p Win32LobAppRegistryDetectionOperationPredicate) Matches(input Win32LobAppRegistryDetection) bool {

	if p.Check32BitOn64System != nil && (input.Check32BitOn64System == nil || *p.Check32BitOn64System != *input.Check32BitOn64System) {
		return false
	}

	if p.DetectionValue != nil && (input.DetectionValue == nil || *p.DetectionValue != *input.DetectionValue) {
		return false
	}

	if p.KeyPath != nil && (input.KeyPath == nil || *p.KeyPath != *input.KeyPath) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ValueName != nil && (input.ValueName == nil || *p.ValueName != *input.ValueName) {
		return false
	}

	return true
}
