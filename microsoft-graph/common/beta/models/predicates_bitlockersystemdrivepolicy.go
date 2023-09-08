package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BitLockerSystemDrivePolicyOperationPredicate struct {
	MinimumPinLength                         *int64
	ODataType                                *string
	PrebootRecoveryEnableMessageAndUrl       *bool
	PrebootRecoveryMessage                   *string
	PrebootRecoveryUrl                       *string
	StartupAuthenticationBlockWithoutTpmChip *bool
	StartupAuthenticationRequired            *bool
}

func (p BitLockerSystemDrivePolicyOperationPredicate) Matches(input BitLockerSystemDrivePolicy) bool {

	if p.MinimumPinLength != nil && (input.MinimumPinLength == nil || *p.MinimumPinLength != *input.MinimumPinLength) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PrebootRecoveryEnableMessageAndUrl != nil && (input.PrebootRecoveryEnableMessageAndUrl == nil || *p.PrebootRecoveryEnableMessageAndUrl != *input.PrebootRecoveryEnableMessageAndUrl) {
		return false
	}

	if p.PrebootRecoveryMessage != nil && (input.PrebootRecoveryMessage == nil || *p.PrebootRecoveryMessage != *input.PrebootRecoveryMessage) {
		return false
	}

	if p.PrebootRecoveryUrl != nil && (input.PrebootRecoveryUrl == nil || *p.PrebootRecoveryUrl != *input.PrebootRecoveryUrl) {
		return false
	}

	if p.StartupAuthenticationBlockWithoutTpmChip != nil && (input.StartupAuthenticationBlockWithoutTpmChip == nil || *p.StartupAuthenticationBlockWithoutTpmChip != *input.StartupAuthenticationBlockWithoutTpmChip) {
		return false
	}

	if p.StartupAuthenticationRequired != nil && (input.StartupAuthenticationRequired == nil || *p.StartupAuthenticationRequired != *input.StartupAuthenticationRequired) {
		return false
	}

	return true
}
