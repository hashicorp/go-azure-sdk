package appplatform

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ JobTriggerConfig = ManualJobTriggerConfig{}

type ManualJobTriggerConfig struct {
	Parallelism      *int64 `json:"parallelism,omitempty"`
	RetryLimit       *int64 `json:"retryLimit,omitempty"`
	TimeoutInSeconds *int64 `json:"timeoutInSeconds,omitempty"`

	// Fields inherited from JobTriggerConfig

	TriggerType TriggerType `json:"triggerType"`
}

func (s ManualJobTriggerConfig) JobTriggerConfig() BaseJobTriggerConfigImpl {
	return BaseJobTriggerConfigImpl{
		TriggerType: s.TriggerType,
	}
}

var _ json.Marshaler = ManualJobTriggerConfig{}

func (s ManualJobTriggerConfig) MarshalJSON() ([]byte, error) {
	type wrapper ManualJobTriggerConfig
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManualJobTriggerConfig: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManualJobTriggerConfig: %+v", err)
	}

	decoded["triggerType"] = "Manual"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManualJobTriggerConfig: %+v", err)
	}

	return encoded, nil
}
