package entities

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EntityTimelineItem = ActivityTimelineItem{}

type ActivityTimelineItem struct {
	BucketEndTimeUTC     string `json:"bucketEndTimeUTC"`
	BucketStartTimeUTC   string `json:"bucketStartTimeUTC"`
	Content              string `json:"content"`
	FirstActivityTimeUTC string `json:"firstActivityTimeUTC"`
	LastActivityTimeUTC  string `json:"lastActivityTimeUTC"`
	QueryId              string `json:"queryId"`
	Title                string `json:"title"`

	// Fields inherited from EntityTimelineItem

	Kind EntityTimelineKind `json:"kind"`
}

func (s ActivityTimelineItem) EntityTimelineItem() BaseEntityTimelineItemImpl {
	return BaseEntityTimelineItemImpl{
		Kind: s.Kind,
	}
}

var _ json.Marshaler = ActivityTimelineItem{}

func (s ActivityTimelineItem) MarshalJSON() ([]byte, error) {
	type wrapper ActivityTimelineItem
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ActivityTimelineItem: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ActivityTimelineItem: %+v", err)
	}

	decoded["kind"] = "Activity"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ActivityTimelineItem: %+v", err)
	}

	return encoded, nil
}
