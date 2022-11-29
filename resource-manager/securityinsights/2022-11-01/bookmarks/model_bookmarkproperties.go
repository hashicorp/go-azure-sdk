package bookmarks

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookmarkProperties struct {
	Created        *string       `json:"created,omitempty"`
	CreatedBy      *UserInfo     `json:"createdBy,omitempty"`
	DisplayName    string        `json:"displayName"`
	EventTime      *string       `json:"eventTime,omitempty"`
	IncidentInfo   *IncidentInfo `json:"incidentInfo,omitempty"`
	Labels         *[]string     `json:"labels,omitempty"`
	Notes          *string       `json:"notes,omitempty"`
	Query          string        `json:"query"`
	QueryEndTime   *string       `json:"queryEndTime,omitempty"`
	QueryResult    *string       `json:"queryResult,omitempty"`
	QueryStartTime *string       `json:"queryStartTime,omitempty"`
	Updated        *string       `json:"updated,omitempty"`
	UpdatedBy      *UserInfo     `json:"updatedBy,omitempty"`
}

func (o *BookmarkProperties) GetCreatedAsTime() (*time.Time, error) {
	if o.Created == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.Created, "2006-01-02T15:04:05Z07:00")
}

func (o *BookmarkProperties) SetCreatedAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.Created = &formatted
}

func (o *BookmarkProperties) GetEventTimeAsTime() (*time.Time, error) {
	if o.EventTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EventTime, "2006-01-02T15:04:05Z07:00")
}

func (o *BookmarkProperties) SetEventTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EventTime = &formatted
}

func (o *BookmarkProperties) GetQueryEndTimeAsTime() (*time.Time, error) {
	if o.QueryEndTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.QueryEndTime, "2006-01-02T15:04:05Z07:00")
}

func (o *BookmarkProperties) SetQueryEndTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.QueryEndTime = &formatted
}

func (o *BookmarkProperties) GetQueryStartTimeAsTime() (*time.Time, error) {
	if o.QueryStartTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.QueryStartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *BookmarkProperties) SetQueryStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.QueryStartTime = &formatted
}

func (o *BookmarkProperties) GetUpdatedAsTime() (*time.Time, error) {
	if o.Updated == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.Updated, "2006-01-02T15:04:05Z07:00")
}

func (o *BookmarkProperties) SetUpdatedAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.Updated = &formatted
}
