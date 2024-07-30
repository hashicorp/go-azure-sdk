package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewStageSettings struct {
	DecisionsThatWillMoveToNextStage *[]string                                   `json:"decisionsThatWillMoveToNextStage,omitempty"`
	DependsOn                        *[]string                                   `json:"dependsOn,omitempty"`
	DurationInDays                   *int64                                      `json:"durationInDays,omitempty"`
	FallbackReviewers                *[]AccessReviewReviewerScope                `json:"fallbackReviewers,omitempty"`
	ODataType                        *string                                     `json:"@odata.type,omitempty"`
	RecommendationInsightSettings    *[]AccessReviewRecommendationInsightSetting `json:"recommendationInsightSettings,omitempty"`
	RecommendationsEnabled           *bool                                       `json:"recommendationsEnabled,omitempty"`
	Reviewers                        *[]AccessReviewReviewerScope                `json:"reviewers,omitempty"`
	StageId                          *string                                     `json:"stageId,omitempty"`
}
