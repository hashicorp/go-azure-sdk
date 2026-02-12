package iotcentrals

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ JobData = DeviceTemplateMigrationJobData{}

type DeviceTemplateMigrationJobData struct {
	Template string `json:"template"`

	// Fields inherited from JobData

	Type string `json:"type"`
}

func (s DeviceTemplateMigrationJobData) JobData() BaseJobDataImpl {
	return BaseJobDataImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = DeviceTemplateMigrationJobData{}

func (s DeviceTemplateMigrationJobData) MarshalJSON() ([]byte, error) {
	type wrapper DeviceTemplateMigrationJobData
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceTemplateMigrationJobData: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceTemplateMigrationJobData: %+v", err)
	}

	decoded["type"] = "deviceTemplateMigration"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceTemplateMigrationJobData: %+v", err)
	}

	return encoded, nil
}
