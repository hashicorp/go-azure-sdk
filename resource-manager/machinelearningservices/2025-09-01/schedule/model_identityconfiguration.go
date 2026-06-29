package schedule

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityConfiguration interface {
	IdentityConfiguration() BaseIdentityConfigurationImpl
}

var _ IdentityConfiguration = BaseIdentityConfigurationImpl{}

type BaseIdentityConfigurationImpl struct {
	IdentityType IdentityConfigurationType `json:"identityType"`
}

func (s BaseIdentityConfigurationImpl) IdentityConfiguration() BaseIdentityConfigurationImpl {
	return s
}

var _ IdentityConfiguration = RawIdentityConfigurationImpl{}

// RawIdentityConfigurationImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawIdentityConfigurationImpl struct {
	identityConfiguration BaseIdentityConfigurationImpl
	Type                  string
	Values                map[string]interface{}
}

func (s RawIdentityConfigurationImpl) IdentityConfiguration() BaseIdentityConfigurationImpl {
	return s.identityConfiguration
}

func (s RawIdentityConfigurationImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalIdentityConfigurationImplementation(input []byte) (IdentityConfiguration, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling IdentityConfiguration into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["identityType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "AMLToken") {
		var out AmlToken
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AmlToken: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Managed") {
		var out ManagedIdentity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedIdentity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "UserIdentity") {
		var out UserIdentity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserIdentity: %+v", err)
		}
		return out, nil
	}

	var parent BaseIdentityConfigurationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseIdentityConfigurationImpl: %+v", err)
	}

	return RawIdentityConfigurationImpl{
		identityConfiguration: parent,
		Type:                  value,
		Values:                temp,
	}, nil

}
