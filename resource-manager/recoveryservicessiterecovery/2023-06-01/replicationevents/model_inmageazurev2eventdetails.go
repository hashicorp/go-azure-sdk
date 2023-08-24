package replicationevents

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EventProviderSpecificDetails = InMageAzureV2EventDetails{}

type InMageAzureV2EventDetails struct {
	Category         *string `json:"category,omitempty"`
	Component        *string `json:"component,omitempty"`
	CorrectiveAction *string `json:"correctiveAction,omitempty"`
	Details          *string `json:"details,omitempty"`
	EventType        *string `json:"eventType,omitempty"`
	SiteName         *string `json:"siteName,omitempty"`
	Summary          *string `json:"summary,omitempty"`

	// Fields inherited from EventProviderSpecificDetails
}

var _ json.Marshaler = InMageAzureV2EventDetails{}

func (s InMageAzureV2EventDetails) MarshalJSON() ([]byte, error) {
	type wrapper InMageAzureV2EventDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling InMageAzureV2EventDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling InMageAzureV2EventDetails: %+v", err)
	}
	decoded["instanceType"] = "InMageAzureV2"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling InMageAzureV2EventDetails: %+v", err)
	}

	return encoded, nil
}
