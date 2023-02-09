package sharesubscription

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SourceShareSynchronizationSetting = ScheduledSourceSynchronizationSetting{}

type ScheduledSourceSynchronizationSetting struct {
	Properties *ScheduledSourceShareSynchronizationSettingProperties `json:"properties,omitempty"`

	// Fields inherited from SourceShareSynchronizationSetting
}

var _ json.Marshaler = ScheduledSourceSynchronizationSetting{}

func (s ScheduledSourceSynchronizationSetting) MarshalJSON() ([]byte, error) {
	type wrapper ScheduledSourceSynchronizationSetting
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ScheduledSourceSynchronizationSetting: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ScheduledSourceSynchronizationSetting: %+v", err)
	}
	decoded["kind"] = "ScheduleBased"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ScheduledSourceSynchronizationSetting: %+v", err)
	}

	return encoded, nil
}
