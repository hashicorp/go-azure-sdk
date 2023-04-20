package prediction

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PredictionResponseProperties struct {
	Category           *Category         `json:"category,omitempty"`
	ExtendedProperties *interface{}      `json:"extendedProperties,omitempty"`
	Impact             *Impact           `json:"impact,omitempty"`
	ImpactedField      *string           `json:"impactedField,omitempty"`
	LastUpdated        *string           `json:"lastUpdated,omitempty"`
	PredictionType     *PredictionType   `json:"predictionType,omitempty"`
	ShortDescription   *ShortDescription `json:"shortDescription,omitempty"`
}

func (o *PredictionResponseProperties) GetLastUpdatedAsTime() (*time.Time, error) {
	if o.LastUpdated == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastUpdated, "2006-01-02T15:04:05Z07:00")
}

func (o *PredictionResponseProperties) SetLastUpdatedAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastUpdated = &formatted
}
