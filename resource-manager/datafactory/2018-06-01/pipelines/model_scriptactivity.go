package pipelines

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Activity = ScriptActivity{}

type ScriptActivity struct {
	LinkedServiceName *LinkedServiceReference      `json:"linkedServiceName,omitempty"`
	Policy            *ActivityPolicy              `json:"policy,omitempty"`
	TypeProperties    ScriptActivityTypeProperties `json:"typeProperties"`

	// Fields inherited from Activity
	DependsOn        *[]ActivityDependency     `json:"dependsOn,omitempty"`
	Description      *string                   `json:"description,omitempty"`
	Name             string                    `json:"name"`
	OnInactiveMarkAs *ActivityOnInactiveMarkAs `json:"onInactiveMarkAs,omitempty"`
	State            *ActivityState            `json:"state,omitempty"`
	UserProperties   *[]UserProperty           `json:"userProperties,omitempty"`
}

var _ json.Marshaler = ScriptActivity{}

func (s ScriptActivity) MarshalJSON() ([]byte, error) {
	type wrapper ScriptActivity
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ScriptActivity: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ScriptActivity: %+v", err)
	}
	decoded["type"] = "Script"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ScriptActivity: %+v", err)
	}

	return encoded, nil
}
