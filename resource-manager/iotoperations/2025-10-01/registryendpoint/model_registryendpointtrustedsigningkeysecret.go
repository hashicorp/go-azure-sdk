package registryendpoint

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ RegistryEndpointTrustedSigningKey = RegistryEndpointTrustedSigningKeySecret{}

type RegistryEndpointTrustedSigningKeySecret struct {
	SecretRef string `json:"secretRef"`

	// Fields inherited from RegistryEndpointTrustedSigningKey

	Type RegistryEndpointTrustedSigningKeyType `json:"type"`
}

func (s RegistryEndpointTrustedSigningKeySecret) RegistryEndpointTrustedSigningKey() BaseRegistryEndpointTrustedSigningKeyImpl {
	return BaseRegistryEndpointTrustedSigningKeyImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = RegistryEndpointTrustedSigningKeySecret{}

func (s RegistryEndpointTrustedSigningKeySecret) MarshalJSON() ([]byte, error) {
	type wrapper RegistryEndpointTrustedSigningKeySecret
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling RegistryEndpointTrustedSigningKeySecret: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling RegistryEndpointTrustedSigningKeySecret: %+v", err)
	}

	decoded["type"] = "Secret"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling RegistryEndpointTrustedSigningKeySecret: %+v", err)
	}

	return encoded, nil
}
