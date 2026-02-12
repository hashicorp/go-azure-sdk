package iotcentrals

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Attestation = TpmAttestation{}

type TpmAttestation struct {
	Tpm Tpm `json:"tpm"`

	// Fields inherited from Attestation

	Type string `json:"type"`
}

func (s TpmAttestation) Attestation() BaseAttestationImpl {
	return BaseAttestationImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = TpmAttestation{}

func (s TpmAttestation) MarshalJSON() ([]byte, error) {
	type wrapper TpmAttestation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TpmAttestation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TpmAttestation: %+v", err)
	}

	decoded["type"] = "tpm"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TpmAttestation: %+v", err)
	}

	return encoded, nil
}
