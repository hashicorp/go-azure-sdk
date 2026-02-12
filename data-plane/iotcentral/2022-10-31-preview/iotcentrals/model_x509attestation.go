package iotcentrals

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Attestation = X509Attestation{}

type X509Attestation struct {
	X509 X509 `json:"x509"`

	// Fields inherited from Attestation

	Type string `json:"type"`
}

func (s X509Attestation) Attestation() BaseAttestationImpl {
	return BaseAttestationImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = X509Attestation{}

func (s X509Attestation) MarshalJSON() ([]byte, error) {
	type wrapper X509Attestation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling X509Attestation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling X509Attestation: %+v", err)
	}

	decoded["type"] = "x509"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling X509Attestation: %+v", err)
	}

	return encoded, nil
}
