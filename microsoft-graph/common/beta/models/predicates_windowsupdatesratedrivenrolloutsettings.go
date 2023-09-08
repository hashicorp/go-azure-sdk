package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesRateDrivenRolloutSettingsOperationPredicate struct {
	DevicesPerOffer       *int64
	DurationBetweenOffers *string
	ODataType             *string
}

func (p WindowsUpdatesRateDrivenRolloutSettingsOperationPredicate) Matches(input WindowsUpdatesRateDrivenRolloutSettings) bool {

	if p.DevicesPerOffer != nil && (input.DevicesPerOffer == nil || *p.DevicesPerOffer != *input.DevicesPerOffer) {
		return false
	}

	if p.DurationBetweenOffers != nil && (input.DurationBetweenOffers == nil || *p.DurationBetweenOffers != *input.DurationBetweenOffers) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
