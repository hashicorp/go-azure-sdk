package discoveryrules

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiscoveryRuleProperties struct {
	AddRecommendedSignals   DiscoveryRuleRecommendedSignalsBehavior        `json:"addRecommendedSignals"`
	AddResourceHealthSignal *ResourceHealthAvailabilityStateSignalBehavior `json:"addResourceHealthSignal,omitempty"`
	AuthenticationSetting   string                                         `json:"authenticationSetting"`
	DiscoverRelationships   DiscoveryRuleRelationshipDiscoveryBehavior     `json:"discoverRelationships"`
	DisplayName             *string                                        `json:"displayName,omitempty"`
	EntityName              *string                                        `json:"entityName,omitempty"`
	Error                   *DiscoveryError                                `json:"error,omitempty"`
	ProvisioningState       *HealthModelProvisioningState                  `json:"provisioningState,omitempty"`
	Specification           DiscoveryRuleSpecification                     `json:"specification"`
}

var _ json.Unmarshaler = &DiscoveryRuleProperties{}

func (s *DiscoveryRuleProperties) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AddRecommendedSignals   DiscoveryRuleRecommendedSignalsBehavior        `json:"addRecommendedSignals"`
		AddResourceHealthSignal *ResourceHealthAvailabilityStateSignalBehavior `json:"addResourceHealthSignal,omitempty"`
		AuthenticationSetting   string                                         `json:"authenticationSetting"`
		DiscoverRelationships   DiscoveryRuleRelationshipDiscoveryBehavior     `json:"discoverRelationships"`
		DisplayName             *string                                        `json:"displayName,omitempty"`
		EntityName              *string                                        `json:"entityName,omitempty"`
		Error                   *DiscoveryError                                `json:"error,omitempty"`
		ProvisioningState       *HealthModelProvisioningState                  `json:"provisioningState,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AddRecommendedSignals = decoded.AddRecommendedSignals
	s.AddResourceHealthSignal = decoded.AddResourceHealthSignal
	s.AuthenticationSetting = decoded.AuthenticationSetting
	s.DiscoverRelationships = decoded.DiscoverRelationships
	s.DisplayName = decoded.DisplayName
	s.EntityName = decoded.EntityName
	s.Error = decoded.Error
	s.ProvisioningState = decoded.ProvisioningState

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DiscoveryRuleProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["specification"]; ok {
		impl, err := UnmarshalDiscoveryRuleSpecificationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Specification' for 'DiscoveryRuleProperties': %+v", err)
		}
		s.Specification = impl
	}

	return nil
}
