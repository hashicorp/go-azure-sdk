package replicationprotectionclusters

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharedDiskReplicationItemProperties struct {
	ActiveLocation                    *string                                       `json:"activeLocation,omitempty"`
	AllowedOperations                 *[]string                                     `json:"allowedOperations,omitempty"`
	CurrentScenario                   *CurrentScenarioDetails                       `json:"currentScenario,omitempty"`
	HealthErrors                      *[]HealthError                                `json:"healthErrors,omitempty"`
	ProtectionState                   *string                                       `json:"protectionState,omitempty"`
	ReplicationHealth                 *string                                       `json:"replicationHealth,omitempty"`
	SharedDiskProviderSpecificDetails SharedDiskReplicationProviderSpecificSettings `json:"sharedDiskProviderSpecificDetails"`
	TestFailoverState                 *string                                       `json:"testFailoverState,omitempty"`
}

var _ json.Unmarshaler = &SharedDiskReplicationItemProperties{}

func (s *SharedDiskReplicationItemProperties) UnmarshalJSON(bytes []byte) error {
	type alias SharedDiskReplicationItemProperties
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into SharedDiskReplicationItemProperties: %+v", err)
	}

	s.ActiveLocation = decoded.ActiveLocation
	s.AllowedOperations = decoded.AllowedOperations
	s.CurrentScenario = decoded.CurrentScenario
	s.HealthErrors = decoded.HealthErrors
	s.ProtectionState = decoded.ProtectionState
	s.ReplicationHealth = decoded.ReplicationHealth
	s.TestFailoverState = decoded.TestFailoverState

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SharedDiskReplicationItemProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["sharedDiskProviderSpecificDetails"]; ok {
		impl, err := unmarshalSharedDiskReplicationProviderSpecificSettingsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'SharedDiskProviderSpecificDetails' for 'SharedDiskReplicationItemProperties': %+v", err)
		}
		s.SharedDiskProviderSpecificDetails = impl
	}
	return nil
}
