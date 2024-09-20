package entities

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EntityTimelineItem = AnomalyTimelineItem{}

type AnomalyTimelineItem struct {
	AzureResourceId string    `json:"azureResourceId"`
	Description     *string   `json:"description,omitempty"`
	DisplayName     string    `json:"displayName"`
	EndTimeUtc      string    `json:"endTimeUtc"`
	Intent          *string   `json:"intent,omitempty"`
	ProductName     *string   `json:"productName,omitempty"`
	Reasons         *[]string `json:"reasons,omitempty"`
	StartTimeUtc    string    `json:"startTimeUtc"`
	Techniques      *[]string `json:"techniques,omitempty"`
	TimeGenerated   string    `json:"timeGenerated"`
	Vendor          *string   `json:"vendor,omitempty"`

	// Fields inherited from EntityTimelineItem

	Kind EntityTimelineKind `json:"kind"`
}

func (s AnomalyTimelineItem) EntityTimelineItem() BaseEntityTimelineItemImpl {
	return BaseEntityTimelineItemImpl{
		Kind: s.Kind,
	}
}

var _ json.Marshaler = AnomalyTimelineItem{}

func (s AnomalyTimelineItem) MarshalJSON() ([]byte, error) {
	type wrapper AnomalyTimelineItem
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AnomalyTimelineItem: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AnomalyTimelineItem: %+v", err)
	}

	decoded["kind"] = "Anomaly"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AnomalyTimelineItem: %+v", err)
	}

	return encoded, nil
}
