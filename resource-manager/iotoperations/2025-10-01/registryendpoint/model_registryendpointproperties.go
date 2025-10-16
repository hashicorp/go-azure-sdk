package registryendpoint

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegistryEndpointProperties struct {
	Authentication    RegistryEndpointAuthentication       `json:"authentication"`
	CodeSigningCas    *[]RegistryEndpointTrustedSigningKey `json:"codeSigningCas,omitempty"`
	HealthState       *ResourceHealthState                 `json:"healthState,omitempty"`
	Host              string                               `json:"host"`
	ProvisioningState *ProvisioningState                   `json:"provisioningState,omitempty"`
}

var _ json.Unmarshaler = &RegistryEndpointProperties{}

func (s *RegistryEndpointProperties) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		HealthState       *ResourceHealthState `json:"healthState,omitempty"`
		Host              string               `json:"host"`
		ProvisioningState *ProvisioningState   `json:"provisioningState,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.HealthState = decoded.HealthState
	s.Host = decoded.Host
	s.ProvisioningState = decoded.ProvisioningState

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling RegistryEndpointProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["authentication"]; ok {
		impl, err := UnmarshalRegistryEndpointAuthenticationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Authentication' for 'RegistryEndpointProperties': %+v", err)
		}
		s.Authentication = impl
	}

	if v, ok := temp["codeSigningCas"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling CodeSigningCas into list []json.RawMessage: %+v", err)
		}

		output := make([]RegistryEndpointTrustedSigningKey, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalRegistryEndpointTrustedSigningKeyImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'CodeSigningCas' for 'RegistryEndpointProperties': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.CodeSigningCas = &output
	}

	return nil
}
