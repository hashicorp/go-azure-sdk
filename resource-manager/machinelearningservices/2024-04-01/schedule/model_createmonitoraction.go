package schedule

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ScheduleActionBase = CreateMonitorAction{}

type CreateMonitorAction struct {
	MonitorDefinition MonitorDefinition `json:"monitorDefinition"`

	// Fields inherited from ScheduleActionBase
}

var _ json.Marshaler = CreateMonitorAction{}

func (s CreateMonitorAction) MarshalJSON() ([]byte, error) {
	type wrapper CreateMonitorAction
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CreateMonitorAction: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CreateMonitorAction: %+v", err)
	}
	decoded["actionType"] = "CreateMonitor"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CreateMonitorAction: %+v", err)
	}

	return encoded, nil
}
