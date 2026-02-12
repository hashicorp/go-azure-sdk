package iotcentrals

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ JobData = DeviceManifestMigrationJobData{}

type DeviceManifestMigrationJobData struct {
	Manifest string `json:"manifest"`

	// Fields inherited from JobData

	Type string `json:"type"`
}

func (s DeviceManifestMigrationJobData) JobData() BaseJobDataImpl {
	return BaseJobDataImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = DeviceManifestMigrationJobData{}

func (s DeviceManifestMigrationJobData) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManifestMigrationJobData
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManifestMigrationJobData: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManifestMigrationJobData: %+v", err)
	}

	decoded["type"] = "deviceManifestMigration"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManifestMigrationJobData: %+v", err)
	}

	return encoded, nil
}
