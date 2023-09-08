package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceConfigurationConflictSummaryOperationPredicate struct {
	DeviceCheckinsImpacted *int64
	Id                     *string
	ODataType              *string
}

func (p DeviceConfigurationConflictSummaryOperationPredicate) Matches(input DeviceConfigurationConflictSummary) bool {

	if p.DeviceCheckinsImpacted != nil && (input.DeviceCheckinsImpacted == nil || *p.DeviceCheckinsImpacted != *input.DeviceCheckinsImpacted) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
