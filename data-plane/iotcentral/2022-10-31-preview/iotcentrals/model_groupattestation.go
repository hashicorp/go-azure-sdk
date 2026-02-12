package iotcentrals

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupAttestation interface {
	GroupAttestation() BaseGroupAttestationImpl
}

var _ GroupAttestation = BaseGroupAttestationImpl{}

type BaseGroupAttestationImpl struct {
	Type string `json:"type"`
}

func (s BaseGroupAttestationImpl) GroupAttestation() BaseGroupAttestationImpl {
	return s
}

var _ GroupAttestation = RawGroupAttestationImpl{}

// RawGroupAttestationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawGroupAttestationImpl struct {
	groupAttestation BaseGroupAttestationImpl
	Type             string
	Values           map[string]interface{}
}

func (s RawGroupAttestationImpl) GroupAttestation() BaseGroupAttestationImpl {
	return s.groupAttestation
}

func UnmarshalGroupAttestationImplementation(input []byte) (GroupAttestation, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling GroupAttestation into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "symmetricKey") {
		var out GroupSymmetricKeyAttestation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupSymmetricKeyAttestation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "x509") {
		var out GroupX509Attestation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupX509Attestation: %+v", err)
		}
		return out, nil
	}

	var parent BaseGroupAttestationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseGroupAttestationImpl: %+v", err)
	}

	return RawGroupAttestationImpl{
		groupAttestation: parent,
		Type:             value,
		Values:           temp,
	}, nil

}
