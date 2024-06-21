package bookmarks

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
