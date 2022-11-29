package entities

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EntityTimelineItem = BookmarkTimelineItem{}

type BookmarkTimelineItem struct {
	AzureResourceId string    `json:"azureResourceId"`
	CreatedBy       *UserInfo `json:"createdBy,omitempty"`
	DisplayName     *string   `json:"displayName,omitempty"`
	EndTimeUtc      *string   `json:"endTimeUtc,omitempty"`
	EventTime       *string   `json:"eventTime,omitempty"`
	Labels          *[]string `json:"labels,omitempty"`
	Notes           *string   `json:"notes,omitempty"`
	StartTimeUtc    *string   `json:"startTimeUtc,omitempty"`

	// Fields inherited from EntityTimelineItem
}

var _ json.Marshaler = BookmarkTimelineItem{}

func (s BookmarkTimelineItem) MarshalJSON() ([]byte, error) {
	type wrapper BookmarkTimelineItem
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BookmarkTimelineItem: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BookmarkTimelineItem: %+v", err)
	}
	decoded["kind"] = "Bookmark"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BookmarkTimelineItem: %+v", err)
	}

	return encoded, nil
}
