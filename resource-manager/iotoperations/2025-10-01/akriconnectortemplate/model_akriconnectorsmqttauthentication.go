package akriconnectortemplate

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AkriConnectorsMqttAuthentication interface {
	AkriConnectorsMqttAuthentication() BaseAkriConnectorsMqttAuthenticationImpl
}

var _ AkriConnectorsMqttAuthentication = BaseAkriConnectorsMqttAuthenticationImpl{}

type BaseAkriConnectorsMqttAuthenticationImpl struct {
	Method AkriConnectorsMqttAuthenticationMethod `json:"method"`
}

func (s BaseAkriConnectorsMqttAuthenticationImpl) AkriConnectorsMqttAuthentication() BaseAkriConnectorsMqttAuthenticationImpl {
	return s
}

var _ AkriConnectorsMqttAuthentication = RawAkriConnectorsMqttAuthenticationImpl{}

// RawAkriConnectorsMqttAuthenticationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAkriConnectorsMqttAuthenticationImpl struct {
	akriConnectorsMqttAuthentication BaseAkriConnectorsMqttAuthenticationImpl
	Type                             string
	Values                           map[string]interface{}
}

func (s RawAkriConnectorsMqttAuthenticationImpl) AkriConnectorsMqttAuthentication() BaseAkriConnectorsMqttAuthenticationImpl {
	return s.akriConnectorsMqttAuthentication
}

func UnmarshalAkriConnectorsMqttAuthenticationImplementation(input []byte) (AkriConnectorsMqttAuthentication, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AkriConnectorsMqttAuthentication into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["method"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "ServiceAccountToken") {
		var out AkriConnectorsServiceAccountAuthentication
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AkriConnectorsServiceAccountAuthentication: %+v", err)
		}
		return out, nil
	}

	var parent BaseAkriConnectorsMqttAuthenticationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAkriConnectorsMqttAuthenticationImpl: %+v", err)
	}

	return RawAkriConnectorsMqttAuthenticationImpl{
		akriConnectorsMqttAuthentication: parent,
		Type:                             value,
		Values:                           temp,
	}, nil

}
