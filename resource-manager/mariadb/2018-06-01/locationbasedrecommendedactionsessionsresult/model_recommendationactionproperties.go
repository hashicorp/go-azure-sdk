package locationbasedrecommendedactionsessionsresult

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendationActionProperties struct {
	ActionId           *int64             `json:"actionId,omitempty"`
	AdvisorName        *string            `json:"advisorName,omitempty"`
	CreatedTime        *string            `json:"createdTime,omitempty"`
	Details            *map[string]string `json:"details,omitempty"`
	ExpirationTime     *string            `json:"expirationTime,omitempty"`
	Reason             *string            `json:"reason,omitempty"`
	RecommendationType *string            `json:"recommendationType,omitempty"`
	SessionId          *string            `json:"sessionId,omitempty"`
}

func (o *RecommendationActionProperties) GetCreatedTimeAsTime() (*time.Time, error) {
	if o.CreatedTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreatedTime, "2006-01-02T15:04:05Z07:00")
}

func (o *RecommendationActionProperties) SetCreatedTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CreatedTime = &formatted
}

func (o *RecommendationActionProperties) GetExpirationTimeAsTime() (*time.Time, error) {
	if o.ExpirationTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ExpirationTime, "2006-01-02T15:04:05Z07:00")
}

func (o *RecommendationActionProperties) SetExpirationTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ExpirationTime = &formatted
}
