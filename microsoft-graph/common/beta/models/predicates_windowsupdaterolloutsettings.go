package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdateRolloutSettingsOperationPredicate struct {
	ODataType               *string
	OfferEndDateTimeInUTC   *string
	OfferIntervalInDays     *int64
	OfferStartDateTimeInUTC *string
}

func (p WindowsUpdateRolloutSettingsOperationPredicate) Matches(input WindowsUpdateRolloutSettings) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OfferEndDateTimeInUTC != nil && (input.OfferEndDateTimeInUTC == nil || *p.OfferEndDateTimeInUTC != *input.OfferEndDateTimeInUTC) {
		return false
	}

	if p.OfferIntervalInDays != nil && (input.OfferIntervalInDays == nil || *p.OfferIntervalInDays != *input.OfferIntervalInDays) {
		return false
	}

	if p.OfferStartDateTimeInUTC != nil && (input.OfferStartDateTimeInUTC == nil || *p.OfferStartDateTimeInUTC != *input.OfferStartDateTimeInUTC) {
		return false
	}

	return true
}
