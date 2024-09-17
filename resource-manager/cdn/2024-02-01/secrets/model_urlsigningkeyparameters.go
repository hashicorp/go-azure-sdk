package secrets

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecretParameters = UrlSigningKeyParameters{}

type UrlSigningKeyParameters struct {
	KeyId         string            `json:"keyId"`
	SecretSource  ResourceReference `json:"secretSource"`
	SecretVersion *string           `json:"secretVersion,omitempty"`

	// Fields inherited from SecretParameters

	Type SecretType `json:"type"`
}

func (s UrlSigningKeyParameters) SecretParameters() BaseSecretParametersImpl {
	return BaseSecretParametersImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = UrlSigningKeyParameters{}

func (s UrlSigningKeyParameters) MarshalJSON() ([]byte, error) {
	type wrapper UrlSigningKeyParameters
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UrlSigningKeyParameters: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UrlSigningKeyParameters: %+v", err)
	}

	decoded["type"] = "UrlSigningKey"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UrlSigningKeyParameters: %+v", err)
	}

	return encoded, nil
}
