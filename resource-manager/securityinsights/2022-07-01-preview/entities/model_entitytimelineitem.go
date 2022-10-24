package entities

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntityTimelineItem interface {
}

func unmarshalEntityTimelineItemImplementation(input []byte) (EntityTimelineItem, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling EntityTimelineItem into map[string]interface: %+v", err)
	}

	value, ok := temp["kind"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Activity") {
		var out ActivityTimelineItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ActivityTimelineItem: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Anomaly") {
		var out AnomalyTimelineItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AnomalyTimelineItem: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Bookmark") {
		var out BookmarkTimelineItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BookmarkTimelineItem: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "SecurityAlert") {
		var out SecurityAlertTimelineItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityAlertTimelineItem: %+v", err)
		}
		return out, nil
	}

	type RawEntityTimelineItemImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawEntityTimelineItemImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
