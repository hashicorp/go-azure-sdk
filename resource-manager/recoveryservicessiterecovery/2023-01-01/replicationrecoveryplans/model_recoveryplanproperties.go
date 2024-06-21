package replicationrecoveryplans

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecoveryPlanProperties struct {
	AllowedOperations                *[]string                              `json:"allowedOperations,omitempty"`
	CurrentScenario                  *CurrentScenarioDetails                `json:"currentScenario,omitempty"`
	CurrentScenarioStatus            *string                                `json:"currentScenarioStatus,omitempty"`
	CurrentScenarioStatusDescription *string                                `json:"currentScenarioStatusDescription,omitempty"`
	FailoverDeploymentModel          *string                                `json:"failoverDeploymentModel,omitempty"`
	FriendlyName                     *string                                `json:"friendlyName,omitempty"`
	Groups                           *[]RecoveryPlanGroup                   `json:"groups,omitempty"`
	LastPlannedFailoverTime          *string                                `json:"lastPlannedFailoverTime,omitempty"`
	LastTestFailoverTime             *string                                `json:"lastTestFailoverTime,omitempty"`
	LastUnplannedFailoverTime        *string                                `json:"lastUnplannedFailoverTime,omitempty"`
	PrimaryFabricFriendlyName        *string                                `json:"primaryFabricFriendlyName,omitempty"`
	PrimaryFabricId                  *string                                `json:"primaryFabricId,omitempty"`
	ProviderSpecificDetails          *[]RecoveryPlanProviderSpecificDetails `json:"providerSpecificDetails,omitempty"`
	RecoveryFabricFriendlyName       *string                                `json:"recoveryFabricFriendlyName,omitempty"`
	RecoveryFabricId                 *string                                `json:"recoveryFabricId,omitempty"`
	ReplicationProviders             *[]string                              `json:"replicationProviders,omitempty"`
}

var _ json.Unmarshaler = &RecoveryPlanProperties{}

func (s *RecoveryPlanProperties) UnmarshalJSON(bytes []byte) error {
	type alias RecoveryPlanProperties
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into RecoveryPlanProperties: %+v", err)
	}

	s.AllowedOperations = decoded.AllowedOperations
	s.CurrentScenario = decoded.CurrentScenario
	s.CurrentScenarioStatus = decoded.CurrentScenarioStatus
	s.CurrentScenarioStatusDescription = decoded.CurrentScenarioStatusDescription
	s.FailoverDeploymentModel = decoded.FailoverDeploymentModel
	s.FriendlyName = decoded.FriendlyName
	s.Groups = decoded.Groups
	s.LastPlannedFailoverTime = decoded.LastPlannedFailoverTime
	s.LastTestFailoverTime = decoded.LastTestFailoverTime
	s.LastUnplannedFailoverTime = decoded.LastUnplannedFailoverTime
	s.PrimaryFabricFriendlyName = decoded.PrimaryFabricFriendlyName
	s.PrimaryFabricId = decoded.PrimaryFabricId
	s.RecoveryFabricFriendlyName = decoded.RecoveryFabricFriendlyName
	s.RecoveryFabricId = decoded.RecoveryFabricId
	s.ReplicationProviders = decoded.ReplicationProviders

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling RecoveryPlanProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["providerSpecificDetails"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ProviderSpecificDetails into list []json.RawMessage: %+v", err)
		}

		output := make([]RecoveryPlanProviderSpecificDetails, 0)
		for i, val := range listTemp {
			impl, err := unmarshalRecoveryPlanProviderSpecificDetailsImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ProviderSpecificDetails' for 'RecoveryPlanProperties': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ProviderSpecificDetails = &output
	}
	return nil
}
