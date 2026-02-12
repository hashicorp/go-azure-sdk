package iotcentrals

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ GroupAttestation = GroupX509Attestation{}

type GroupX509Attestation struct {
	X509 *SigningX509 `json:"x509,omitempty"`

	// Fields inherited from GroupAttestation

	Type string `json:"type"`
}

func (s GroupX509Attestation) GroupAttestation() BaseGroupAttestationImpl {
	return BaseGroupAttestationImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = GroupX509Attestation{}

func (s GroupX509Attestation) MarshalJSON() ([]byte, error) {
	type wrapper GroupX509Attestation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GroupX509Attestation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GroupX509Attestation: %+v", err)
	}

	decoded["type"] = "x509"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GroupX509Attestation: %+v", err)
	}

	return encoded, nil
}
