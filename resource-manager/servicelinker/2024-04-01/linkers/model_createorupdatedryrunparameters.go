package linkers

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DryrunParameters = CreateOrUpdateDryrunParameters{}

type CreateOrUpdateDryrunParameters struct {
	AuthInfo              AuthInfoBase           `json:"authInfo"`
	ClientType            *ClientType            `json:"clientType,omitempty"`
	ConfigurationInfo     *ConfigurationInfo     `json:"configurationInfo,omitempty"`
	ProvisioningState     *string                `json:"provisioningState,omitempty"`
	PublicNetworkSolution *PublicNetworkSolution `json:"publicNetworkSolution,omitempty"`
	Scope                 *string                `json:"scope,omitempty"`
	SecretStore           *SecretStore           `json:"secretStore,omitempty"`
	TargetService         TargetServiceBase      `json:"targetService"`
	VNetSolution          *VNetSolution          `json:"vNetSolution,omitempty"`

	// Fields inherited from DryrunParameters

	ActionName DryrunActionName `json:"actionName"`
}

func (s CreateOrUpdateDryrunParameters) DryrunParameters() BaseDryrunParametersImpl {
	return BaseDryrunParametersImpl{
		ActionName: s.ActionName,
	}
}

var _ json.Marshaler = CreateOrUpdateDryrunParameters{}

func (s CreateOrUpdateDryrunParameters) MarshalJSON() ([]byte, error) {
	type wrapper CreateOrUpdateDryrunParameters
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CreateOrUpdateDryrunParameters: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CreateOrUpdateDryrunParameters: %+v", err)
	}

	decoded["actionName"] = "createOrUpdate"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CreateOrUpdateDryrunParameters: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &CreateOrUpdateDryrunParameters{}

func (s *CreateOrUpdateDryrunParameters) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ClientType            *ClientType            `json:"clientType,omitempty"`
		ConfigurationInfo     *ConfigurationInfo     `json:"configurationInfo,omitempty"`
		ProvisioningState     *string                `json:"provisioningState,omitempty"`
		PublicNetworkSolution *PublicNetworkSolution `json:"publicNetworkSolution,omitempty"`
		Scope                 *string                `json:"scope,omitempty"`
		SecretStore           *SecretStore           `json:"secretStore,omitempty"`
		VNetSolution          *VNetSolution          `json:"vNetSolution,omitempty"`
		ActionName            DryrunActionName       `json:"actionName"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ClientType = decoded.ClientType
	s.ConfigurationInfo = decoded.ConfigurationInfo
	s.ProvisioningState = decoded.ProvisioningState
	s.PublicNetworkSolution = decoded.PublicNetworkSolution
	s.Scope = decoded.Scope
	s.SecretStore = decoded.SecretStore
	s.VNetSolution = decoded.VNetSolution
	s.ActionName = decoded.ActionName

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling CreateOrUpdateDryrunParameters into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["authInfo"]; ok {
		impl, err := UnmarshalAuthInfoBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'AuthInfo' for 'CreateOrUpdateDryrunParameters': %+v", err)
		}
		s.AuthInfo = impl
	}

	if v, ok := temp["targetService"]; ok {
		impl, err := UnmarshalTargetServiceBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'TargetService' for 'CreateOrUpdateDryrunParameters': %+v", err)
		}
		s.TargetService = impl
	}

	return nil
}
