package registryendpoint

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegistryEndpointAuthentication interface {
	RegistryEndpointAuthentication() BaseRegistryEndpointAuthenticationImpl
}

var _ RegistryEndpointAuthentication = BaseRegistryEndpointAuthenticationImpl{}

type BaseRegistryEndpointAuthenticationImpl struct {
	Method RegistryEndpointAuthenticationMethod `json:"method"`
}

func (s BaseRegistryEndpointAuthenticationImpl) RegistryEndpointAuthentication() BaseRegistryEndpointAuthenticationImpl {
	return s
}

var _ RegistryEndpointAuthentication = RawRegistryEndpointAuthenticationImpl{}

// RawRegistryEndpointAuthenticationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawRegistryEndpointAuthenticationImpl struct {
	registryEndpointAuthentication BaseRegistryEndpointAuthenticationImpl
	Type                           string
	Values                         map[string]interface{}
}

func (s RawRegistryEndpointAuthenticationImpl) RegistryEndpointAuthentication() BaseRegistryEndpointAuthenticationImpl {
	return s.registryEndpointAuthentication
}

func UnmarshalRegistryEndpointAuthenticationImplementation(input []byte) (RegistryEndpointAuthentication, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling RegistryEndpointAuthentication into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["method"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "Anonymous") {
		var out RegistryEndpointAnonymousAuthentication
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RegistryEndpointAnonymousAuthentication: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "ArtifactPullSecret") {
		var out RegistryEndpointArtifactPullSecretAuthentication
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RegistryEndpointArtifactPullSecretAuthentication: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "SystemAssignedManagedIdentity") {
		var out RegistryEndpointSystemAssignedIdentityAuthentication
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RegistryEndpointSystemAssignedIdentityAuthentication: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "UserAssignedManagedIdentity") {
		var out RegistryEndpointUserAssignedIdentityAuthentication
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RegistryEndpointUserAssignedIdentityAuthentication: %+v", err)
		}
		return out, nil
	}

	var parent BaseRegistryEndpointAuthenticationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseRegistryEndpointAuthenticationImpl: %+v", err)
	}

	return RawRegistryEndpointAuthenticationImpl{
		registryEndpointAuthentication: parent,
		Type:                           value,
		Values:                         temp,
	}, nil

}
