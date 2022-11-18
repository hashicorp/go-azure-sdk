package service

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ServiceResourceProperties = StatefulServiceProperties{}

type StatefulServiceProperties struct {
	HasPersistedState          *bool   `json:"hasPersistedState,omitempty"`
	MinReplicaSetSize          *int64  `json:"minReplicaSetSize,omitempty"`
	QuorumLossWaitDuration     *string `json:"quorumLossWaitDuration,omitempty"`
	ReplicaRestartWaitDuration *string `json:"replicaRestartWaitDuration,omitempty"`
	ServicePlacementTimeLimit  *string `json:"servicePlacementTimeLimit,omitempty"`
	StandByReplicaKeepDuration *string `json:"standByReplicaKeepDuration,omitempty"`
	TargetReplicaSetSize       *int64  `json:"targetReplicaSetSize,omitempty"`

	// Fields inherited from ServiceResourceProperties
	CorrelationScheme            *[]ServiceCorrelation         `json:"correlationScheme,omitempty"`
	DefaultMoveCost              *MoveCost                     `json:"defaultMoveCost,omitempty"`
	PartitionDescription         Partition                     `json:"partitionDescription"`
	PlacementConstraints         *string                       `json:"placementConstraints,omitempty"`
	ProvisioningState            *string                       `json:"provisioningState,omitempty"`
	ScalingPolicies              *[]ScalingPolicy              `json:"scalingPolicies,omitempty"`
	ServiceLoadMetrics           *[]ServiceLoadMetric          `json:"serviceLoadMetrics,omitempty"`
	ServicePackageActivationMode *ServicePackageActivationMode `json:"servicePackageActivationMode,omitempty"`
	ServicePlacementPolicies     *[]ServicePlacementPolicy     `json:"servicePlacementPolicies,omitempty"`
	ServiceTypeName              string                        `json:"serviceTypeName"`
}

var _ json.Marshaler = StatefulServiceProperties{}

func (s StatefulServiceProperties) MarshalJSON() ([]byte, error) {
	type wrapper StatefulServiceProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling StatefulServiceProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling StatefulServiceProperties: %+v", err)
	}
	decoded["serviceKind"] = "Stateful"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling StatefulServiceProperties: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &StatefulServiceProperties{}

func (s *StatefulServiceProperties) UnmarshalJSON(bytes []byte) error {
	type alias StatefulServiceProperties
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into StatefulServiceProperties: %+v", err)
	}

	s.CorrelationScheme = decoded.CorrelationScheme
	s.DefaultMoveCost = decoded.DefaultMoveCost
	s.HasPersistedState = decoded.HasPersistedState
	s.MinReplicaSetSize = decoded.MinReplicaSetSize
	s.PlacementConstraints = decoded.PlacementConstraints
	s.ProvisioningState = decoded.ProvisioningState
	s.QuorumLossWaitDuration = decoded.QuorumLossWaitDuration
	s.ReplicaRestartWaitDuration = decoded.ReplicaRestartWaitDuration
	s.ScalingPolicies = decoded.ScalingPolicies
	s.ServiceLoadMetrics = decoded.ServiceLoadMetrics
	s.ServicePackageActivationMode = decoded.ServicePackageActivationMode
	s.ServicePlacementTimeLimit = decoded.ServicePlacementTimeLimit
	s.ServiceTypeName = decoded.ServiceTypeName
	s.StandByReplicaKeepDuration = decoded.StandByReplicaKeepDuration
	s.TargetReplicaSetSize = decoded.TargetReplicaSetSize

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling StatefulServiceProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["partitionDescription"]; ok {
		impl, err := unmarshalPartitionImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'PartitionDescription' for 'StatefulServiceProperties': %+v", err)
		}
		s.PartitionDescription = impl
	}

	if v, ok := temp["servicePlacementPolicies"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ServicePlacementPolicies into list []json.RawMessage: %+v", err)
		}

		output := make([]ServicePlacementPolicy, 0)
		for i, val := range listTemp {
			impl, err := unmarshalServicePlacementPolicyImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ServicePlacementPolicies' for 'StatefulServiceProperties': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ServicePlacementPolicies = &output
	}
	return nil
}
