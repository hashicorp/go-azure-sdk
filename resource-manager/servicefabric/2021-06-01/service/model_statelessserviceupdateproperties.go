package service

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ServiceResourceUpdateProperties = StatelessServiceUpdateProperties{}

type StatelessServiceUpdateProperties struct {
	InstanceCloseDelayDuration *string `json:"instanceCloseDelayDuration,omitempty"`
	InstanceCount              *int64  `json:"instanceCount,omitempty"`

	// Fields inherited from ServiceResourceUpdateProperties

	CorrelationScheme        *[]ServiceCorrelationDescription     `json:"correlationScheme,omitempty"`
	DefaultMoveCost          *MoveCost                            `json:"defaultMoveCost,omitempty"`
	PlacementConstraints     *string                              `json:"placementConstraints,omitempty"`
	ServiceKind              ServiceKind                          `json:"serviceKind"`
	ServiceLoadMetrics       *[]ServiceLoadMetricDescription      `json:"serviceLoadMetrics,omitempty"`
	ServicePlacementPolicies *[]ServicePlacementPolicyDescription `json:"servicePlacementPolicies,omitempty"`
}

func (s StatelessServiceUpdateProperties) ServiceResourceUpdateProperties() BaseServiceResourceUpdatePropertiesImpl {
	return BaseServiceResourceUpdatePropertiesImpl{
		CorrelationScheme:        s.CorrelationScheme,
		DefaultMoveCost:          s.DefaultMoveCost,
		PlacementConstraints:     s.PlacementConstraints,
		ServiceKind:              s.ServiceKind,
		ServiceLoadMetrics:       s.ServiceLoadMetrics,
		ServicePlacementPolicies: s.ServicePlacementPolicies,
	}
}

var _ json.Marshaler = StatelessServiceUpdateProperties{}

func (s StatelessServiceUpdateProperties) MarshalJSON() ([]byte, error) {
	type wrapper StatelessServiceUpdateProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling StatelessServiceUpdateProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling StatelessServiceUpdateProperties: %+v", err)
	}

	decoded["serviceKind"] = "Stateless"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling StatelessServiceUpdateProperties: %+v", err)
	}

	return encoded, nil
}
