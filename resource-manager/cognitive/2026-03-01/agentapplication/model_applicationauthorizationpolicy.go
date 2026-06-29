package agentapplication

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationAuthorizationPolicy interface {
	ApplicationAuthorizationPolicy() BaseApplicationAuthorizationPolicyImpl
}

var _ ApplicationAuthorizationPolicy = BaseApplicationAuthorizationPolicyImpl{}

type BaseApplicationAuthorizationPolicyImpl struct {
	Type BuiltInAuthorizationScheme `json:"type"`
}

func (s BaseApplicationAuthorizationPolicyImpl) ApplicationAuthorizationPolicy() BaseApplicationAuthorizationPolicyImpl {
	return s
}

var _ ApplicationAuthorizationPolicy = RawApplicationAuthorizationPolicyImpl{}

// RawApplicationAuthorizationPolicyImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawApplicationAuthorizationPolicyImpl struct {
	applicationAuthorizationPolicy BaseApplicationAuthorizationPolicyImpl
	Type                           string
	Values                         map[string]interface{}
}

func (s RawApplicationAuthorizationPolicyImpl) ApplicationAuthorizationPolicy() BaseApplicationAuthorizationPolicyImpl {
	return s.applicationAuthorizationPolicy
}

func (s RawApplicationAuthorizationPolicyImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalApplicationAuthorizationPolicyImplementation(input []byte) (ApplicationAuthorizationPolicy, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ApplicationAuthorizationPolicy into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "Channels") {
		var out ChannelsBuiltInAuthorizationPolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ChannelsBuiltInAuthorizationPolicy: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "OrganizationScope") {
		var out OrganizationSharedBuiltInAuthorizationPolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OrganizationSharedBuiltInAuthorizationPolicy: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Default") {
		var out RoleBasedBuiltInAuthorizationPolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RoleBasedBuiltInAuthorizationPolicy: %+v", err)
		}
		return out, nil
	}

	var parent BaseApplicationAuthorizationPolicyImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseApplicationAuthorizationPolicyImpl: %+v", err)
	}

	return RawApplicationAuthorizationPolicyImpl{
		applicationAuthorizationPolicy: parent,
		Type:                           value,
		Values:                         temp,
	}, nil

}
