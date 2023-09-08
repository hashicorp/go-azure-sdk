package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceComplianceLocalActionLockDeviceWithPasscode struct {
	GracePeriodInMinutes                 *int64  `json:"gracePeriodInMinutes,omitempty"`
	Id                                   *string `json:"id,omitempty"`
	ODataType                            *string `json:"@odata.type,omitempty"`
	Passcode                             *string `json:"passcode,omitempty"`
	PasscodeSignInFailureCountBeforeWipe *int64  `json:"passcodeSignInFailureCountBeforeWipe,omitempty"`
}
