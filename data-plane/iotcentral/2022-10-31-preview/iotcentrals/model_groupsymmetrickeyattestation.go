package iotcentrals

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ GroupAttestation = GroupSymmetricKeyAttestation{}

type GroupSymmetricKeyAttestation struct {
	SymmetricKey *SymmetricKey `json:"symmetricKey,omitempty"`

	// Fields inherited from GroupAttestation

	Type string `json:"type"`
}

func (s GroupSymmetricKeyAttestation) GroupAttestation() BaseGroupAttestationImpl {
	return BaseGroupAttestationImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = GroupSymmetricKeyAttestation{}

func (s GroupSymmetricKeyAttestation) MarshalJSON() ([]byte, error) {
	type wrapper GroupSymmetricKeyAttestation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GroupSymmetricKeyAttestation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GroupSymmetricKeyAttestation: %+v", err)
	}

	decoded["type"] = "symmetricKey"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GroupSymmetricKeyAttestation: %+v", err)
	}

	return encoded, nil
}
