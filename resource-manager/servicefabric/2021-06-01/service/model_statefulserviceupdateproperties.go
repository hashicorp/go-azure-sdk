package service

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ServiceResourceUpdateProperties = StatefulServiceUpdateProperties{}

type StatefulServiceUpdateProperties struct {
	MinReplicaSetSize          *int64  `json:"minReplicaSetSize,omitempty"`
	QuorumLossWaitDuration     *string `json:"quorumLossWaitDuration,omitempty"`
	ReplicaRestartWaitDuration *string `json:"replicaRestartWaitDuration,omitempty"`
	StandByReplicaKeepDuration *string `json:"standByReplicaKeepDuration,omitempty"`
	TargetReplicaSetSize       *int64  `json:"targetReplicaSetSize,omitempty"`

	// Fields inherited from ServiceResourceUpdateProperties
	CorrelationScheme        *[]ServiceCorrelationDescription     `json:"correlationScheme,omitempty"`
	DefaultMoveCost          *MoveCost                            `json:"defaultMoveCost,omitempty"`
	PlacementConstraints     *string                              `json:"placementConstraints,omitempty"`
	ServiceLoadMetrics       *[]ServiceLoadMetricDescription      `json:"serviceLoadMetrics,omitempty"`
	ServicePlacementPolicies *[]ServicePlacementPolicyDescription `json:"servicePlacementPolicies,omitempty"`
}

var _ json.Marshaler = StatefulServiceUpdateProperties{}

func (s StatefulServiceUpdateProperties) MarshalJSON() ([]byte, error) {
	type wrapper StatefulServiceUpdateProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling StatefulServiceUpdateProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling StatefulServiceUpdateProperties: %+v", err)
	}
	decoded["serviceKind"] = "Stateful"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling StatefulServiceUpdateProperties: %+v", err)
	}

	return encoded, nil
}
