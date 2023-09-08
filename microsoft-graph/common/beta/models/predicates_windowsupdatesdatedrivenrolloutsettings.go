package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesDateDrivenRolloutSettingsOperationPredicate struct {
	DurationBetweenOffers *string
	EndDateTime           *string
	ODataType             *string
}

func (p WindowsUpdatesDateDrivenRolloutSettingsOperationPredicate) Matches(input WindowsUpdatesDateDrivenRolloutSettings) bool {

	if p.DurationBetweenOffers != nil && (input.DurationBetweenOffers == nil || *p.DurationBetweenOffers != *input.DurationBetweenOffers) {
		return false
	}

	if p.EndDateTime != nil && (input.EndDateTime == nil || *p.EndDateTime != *input.EndDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
