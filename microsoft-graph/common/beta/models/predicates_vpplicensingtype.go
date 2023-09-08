package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VppLicensingTypeOperationPredicate struct {
	ODataType               *string
	SupportDeviceLicensing  *bool
	SupportUserLicensing    *bool
	SupportsDeviceLicensing *bool
	SupportsUserLicensing   *bool
}

func (p VppLicensingTypeOperationPredicate) Matches(input VppLicensingType) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SupportDeviceLicensing != nil && (input.SupportDeviceLicensing == nil || *p.SupportDeviceLicensing != *input.SupportDeviceLicensing) {
		return false
	}

	if p.SupportUserLicensing != nil && (input.SupportUserLicensing == nil || *p.SupportUserLicensing != *input.SupportUserLicensing) {
		return false
	}

	if p.SupportsDeviceLicensing != nil && (input.SupportsDeviceLicensing == nil || *p.SupportsDeviceLicensing != *input.SupportsDeviceLicensing) {
		return false
	}

	if p.SupportsUserLicensing != nil && (input.SupportsUserLicensing == nil || *p.SupportsUserLicensing != *input.SupportsUserLicensing) {
		return false
	}

	return true
}
