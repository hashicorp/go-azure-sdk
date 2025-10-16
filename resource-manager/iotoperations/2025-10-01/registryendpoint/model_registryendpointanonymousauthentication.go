package registryendpoint

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ RegistryEndpointAuthentication = RegistryEndpointAnonymousAuthentication{}

type RegistryEndpointAnonymousAuthentication struct {
	AnonymousSettings interface{} `json:"anonymousSettings"`

	// Fields inherited from RegistryEndpointAuthentication

	Method RegistryEndpointAuthenticationMethod `json:"method"`
}

func (s RegistryEndpointAnonymousAuthentication) RegistryEndpointAuthentication() BaseRegistryEndpointAuthenticationImpl {
	return BaseRegistryEndpointAuthenticationImpl{
		Method: s.Method,
	}
}

var _ json.Marshaler = RegistryEndpointAnonymousAuthentication{}

func (s RegistryEndpointAnonymousAuthentication) MarshalJSON() ([]byte, error) {
	type wrapper RegistryEndpointAnonymousAuthentication
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling RegistryEndpointAnonymousAuthentication: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling RegistryEndpointAnonymousAuthentication: %+v", err)
	}

	decoded["method"] = "Anonymous"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling RegistryEndpointAnonymousAuthentication: %+v", err)
	}

	return encoded, nil
}
