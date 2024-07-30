package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewScheduleSettings struct {
	ApplyActions                         *[]AccessReviewApplyAction                  `json:"applyActions,omitempty"`
	AutoApplyDecisionsEnabled            *bool                                       `json:"autoApplyDecisionsEnabled,omitempty"`
	DecisionHistoriesForReviewersEnabled *bool                                       `json:"decisionHistoriesForReviewersEnabled,omitempty"`
	DefaultDecision                      *string                                     `json:"defaultDecision,omitempty"`
	DefaultDecisionEnabled               *bool                                       `json:"defaultDecisionEnabled,omitempty"`
	InstanceDurationInDays               *int64                                      `json:"instanceDurationInDays,omitempty"`
	JustificationRequiredOnApproval      *bool                                       `json:"justificationRequiredOnApproval,omitempty"`
	MailNotificationsEnabled             *bool                                       `json:"mailNotificationsEnabled,omitempty"`
	ODataType                            *string                                     `json:"@odata.type,omitempty"`
	RecommendationInsightSettings        *[]AccessReviewRecommendationInsightSetting `json:"recommendationInsightSettings,omitempty"`
	RecommendationLookBackDuration       *string                                     `json:"recommendationLookBackDuration,omitempty"`
	RecommendationsEnabled               *bool                                       `json:"recommendationsEnabled,omitempty"`
	Recurrence                           *PatternedRecurrence                        `json:"recurrence,omitempty"`
	ReminderNotificationsEnabled         *bool                                       `json:"reminderNotificationsEnabled,omitempty"`
}
