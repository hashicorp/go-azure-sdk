package registryendpoint

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ RegistryEndpointAuthentication = RegistryEndpointArtifactPullSecretAuthentication{}

type RegistryEndpointArtifactPullSecretAuthentication struct {
	ArtifactPullSecretSettings RegistryEndpointArtifactPullSecretSettings `json:"artifactPullSecretSettings"`

	// Fields inherited from RegistryEndpointAuthentication

	Method RegistryEndpointAuthenticationMethod `json:"method"`
}

func (s RegistryEndpointArtifactPullSecretAuthentication) RegistryEndpointAuthentication() BaseRegistryEndpointAuthenticationImpl {
	return BaseRegistryEndpointAuthenticationImpl{
		Method: s.Method,
	}
}

var _ json.Marshaler = RegistryEndpointArtifactPullSecretAuthentication{}

func (s RegistryEndpointArtifactPullSecretAuthentication) MarshalJSON() ([]byte, error) {
	type wrapper RegistryEndpointArtifactPullSecretAuthentication
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling RegistryEndpointArtifactPullSecretAuthentication: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling RegistryEndpointArtifactPullSecretAuthentication: %+v", err)
	}

	decoded["method"] = "ArtifactPullSecret"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling RegistryEndpointArtifactPullSecretAuthentication: %+v", err)
	}

	return encoded, nil
}
