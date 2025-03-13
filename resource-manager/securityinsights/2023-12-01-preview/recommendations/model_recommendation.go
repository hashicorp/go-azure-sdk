package recommendations

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Recommendation struct {
	Actions                 []RecommendedAction `json:"actions"`
	AdditionalProperties    *map[string]string  `json:"additionalProperties,omitempty"`
	Category                Category            `json:"category"`
	Content                 *Content            `json:"content,omitempty"`
	Context                 Context             `json:"context"`
	Description             string              `json:"description"`
	DisplayUntilTimeUtc     *string             `json:"displayUntilTimeUtc,omitempty"`
	HideUntilTimeUtc        *string             `json:"hideUntilTimeUtc,omitempty"`
	Id                      string              `json:"id"`
	Instructions            Instructions        `json:"instructions"`
	LastEvaluatedTimeUtc    string              `json:"lastEvaluatedTimeUtc"`
	Priority                Priority            `json:"priority"`
	RecommendationTypeId    string              `json:"recommendationTypeId"`
	RecommendationTypeTitle string              `json:"recommendationTypeTitle"`
	ResourceId              *string             `json:"resourceId,omitempty"`
	State                   State               `json:"state"`
	Title                   string              `json:"title"`
	Visible                 *bool               `json:"visible,omitempty"`
	WorkspaceId             string              `json:"workspaceId"`
}

func (o *Recommendation) GetDisplayUntilTimeUtcAsTime() (*time.Time, error) {
	if o.DisplayUntilTimeUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.DisplayUntilTimeUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *Recommendation) SetDisplayUntilTimeUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.DisplayUntilTimeUtc = &formatted
}

func (o *Recommendation) GetHideUntilTimeUtcAsTime() (*time.Time, error) {
	if o.HideUntilTimeUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.HideUntilTimeUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *Recommendation) SetHideUntilTimeUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.HideUntilTimeUtc = &formatted
}

func (o *Recommendation) GetLastEvaluatedTimeUtcAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.LastEvaluatedTimeUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *Recommendation) SetLastEvaluatedTimeUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastEvaluatedTimeUtc = formatted
}
