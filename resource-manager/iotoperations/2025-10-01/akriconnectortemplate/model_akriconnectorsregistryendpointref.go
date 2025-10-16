package akriconnectortemplate

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AkriConnectorsRegistrySettings = AkriConnectorsRegistryEndpointRef{}

type AkriConnectorsRegistryEndpointRef struct {
	RegistryEndpointRef string `json:"registryEndpointRef"`

	// Fields inherited from AkriConnectorsRegistrySettings

	RegistrySettingsType AkriConnectorsRegistrySettingsType `json:"registrySettingsType"`
}

func (s AkriConnectorsRegistryEndpointRef) AkriConnectorsRegistrySettings() BaseAkriConnectorsRegistrySettingsImpl {
	return BaseAkriConnectorsRegistrySettingsImpl{
		RegistrySettingsType: s.RegistrySettingsType,
	}
}

var _ json.Marshaler = AkriConnectorsRegistryEndpointRef{}

func (s AkriConnectorsRegistryEndpointRef) MarshalJSON() ([]byte, error) {
	type wrapper AkriConnectorsRegistryEndpointRef
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AkriConnectorsRegistryEndpointRef: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AkriConnectorsRegistryEndpointRef: %+v", err)
	}

	decoded["registrySettingsType"] = "RegistryEndpointRef"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AkriConnectorsRegistryEndpointRef: %+v", err)
	}

	return encoded, nil
}
