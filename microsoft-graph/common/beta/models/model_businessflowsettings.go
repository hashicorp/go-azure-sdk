package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BusinessFlowSettings struct {
	AccessRecommendationsEnabled    *bool                           `json:"accessRecommendationsEnabled,omitempty"`
	ActivityDurationInDays          *int64                          `json:"activityDurationInDays,omitempty"`
	AutoApplyReviewResultsEnabled   *bool                           `json:"autoApplyReviewResultsEnabled,omitempty"`
	AutoReviewEnabled               *bool                           `json:"autoReviewEnabled,omitempty"`
	AutoReviewSettings              *AutoReviewSettings             `json:"autoReviewSettings,omitempty"`
	DurationInDays                  *int64                          `json:"durationInDays,omitempty"`
	JustificationRequiredOnApproval *bool                           `json:"justificationRequiredOnApproval,omitempty"`
	MailNotificationsEnabled        *bool                           `json:"mailNotificationsEnabled,omitempty"`
	ODataType                       *string                         `json:"@odata.type,omitempty"`
	RecurrenceSettings              *AccessReviewRecurrenceSettings `json:"recurrenceSettings,omitempty"`
	RemindersEnabled                *bool                           `json:"remindersEnabled,omitempty"`
}
