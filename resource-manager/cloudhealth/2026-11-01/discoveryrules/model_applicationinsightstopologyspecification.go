package discoveryrules

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DiscoveryRuleSpecification = ApplicationInsightsTopologySpecification{}

type ApplicationInsightsTopologySpecification struct {
	ApplicationInsightsResourceId string `json:"applicationInsightsResourceId"`

	// Fields inherited from DiscoveryRuleSpecification

	Kind DiscoveryRuleKind `json:"kind"`
}

func (s ApplicationInsightsTopologySpecification) DiscoveryRuleSpecification() BaseDiscoveryRuleSpecificationImpl {
	return BaseDiscoveryRuleSpecificationImpl{
		Kind: s.Kind,
	}
}

var _ json.Marshaler = ApplicationInsightsTopologySpecification{}

func (s ApplicationInsightsTopologySpecification) MarshalJSON() ([]byte, error) {
	type wrapper ApplicationInsightsTopologySpecification
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ApplicationInsightsTopologySpecification: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ApplicationInsightsTopologySpecification: %+v", err)
	}

	decoded["kind"] = "ApplicationInsightsTopology"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ApplicationInsightsTopologySpecification: %+v", err)
	}

	return encoded, nil
}
