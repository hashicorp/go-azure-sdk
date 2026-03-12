package triggers

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ TriggerProperties = SourceImageTriggerProperties{}

type SourceImageTriggerProperties struct {

	// Fields inherited from TriggerProperties

	Kind              string             `json:"kind"`
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty"`
	Status            *TriggerStatus     `json:"status,omitempty"`
}

func (s SourceImageTriggerProperties) TriggerProperties() BaseTriggerPropertiesImpl {
	return BaseTriggerPropertiesImpl{
		Kind:              s.Kind,
		ProvisioningState: s.ProvisioningState,
		Status:            s.Status,
	}
}

var _ json.Marshaler = SourceImageTriggerProperties{}

func (s SourceImageTriggerProperties) MarshalJSON() ([]byte, error) {
	type wrapper SourceImageTriggerProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SourceImageTriggerProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SourceImageTriggerProperties: %+v", err)
	}

	decoded["kind"] = "SourceImage"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SourceImageTriggerProperties: %+v", err)
	}

	return encoded, nil
}
