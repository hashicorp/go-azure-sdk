package registryendpoint

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ RegistryEndpointAuthentication = RegistryEndpointSystemAssignedIdentityAuthentication{}

type RegistryEndpointSystemAssignedIdentityAuthentication struct {
	SystemAssignedManagedIdentitySettings RegistryEndpointSystemAssignedManagedIdentitySettings `json:"systemAssignedManagedIdentitySettings"`

	// Fields inherited from RegistryEndpointAuthentication

	Method RegistryEndpointAuthenticationMethod `json:"method"`
}

func (s RegistryEndpointSystemAssignedIdentityAuthentication) RegistryEndpointAuthentication() BaseRegistryEndpointAuthenticationImpl {
	return BaseRegistryEndpointAuthenticationImpl{
		Method: s.Method,
	}
}

var _ json.Marshaler = RegistryEndpointSystemAssignedIdentityAuthentication{}

func (s RegistryEndpointSystemAssignedIdentityAuthentication) MarshalJSON() ([]byte, error) {
	type wrapper RegistryEndpointSystemAssignedIdentityAuthentication
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling RegistryEndpointSystemAssignedIdentityAuthentication: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling RegistryEndpointSystemAssignedIdentityAuthentication: %+v", err)
	}

	decoded["method"] = "SystemAssignedManagedIdentity"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling RegistryEndpointSystemAssignedIdentityAuthentication: %+v", err)
	}

	return encoded, nil
}
