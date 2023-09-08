package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppRequirementOperationPredicate struct {
	DetectionValue *string
	ODataType      *string
}

func (p Win32LobAppRequirementOperationPredicate) Matches(input Win32LobAppRequirement) bool {

	if p.DetectionValue != nil && (input.DetectionValue == nil || *p.DetectionValue != *input.DetectionValue) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
