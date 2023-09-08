package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WatermarkProtectionValuesOperationPredicate struct {
	IsEnabledForContentSharing *bool
	IsEnabledForVideo          *bool
	ODataType                  *string
}

func (p WatermarkProtectionValuesOperationPredicate) Matches(input WatermarkProtectionValues) bool {

	if p.IsEnabledForContentSharing != nil && (input.IsEnabledForContentSharing == nil || *p.IsEnabledForContentSharing != *input.IsEnabledForContentSharing) {
		return false
	}

	if p.IsEnabledForVideo != nil && (input.IsEnabledForVideo == nil || *p.IsEnabledForVideo != *input.IsEnabledForVideo) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
