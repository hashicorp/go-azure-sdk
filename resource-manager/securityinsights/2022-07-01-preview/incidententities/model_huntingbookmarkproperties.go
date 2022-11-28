package incidententities

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HuntingBookmarkProperties struct {
	AdditionalData *map[string]interface{} `json:"additionalData,omitempty"`
	Created        *string                 `json:"created,omitempty"`
	CreatedBy      *UserInfo               `json:"createdBy"`
	DisplayName    string                  `json:"displayName"`
	EventTime      *string                 `json:"eventTime,omitempty"`
	FriendlyName   *string                 `json:"friendlyName,omitempty"`
	IncidentInfo   *IncidentInfo           `json:"incidentInfo"`
	Labels         *[]string               `json:"labels,omitempty"`
	Notes          *string                 `json:"notes,omitempty"`
	Query          string                  `json:"query"`
	QueryResult    *string                 `json:"queryResult,omitempty"`
	Updated        *string                 `json:"updated,omitempty"`
	UpdatedBy      *UserInfo               `json:"updatedBy"`
}

func (o *HuntingBookmarkProperties) GetCreatedAsTime() (*time.Time, error) {
	if o.Created == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.Created, "2006-01-02T15:04:05Z07:00")
}

func (o *HuntingBookmarkProperties) SetCreatedAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.Created = &formatted
}

func (o *HuntingBookmarkProperties) GetEventTimeAsTime() (*time.Time, error) {
	if o.EventTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EventTime, "2006-01-02T15:04:05Z07:00")
}

func (o *HuntingBookmarkProperties) SetEventTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EventTime = &formatted
}

func (o *HuntingBookmarkProperties) GetUpdatedAsTime() (*time.Time, error) {
	if o.Updated == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.Updated, "2006-01-02T15:04:05Z07:00")
}

func (o *HuntingBookmarkProperties) SetUpdatedAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.Updated = &formatted
}
