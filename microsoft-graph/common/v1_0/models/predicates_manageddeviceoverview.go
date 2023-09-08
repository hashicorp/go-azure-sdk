package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceOverviewOperationPredicate struct {
	DualEnrolledDeviceCount *int64
	EnrolledDeviceCount     *int64
	Id                      *string
	MdmEnrolledCount        *int64
	ODataType               *string
}

func (p ManagedDeviceOverviewOperationPredicate) Matches(input ManagedDeviceOverview) bool {

	if p.DualEnrolledDeviceCount != nil && (input.DualEnrolledDeviceCount == nil || *p.DualEnrolledDeviceCount != *input.DualEnrolledDeviceCount) {
		return false
	}

	if p.EnrolledDeviceCount != nil && (input.EnrolledDeviceCount == nil || *p.EnrolledDeviceCount != *input.EnrolledDeviceCount) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.MdmEnrolledCount != nil && (input.MdmEnrolledCount == nil || *p.MdmEnrolledCount != *input.MdmEnrolledCount) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
