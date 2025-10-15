package tuningoptions

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ObjectRecommendationProperties struct {
	AnalyzedWorkload       *ObjectRecommendationPropertiesAnalyzedWorkload      `json:"analyzedWorkload,omitempty"`
	CurrentState           *string                                              `json:"currentState,omitempty"`
	Details                *ObjectRecommendationDetails                         `json:"details,omitempty"`
	EstimatedImpact        *[]ImpactRecord                                      `json:"estimatedImpact,omitempty"`
	ImplementationDetails  *ObjectRecommendationPropertiesImplementationDetails `json:"implementationDetails,omitempty"`
	ImprovedQueryIds       *[]int64                                             `json:"improvedQueryIds,omitempty"`
	InitialRecommendedTime *string                                              `json:"initialRecommendedTime,omitempty"`
	LastRecommendedTime    *string                                              `json:"lastRecommendedTime,omitempty"`
	RecommendationReason   *string                                              `json:"recommendationReason,omitempty"`
	RecommendationType     *RecommendationType                                  `json:"recommendationType,omitempty"`
	TimesRecommended       *int64                                               `json:"timesRecommended,omitempty"`
}

func (o *ObjectRecommendationProperties) GetInitialRecommendedTimeAsTime() (*time.Time, error) {
	if o.InitialRecommendedTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.InitialRecommendedTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ObjectRecommendationProperties) SetInitialRecommendedTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.InitialRecommendedTime = &formatted
}

func (o *ObjectRecommendationProperties) GetLastRecommendedTimeAsTime() (*time.Time, error) {
	if o.LastRecommendedTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastRecommendedTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ObjectRecommendationProperties) SetLastRecommendedTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastRecommendedTime = &formatted
}
