package pipelines

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Activity = WebActivity{}

type WebActivity struct {
	LinkedServiceName *LinkedServiceReference   `json:"linkedServiceName,omitempty"`
	Policy            *ActivityPolicy           `json:"policy,omitempty"`
	TypeProperties    WebActivityTypeProperties `json:"typeProperties"`

	// Fields inherited from Activity
	DependsOn        *[]ActivityDependency     `json:"dependsOn,omitempty"`
	Description      *string                   `json:"description,omitempty"`
	Name             string                    `json:"name"`
	OnInactiveMarkAs *ActivityOnInactiveMarkAs `json:"onInactiveMarkAs,omitempty"`
	State            *ActivityState            `json:"state,omitempty"`
	UserProperties   *[]UserProperty           `json:"userProperties,omitempty"`
}

var _ json.Marshaler = WebActivity{}

func (s WebActivity) MarshalJSON() ([]byte, error) {
	type wrapper WebActivity
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WebActivity: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WebActivity: %+v", err)
	}
	decoded["type"] = "WebActivity"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WebActivity: %+v", err)
	}

	return encoded, nil
}
