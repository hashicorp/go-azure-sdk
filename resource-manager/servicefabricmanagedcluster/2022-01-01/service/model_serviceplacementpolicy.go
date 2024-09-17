package service

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePlacementPolicy interface {
	ServicePlacementPolicy() BaseServicePlacementPolicyImpl
}

var _ ServicePlacementPolicy = BaseServicePlacementPolicyImpl{}

type BaseServicePlacementPolicyImpl struct {
	Type ServicePlacementPolicyType `json:"type"`
}

func (s BaseServicePlacementPolicyImpl) ServicePlacementPolicy() BaseServicePlacementPolicyImpl {
	return s
}

var _ ServicePlacementPolicy = RawServicePlacementPolicyImpl{}

// RawServicePlacementPolicyImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawServicePlacementPolicyImpl struct {
	servicePlacementPolicy BaseServicePlacementPolicyImpl
	Type                   string
	Values                 map[string]interface{}
}

func (s RawServicePlacementPolicyImpl) ServicePlacementPolicy() BaseServicePlacementPolicyImpl {
	return s.servicePlacementPolicy
}

func UnmarshalServicePlacementPolicyImplementation(input []byte) (ServicePlacementPolicy, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ServicePlacementPolicy into map[string]interface: %+v", err)
	}

	value, ok := temp["type"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "InvalidDomain") {
		var out ServicePlacementInvalidDomainPolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ServicePlacementInvalidDomainPolicy: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "NonPartiallyPlaceService") {
		var out ServicePlacementNonPartiallyPlaceServicePolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ServicePlacementNonPartiallyPlaceServicePolicy: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "PreferredPrimaryDomain") {
		var out ServicePlacementPreferPrimaryDomainPolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ServicePlacementPreferPrimaryDomainPolicy: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "RequiredDomainDistribution") {
		var out ServicePlacementRequireDomainDistributionPolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ServicePlacementRequireDomainDistributionPolicy: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "RequiredDomain") {
		var out ServicePlacementRequiredDomainPolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ServicePlacementRequiredDomainPolicy: %+v", err)
		}
		return out, nil
	}

	var parent BaseServicePlacementPolicyImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseServicePlacementPolicyImpl: %+v", err)
	}

	return RawServicePlacementPolicyImpl{
		servicePlacementPolicy: parent,
		Type:                   value,
		Values:                 temp,
	}, nil

}
