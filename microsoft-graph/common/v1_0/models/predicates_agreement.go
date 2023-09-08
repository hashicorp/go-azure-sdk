package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AgreementOperationPredicate struct {
	DisplayName                       *string
	Id                                *string
	IsPerDeviceAcceptanceRequired     *bool
	IsViewingBeforeAcceptanceRequired *bool
	ODataType                         *string
	UserReacceptRequiredFrequency     *string
}

func (p AgreementOperationPredicate) Matches(input Agreement) bool {

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsPerDeviceAcceptanceRequired != nil && (input.IsPerDeviceAcceptanceRequired == nil || *p.IsPerDeviceAcceptanceRequired != *input.IsPerDeviceAcceptanceRequired) {
		return false
	}

	if p.IsViewingBeforeAcceptanceRequired != nil && (input.IsViewingBeforeAcceptanceRequired == nil || *p.IsViewingBeforeAcceptanceRequired != *input.IsViewingBeforeAcceptanceRequired) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.UserReacceptRequiredFrequency != nil && (input.UserReacceptRequiredFrequency == nil || *p.UserReacceptRequiredFrequency != *input.UserReacceptRequiredFrequency) {
		return false
	}

	return true
}
