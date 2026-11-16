package discoveryrules

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiscoveryRuleSpecification interface {
	DiscoveryRuleSpecification() BaseDiscoveryRuleSpecificationImpl
}

var _ DiscoveryRuleSpecification = BaseDiscoveryRuleSpecificationImpl{}

type BaseDiscoveryRuleSpecificationImpl struct {
	Kind DiscoveryRuleKind `json:"kind"`
}

func (s BaseDiscoveryRuleSpecificationImpl) DiscoveryRuleSpecification() BaseDiscoveryRuleSpecificationImpl {
	return s
}

var _ DiscoveryRuleSpecification = RawDiscoveryRuleSpecificationImpl{}

// RawDiscoveryRuleSpecificationImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawDiscoveryRuleSpecificationImpl struct {
	discoveryRuleSpecification BaseDiscoveryRuleSpecificationImpl
	Type                       string
	Values                     map[string]interface{}
}

func (s RawDiscoveryRuleSpecificationImpl) DiscoveryRuleSpecification() BaseDiscoveryRuleSpecificationImpl {
	return s.discoveryRuleSpecification
}

func (s RawDiscoveryRuleSpecificationImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalDiscoveryRuleSpecificationImplementation(input []byte) (DiscoveryRuleSpecification, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DiscoveryRuleSpecification into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["kind"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "ApplicationInsightsTopology") {
		var out ApplicationInsightsTopologySpecification
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ApplicationInsightsTopologySpecification: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "ResourceGraphQuery") {
		var out ResourceGraphQuerySpecification
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ResourceGraphQuerySpecification: %+v", err)
		}
		return out, nil
	}

	var parent BaseDiscoveryRuleSpecificationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDiscoveryRuleSpecificationImpl: %+v", err)
	}

	return RawDiscoveryRuleSpecificationImpl{
		discoveryRuleSpecification: parent,
		Type:                       value,
		Values:                     temp,
	}, nil

}
