package factories

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FactoryProperties struct {
	CreateTime           *string                                  `json:"createTime,omitempty"`
	Encryption           *EncryptionConfiguration                 `json:"encryption,omitempty"`
	GlobalParameters     *map[string]GlobalParameterSpecification `json:"globalParameters,omitempty"`
	ProvisioningState    *string                                  `json:"provisioningState,omitempty"`
	PublicNetworkAccess  *PublicNetworkAccess                     `json:"publicNetworkAccess,omitempty"`
	PurviewConfiguration *PurviewConfiguration                    `json:"purviewConfiguration,omitempty"`
	RepoConfiguration    FactoryRepoConfiguration                 `json:"repoConfiguration"`
	Version              *string                                  `json:"version,omitempty"`
}

var _ json.Unmarshaler = &FactoryProperties{}

func (s *FactoryProperties) UnmarshalJSON(bytes []byte) error {
	type alias FactoryProperties
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into FactoryProperties: %+v", err)
	}

	s.CreateTime = decoded.CreateTime
	s.Encryption = decoded.Encryption
	s.GlobalParameters = decoded.GlobalParameters
	s.ProvisioningState = decoded.ProvisioningState
	s.PublicNetworkAccess = decoded.PublicNetworkAccess
	s.PurviewConfiguration = decoded.PurviewConfiguration
	s.Version = decoded.Version

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling FactoryProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["repoConfiguration"]; ok {
		impl, err := unmarshalFactoryRepoConfigurationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'RepoConfiguration' for 'FactoryProperties': %+v", err)
		}
		s.RepoConfiguration = impl
	}
	return nil
}
