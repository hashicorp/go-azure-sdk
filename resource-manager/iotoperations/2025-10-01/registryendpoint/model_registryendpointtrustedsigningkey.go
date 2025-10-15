package registryendpoint

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegistryEndpointTrustedSigningKey interface {
	RegistryEndpointTrustedSigningKey() BaseRegistryEndpointTrustedSigningKeyImpl
}

var _ RegistryEndpointTrustedSigningKey = BaseRegistryEndpointTrustedSigningKeyImpl{}

type BaseRegistryEndpointTrustedSigningKeyImpl struct {
	Type RegistryEndpointTrustedSigningKeyType `json:"type"`
}

func (s BaseRegistryEndpointTrustedSigningKeyImpl) RegistryEndpointTrustedSigningKey() BaseRegistryEndpointTrustedSigningKeyImpl {
	return s
}

var _ RegistryEndpointTrustedSigningKey = RawRegistryEndpointTrustedSigningKeyImpl{}

// RawRegistryEndpointTrustedSigningKeyImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawRegistryEndpointTrustedSigningKeyImpl struct {
	registryEndpointTrustedSigningKey BaseRegistryEndpointTrustedSigningKeyImpl
	Type                              string
	Values                            map[string]interface{}
}

func (s RawRegistryEndpointTrustedSigningKeyImpl) RegistryEndpointTrustedSigningKey() BaseRegistryEndpointTrustedSigningKeyImpl {
	return s.registryEndpointTrustedSigningKey
}

func UnmarshalRegistryEndpointTrustedSigningKeyImplementation(input []byte) (RegistryEndpointTrustedSigningKey, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling RegistryEndpointTrustedSigningKey into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "ConfigMap") {
		var out RegistryEndpointTrustedSigningKeyConfigMap
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RegistryEndpointTrustedSigningKeyConfigMap: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Secret") {
		var out RegistryEndpointTrustedSigningKeySecret
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RegistryEndpointTrustedSigningKeySecret: %+v", err)
		}
		return out, nil
	}

	var parent BaseRegistryEndpointTrustedSigningKeyImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseRegistryEndpointTrustedSigningKeyImpl: %+v", err)
	}

	return RawRegistryEndpointTrustedSigningKeyImpl{
		registryEndpointTrustedSigningKey: parent,
		Type:                              value,
		Values:                            temp,
	}, nil

}
