package registryendpoint

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ RegistryEndpointAuthentication = RegistryEndpointUserAssignedIdentityAuthentication{}

type RegistryEndpointUserAssignedIdentityAuthentication struct {
	UserAssignedManagedIdentitySettings RegistryEndpointUserAssignedManagedIdentitySettings `json:"userAssignedManagedIdentitySettings"`

	// Fields inherited from RegistryEndpointAuthentication

	Method RegistryEndpointAuthenticationMethod `json:"method"`
}

func (s RegistryEndpointUserAssignedIdentityAuthentication) RegistryEndpointAuthentication() BaseRegistryEndpointAuthenticationImpl {
	return BaseRegistryEndpointAuthenticationImpl{
		Method: s.Method,
	}
}

var _ json.Marshaler = RegistryEndpointUserAssignedIdentityAuthentication{}

func (s RegistryEndpointUserAssignedIdentityAuthentication) MarshalJSON() ([]byte, error) {
	type wrapper RegistryEndpointUserAssignedIdentityAuthentication
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling RegistryEndpointUserAssignedIdentityAuthentication: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling RegistryEndpointUserAssignedIdentityAuthentication: %+v", err)
	}

	decoded["method"] = "UserAssignedManagedIdentity"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling RegistryEndpointUserAssignedIdentityAuthentication: %+v", err)
	}

	return encoded, nil
}
