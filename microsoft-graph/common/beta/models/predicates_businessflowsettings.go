package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BusinessFlowSettingsOperationPredicate struct {
	AccessRecommendationsEnabled    *bool
	ActivityDurationInDays          *int64
	AutoApplyReviewResultsEnabled   *bool
	AutoReviewEnabled               *bool
	DurationInDays                  *int64
	JustificationRequiredOnApproval *bool
	MailNotificationsEnabled        *bool
	ODataType                       *string
	RemindersEnabled                *bool
}

func (p BusinessFlowSettingsOperationPredicate) Matches(input BusinessFlowSettings) bool {

	if p.AccessRecommendationsEnabled != nil && (input.AccessRecommendationsEnabled == nil || *p.AccessRecommendationsEnabled != *input.AccessRecommendationsEnabled) {
		return false
	}

	if p.ActivityDurationInDays != nil && (input.ActivityDurationInDays == nil || *p.ActivityDurationInDays != *input.ActivityDurationInDays) {
		return false
	}

	if p.AutoApplyReviewResultsEnabled != nil && (input.AutoApplyReviewResultsEnabled == nil || *p.AutoApplyReviewResultsEnabled != *input.AutoApplyReviewResultsEnabled) {
		return false
	}

	if p.AutoReviewEnabled != nil && (input.AutoReviewEnabled == nil || *p.AutoReviewEnabled != *input.AutoReviewEnabled) {
		return false
	}

	if p.DurationInDays != nil && (input.DurationInDays == nil || *p.DurationInDays != *input.DurationInDays) {
		return false
	}

	if p.JustificationRequiredOnApproval != nil && (input.JustificationRequiredOnApproval == nil || *p.JustificationRequiredOnApproval != *input.JustificationRequiredOnApproval) {
		return false
	}

	if p.MailNotificationsEnabled != nil && (input.MailNotificationsEnabled == nil || *p.MailNotificationsEnabled != *input.MailNotificationsEnabled) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RemindersEnabled != nil && (input.RemindersEnabled == nil || *p.RemindersEnabled != *input.RemindersEnabled) {
		return false
	}

	return true
}
