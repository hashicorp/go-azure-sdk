package hosts

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ HostProperties = SpecializedHostProperties{}

type SpecializedHostProperties struct {

	// Fields inherited from HostProperties

	DisplayName       *string                `json:"displayName,omitempty"`
	FaultDomain       *string                `json:"faultDomain,omitempty"`
	Fqdn              *string                `json:"fqdn,omitempty"`
	Kind              HostKind               `json:"kind"`
	Maintenance       *HostMaintenance       `json:"maintenance,omitempty"`
	MoRefId           *string                `json:"moRefId,omitempty"`
	ProvisioningState *HostProvisioningState `json:"provisioningState,omitempty"`
}

func (s SpecializedHostProperties) HostProperties() BaseHostPropertiesImpl {
	return BaseHostPropertiesImpl{
		DisplayName:       s.DisplayName,
		FaultDomain:       s.FaultDomain,
		Fqdn:              s.Fqdn,
		Kind:              s.Kind,
		Maintenance:       s.Maintenance,
		MoRefId:           s.MoRefId,
		ProvisioningState: s.ProvisioningState,
	}
}

var _ json.Marshaler = SpecializedHostProperties{}

func (s SpecializedHostProperties) MarshalJSON() ([]byte, error) {
	type wrapper SpecializedHostProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SpecializedHostProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SpecializedHostProperties: %+v", err)
	}

	decoded["kind"] = "Specialized"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SpecializedHostProperties: %+v", err)
	}

	return encoded, nil
}
