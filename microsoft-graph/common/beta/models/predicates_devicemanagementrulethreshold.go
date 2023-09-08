package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementRuleThresholdOperationPredicate struct {
	ODataType *string
	Target    *int64
}

func (p DeviceManagementRuleThresholdOperationPredicate) Matches(input DeviceManagementRuleThreshold) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Target != nil && (input.Target == nil || *p.Target != *input.Target) {
		return false
	}

	return true
}
