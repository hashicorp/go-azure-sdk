package secrets

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecretParameters interface {
}

// RawSecretParametersImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawSecretParametersImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalSecretParametersImplementation(input []byte) (SecretParameters, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling SecretParameters into map[string]interface: %+v", err)
	}

	value, ok := temp["type"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "AzureFirstPartyManagedCertificate") {
		var out AzureFirstPartyManagedCertificateParameters
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureFirstPartyManagedCertificateParameters: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "CustomerCertificate") {
		var out CustomerCertificateParameters
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CustomerCertificateParameters: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "ManagedCertificate") {
		var out ManagedCertificateParameters
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedCertificateParameters: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "UrlSigningKey") {
		var out UrlSigningKeyParameters
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UrlSigningKeyParameters: %+v", err)
		}
		return out, nil
	}

	out := RawSecretParametersImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
