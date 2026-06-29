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

// RawAkriConnectorsMqttAuthenticationImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawAkriConnectorsMqttAuthenticationImpl struct {
	akriConnectorsMqttAuthentication BaseAkriConnectorsMqttAuthenticationImpl
	Type                             string
	Values                           map[string]interface{}
}

func (s RawAkriConnectorsMqttAuthenticationImpl) AkriConnectorsMqttAuthentication() BaseAkriConnectorsMqttAuthenticationImpl {
	return s.akriConnectorsMqttAuthentication
}

func (s RawAkriConnectorsMqttAuthenticationImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
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
