package recommendations

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendationPatch struct {
	HideUntilTimeUtc *string `json:"hideUntilTimeUtc,omitempty"`
	State            *State  `json:"state,omitempty"`
}

func (o *RecommendationPatch) GetHideUntilTimeUtcAsTime() (*time.Time, error) {
	if o.HideUntilTimeUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.HideUntilTimeUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *RecommendationPatch) SetHideUntilTimeUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.HideUntilTimeUtc = &formatted
}
