package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAutomaticRequestSettingsOperationPredicate struct {
	GracePeriodBeforeAccessRemoval             *string
	ODataType                                  *string
	RemoveAccessWhenTargetLeavesAllowedTargets *bool
	RequestAccessForAllowedTargets             *bool
}

func (p AccessPackageAutomaticRequestSettingsOperationPredicate) Matches(input AccessPackageAutomaticRequestSettings) bool {

	if p.GracePeriodBeforeAccessRemoval != nil && (input.GracePeriodBeforeAccessRemoval == nil || *p.GracePeriodBeforeAccessRemoval != *input.GracePeriodBeforeAccessRemoval) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RemoveAccessWhenTargetLeavesAllowedTargets != nil && (input.RemoveAccessWhenTargetLeavesAllowedTargets == nil || *p.RemoveAccessWhenTargetLeavesAllowedTargets != *input.RemoveAccessWhenTargetLeavesAllowedTargets) {
		return false
	}

	if p.RequestAccessForAllowedTargets != nil && (input.RequestAccessForAllowedTargets == nil || *p.RequestAccessForAllowedTargets != *input.RequestAccessForAllowedTargets) {
		return false
	}

	return true
}
