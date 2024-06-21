package incidentbookmarks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HuntingBookmarkProperties struct {
	AdditionalData *map[string]interface{} `json:"additionalData,omitempty"`
	Created        *string                 `json:"created,omitempty"`
	CreatedBy      *UserInfo               `json:"createdBy,omitempty"`
	DisplayName    string                  `json:"displayName"`
	EventTime      *string                 `json:"eventTime,omitempty"`
	FriendlyName   *string                 `json:"friendlyName,omitempty"`
	IncidentInfo   *IncidentInfo           `json:"incidentInfo,omitempty"`
	Labels         *[]string               `json:"labels,omitempty"`
	Notes          *string                 `json:"notes,omitempty"`
	Query          string                  `json:"query"`
	QueryResult    *string                 `json:"queryResult,omitempty"`
	Updated        *string                 `json:"updated,omitempty"`
	UpdatedBy      *UserInfo               `json:"updatedBy,omitempty"`
}
