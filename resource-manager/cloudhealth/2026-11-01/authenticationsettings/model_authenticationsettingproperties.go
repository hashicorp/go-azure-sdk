package authenticationsettings

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationSettingProperties interface {
	AuthenticationSettingProperties() BaseAuthenticationSettingPropertiesImpl
}

var _ AuthenticationSettingProperties = BaseAuthenticationSettingPropertiesImpl{}

type BaseAuthenticationSettingPropertiesImpl struct {
	AuthenticationKind AuthenticationKind            `json:"authenticationKind"`
	DisplayName        *string                       `json:"displayName,omitempty"`
	ProvisioningState  *HealthModelProvisioningState `json:"provisioningState,omitempty"`
}

func (s BaseAuthenticationSettingPropertiesImpl) AuthenticationSettingProperties() BaseAuthenticationSettingPropertiesImpl {
	return s
}

var _ AuthenticationSettingProperties = RawAuthenticationSettingPropertiesImpl{}

// RawAuthenticationSettingPropertiesImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawAuthenticationSettingPropertiesImpl struct {
	authenticationSettingProperties BaseAuthenticationSettingPropertiesImpl
	Type                            string
	Values                          map[string]interface{}
}

func (s RawAuthenticationSettingPropertiesImpl) AuthenticationSettingProperties() BaseAuthenticationSettingPropertiesImpl {
	return s.authenticationSettingProperties
}

func (s RawAuthenticationSettingPropertiesImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalAuthenticationSettingPropertiesImplementation(input []byte) (AuthenticationSettingProperties, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AuthenticationSettingProperties into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["authenticationKind"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "ManagedIdentity") {
		var out ManagedIdentityAuthenticationSettingProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedIdentityAuthenticationSettingProperties: %+v", err)
		}
		return out, nil
	}

	var parent BaseAuthenticationSettingPropertiesImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAuthenticationSettingPropertiesImpl: %+v", err)
	}

	return RawAuthenticationSettingPropertiesImpl{
		authenticationSettingProperties: parent,
		Type:                            value,
		Values:                          temp,
	}, nil

}
