package akriconnectortemplate

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AkriConnectorsRegistrySettings = AkriConnectorsContainerRegistry{}

type AkriConnectorsContainerRegistry struct {
	ContainerRegistrySettings AkriConnectorsContainerRegistrySettings `json:"containerRegistrySettings"`

	// Fields inherited from AkriConnectorsRegistrySettings

	RegistrySettingsType AkriConnectorsRegistrySettingsType `json:"registrySettingsType"`
}

func (s AkriConnectorsContainerRegistry) AkriConnectorsRegistrySettings() BaseAkriConnectorsRegistrySettingsImpl {
	return BaseAkriConnectorsRegistrySettingsImpl{
		RegistrySettingsType: s.RegistrySettingsType,
	}
}

var _ json.Marshaler = AkriConnectorsContainerRegistry{}

func (s AkriConnectorsContainerRegistry) MarshalJSON() ([]byte, error) {
	type wrapper AkriConnectorsContainerRegistry
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AkriConnectorsContainerRegistry: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AkriConnectorsContainerRegistry: %+v", err)
	}

	decoded["registrySettingsType"] = "ContainerRegistry"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AkriConnectorsContainerRegistry: %+v", err)
	}

	return encoded, nil
}
