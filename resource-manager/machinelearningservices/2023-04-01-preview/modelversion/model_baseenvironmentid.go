package modelversion

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ BaseEnvironmentSource = BaseEnvironmentId{}

type BaseEnvironmentId struct {
	ResourceId string `json:"resourceId"`

	// Fields inherited from BaseEnvironmentSource
}

var _ json.Marshaler = BaseEnvironmentId{}

func (s BaseEnvironmentId) MarshalJSON() ([]byte, error) {
	type wrapper BaseEnvironmentId
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseEnvironmentId: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseEnvironmentId: %+v", err)
	}
	decoded["baseEnvironmentSourceType"] = "EnvironmentAsset"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseEnvironmentId: %+v", err)
	}

	return encoded, nil
}
