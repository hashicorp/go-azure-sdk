package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActivateDeviceEsimActionResultOperationPredicate struct {
	ActionName          *string
	CarrierUrl          *string
	LastUpdatedDateTime *string
	ODataType           *string
	StartDateTime       *string
}

func (p ActivateDeviceEsimActionResultOperationPredicate) Matches(input ActivateDeviceEsimActionResult) bool {

	if p.ActionName != nil && (input.ActionName == nil || *p.ActionName != *input.ActionName) {
		return false
	}

	if p.CarrierUrl != nil && (input.CarrierUrl == nil || *p.CarrierUrl != *input.CarrierUrl) {
		return false
	}

	if p.LastUpdatedDateTime != nil && (input.LastUpdatedDateTime == nil || *p.LastUpdatedDateTime != *input.LastUpdatedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.StartDateTime != nil && (input.StartDateTime == nil || *p.StartDateTime != *input.StartDateTime) {
		return false
	}

	return true
}
