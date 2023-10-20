package featuresetversion

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FeatureWindow struct {
	FeatureWindowEnd   *string `json:"featureWindowEnd,omitempty"`
	FeatureWindowStart *string `json:"featureWindowStart,omitempty"`
}

func (o *FeatureWindow) GetFeatureWindowEndAsTime() (*time.Time, error) {
	if o.FeatureWindowEnd == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.FeatureWindowEnd, "2006-01-02T15:04:05Z07:00")
}

func (o *FeatureWindow) SetFeatureWindowEndAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.FeatureWindowEnd = &formatted
}

func (o *FeatureWindow) GetFeatureWindowStartAsTime() (*time.Time, error) {
	if o.FeatureWindowStart == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.FeatureWindowStart, "2006-01-02T15:04:05Z07:00")
}

func (o *FeatureWindow) SetFeatureWindowStartAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.FeatureWindowStart = &formatted
}
