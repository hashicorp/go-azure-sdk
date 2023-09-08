package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationSignInDetailedSummaryOperationPredicate struct {
	AggregatedEventDateTime *string
	AppDisplayName          *string
	AppId                   *string
	Id                      *string
	ODataType               *string
	SignInCount             *int64
}

func (p ApplicationSignInDetailedSummaryOperationPredicate) Matches(input ApplicationSignInDetailedSummary) bool {

	if p.AggregatedEventDateTime != nil && (input.AggregatedEventDateTime == nil || *p.AggregatedEventDateTime != *input.AggregatedEventDateTime) {
		return false
	}

	if p.AppDisplayName != nil && (input.AppDisplayName == nil || *p.AppDisplayName != *input.AppDisplayName) {
		return false
	}

	if p.AppId != nil && (input.AppId == nil || *p.AppId != *input.AppId) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SignInCount != nil && (input.SignInCount == nil || *p.SignInCount != *input.SignInCount) {
		return false
	}

	return true
}
