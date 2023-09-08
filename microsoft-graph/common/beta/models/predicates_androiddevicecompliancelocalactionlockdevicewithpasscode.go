package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceComplianceLocalActionLockDeviceWithPasscodeOperationPredicate struct {
	GracePeriodInMinutes                 *int64
	Id                                   *string
	ODataType                            *string
	Passcode                             *string
	PasscodeSignInFailureCountBeforeWipe *int64
}

func (p AndroidDeviceComplianceLocalActionLockDeviceWithPasscodeOperationPredicate) Matches(input AndroidDeviceComplianceLocalActionLockDeviceWithPasscode) bool {

	if p.GracePeriodInMinutes != nil && (input.GracePeriodInMinutes == nil || *p.GracePeriodInMinutes != *input.GracePeriodInMinutes) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Passcode != nil && (input.Passcode == nil || *p.Passcode != *input.Passcode) {
		return false
	}

	if p.PasscodeSignInFailureCountBeforeWipe != nil && (input.PasscodeSignInFailureCountBeforeWipe == nil || *p.PasscodeSignInFailureCountBeforeWipe != *input.PasscodeSignInFailureCountBeforeWipe) {
		return false
	}

	return true
}
