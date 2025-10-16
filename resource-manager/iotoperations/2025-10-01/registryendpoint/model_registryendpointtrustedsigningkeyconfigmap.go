package registryendpoint

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ RegistryEndpointTrustedSigningKey = RegistryEndpointTrustedSigningKeyConfigMap{}

type RegistryEndpointTrustedSigningKeyConfigMap struct {
	ConfigMapRef string `json:"configMapRef"`

	// Fields inherited from RegistryEndpointTrustedSigningKey

	Type RegistryEndpointTrustedSigningKeyType `json:"type"`
}

func (s RegistryEndpointTrustedSigningKeyConfigMap) RegistryEndpointTrustedSigningKey() BaseRegistryEndpointTrustedSigningKeyImpl {
	return BaseRegistryEndpointTrustedSigningKeyImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = RegistryEndpointTrustedSigningKeyConfigMap{}

func (s RegistryEndpointTrustedSigningKeyConfigMap) MarshalJSON() ([]byte, error) {
	type wrapper RegistryEndpointTrustedSigningKeyConfigMap
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling RegistryEndpointTrustedSigningKeyConfigMap: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling RegistryEndpointTrustedSigningKeyConfigMap: %+v", err)
	}

	decoded["type"] = "ConfigMap"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling RegistryEndpointTrustedSigningKeyConfigMap: %+v", err)
	}

	return encoded, nil
}
