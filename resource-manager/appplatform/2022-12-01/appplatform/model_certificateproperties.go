package appplatform

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CertificateProperties interface {
}

func unmarshalCertificatePropertiesImplementation(input []byte) (CertificateProperties, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling CertificateProperties into map[string]interface: %+v", err)
	}

	value, ok := temp["type"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "ContentCertificate") {
		var out ContentCertificateProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ContentCertificateProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "KeyVaultCertificate") {
		var out KeyVaultCertificateProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into KeyVaultCertificateProperties: %+v", err)
		}
		return out, nil
	}

	type RawCertificatePropertiesImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawCertificatePropertiesImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
