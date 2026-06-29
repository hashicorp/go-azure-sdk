package iotcentrals

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Attestation interface {
	Attestation() BaseAttestationImpl
}

var _ Attestation = BaseAttestationImpl{}

type BaseAttestationImpl struct {
	Type string `json:"type"`
}

func (s BaseAttestationImpl) Attestation() BaseAttestationImpl {
	return s
}

var _ Attestation = RawAttestationImpl{}

// RawAttestationImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawAttestationImpl struct {
	attestation BaseAttestationImpl
	Type        string
	Values      map[string]interface{}
}

func (s RawAttestationImpl) Attestation() BaseAttestationImpl {
	return s.attestation
}

func (s RawAttestationImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalAttestationImplementation(input []byte) (Attestation, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling Attestation into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "symmetricKey") {
		var out SymmetricKeyAttestation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SymmetricKeyAttestation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "tpm") {
		var out TpmAttestation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TpmAttestation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "x509") {
		var out X509Attestation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into X509Attestation: %+v", err)
		}
		return out, nil
	}

	var parent BaseAttestationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAttestationImpl: %+v", err)
	}

	return RawAttestationImpl{
		attestation: parent,
		Type:        value,
		Values:      temp,
	}, nil

}
