package entities

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntityTimelineItem interface {
	EntityTimelineItem() BaseEntityTimelineItemImpl
}

var _ EntityTimelineItem = BaseEntityTimelineItemImpl{}

type BaseEntityTimelineItemImpl struct {
	Kind EntityTimelineKind `json:"kind"`
}

func (s BaseEntityTimelineItemImpl) EntityTimelineItem() BaseEntityTimelineItemImpl {
	return s
}

var _ EntityTimelineItem = RawEntityTimelineItemImpl{}

// RawEntityTimelineItemImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawEntityTimelineItemImpl struct {
	entityTimelineItem BaseEntityTimelineItemImpl
	Type               string
	Values             map[string]interface{}
}

func (s RawEntityTimelineItemImpl) EntityTimelineItem() BaseEntityTimelineItemImpl {
	return s.entityTimelineItem
}

func (s RawEntityTimelineItemImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalEntityTimelineItemImplementation(input []byte) (EntityTimelineItem, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling EntityTimelineItem into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["kind"]; ok {
		value = fmt.Sprintf("%v", v)
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

	var parent BaseEntityTimelineItemImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseEntityTimelineItemImpl: %+v", err)
	}

	return RawEntityTimelineItemImpl{
		entityTimelineItem: parent,
		Type:               value,
		Values:             temp,
	}, nil

}
