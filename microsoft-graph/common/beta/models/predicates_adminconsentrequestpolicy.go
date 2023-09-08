package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdminConsentRequestPolicyOperationPredicate struct {
	Id                    *string
	IsEnabled             *bool
	NotifyReviewers       *bool
	ODataType             *string
	RemindersEnabled      *bool
	RequestDurationInDays *int64
	Version               *int64
}

func (p AdminConsentRequestPolicyOperationPredicate) Matches(input AdminConsentRequestPolicy) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsEnabled != nil && (input.IsEnabled == nil || *p.IsEnabled != *input.IsEnabled) {
		return false
	}

	if p.NotifyReviewers != nil && (input.NotifyReviewers == nil || *p.NotifyReviewers != *input.NotifyReviewers) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RemindersEnabled != nil && (input.RemindersEnabled == nil || *p.RemindersEnabled != *input.RemindersEnabled) {
		return false
	}

	if p.RequestDurationInDays != nil && (input.RequestDurationInDays == nil || *p.RequestDurationInDays != *input.RequestDurationInDays) {
		return false
	}

	if p.Version != nil && (input.Version == nil || *p.Version != *input.Version) {
		return false
	}

	return true
}
