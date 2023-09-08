package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsInformationProtectionAppLearningSummaryOperationPredicate struct {
	ApplicationName *string
	DeviceCount     *int64
	Id              *string
	ODataType       *string
}

func (p WindowsInformationProtectionAppLearningSummaryOperationPredicate) Matches(input WindowsInformationProtectionAppLearningSummary) bool {

	if p.ApplicationName != nil && (input.ApplicationName == nil || *p.ApplicationName != *input.ApplicationName) {
		return false
	}

	if p.DeviceCount != nil && (input.DeviceCount == nil || *p.DeviceCount != *input.DeviceCount) {
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
