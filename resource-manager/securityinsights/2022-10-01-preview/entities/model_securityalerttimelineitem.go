package entities

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EntityTimelineItem = SecurityAlertTimelineItem{}

type SecurityAlertTimelineItem struct {
	AlertType       string           `json:"alertType"`
	AzureResourceId string           `json:"azureResourceId"`
	Description     *string          `json:"description,omitempty"`
	DisplayName     string           `json:"displayName"`
	EndTimeUtc      string           `json:"endTimeUtc"`
	Intent          *KillChainIntent `json:"intent,omitempty"`
	ProductName     *string          `json:"productName,omitempty"`
	Severity        AlertSeverity    `json:"severity"`
	StartTimeUtc    string           `json:"startTimeUtc"`
	Techniques      *[]string        `json:"techniques,omitempty"`
	TimeGenerated   string           `json:"timeGenerated"`

	// Fields inherited from EntityTimelineItem
}

var _ json.Marshaler = SecurityAlertTimelineItem{}

func (s SecurityAlertTimelineItem) MarshalJSON() ([]byte, error) {
	type wrapper SecurityAlertTimelineItem
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityAlertTimelineItem: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityAlertTimelineItem: %+v", err)
	}
	decoded["kind"] = "SecurityAlert"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityAlertTimelineItem: %+v", err)
	}

	return encoded, nil
}
