package iotcentrals

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Attestation = SymmetricKeyAttestation{}

type SymmetricKeyAttestation struct {
	SymmetricKey SymmetricKey `json:"symmetricKey"`

	// Fields inherited from Attestation

	Type string `json:"type"`
}

func (s SymmetricKeyAttestation) Attestation() BaseAttestationImpl {
	return BaseAttestationImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = SymmetricKeyAttestation{}

func (s SymmetricKeyAttestation) MarshalJSON() ([]byte, error) {
	type wrapper SymmetricKeyAttestation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SymmetricKeyAttestation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SymmetricKeyAttestation: %+v", err)
	}

	decoded["type"] = "symmetricKey"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SymmetricKeyAttestation: %+v", err)
	}

	return encoded, nil
}
