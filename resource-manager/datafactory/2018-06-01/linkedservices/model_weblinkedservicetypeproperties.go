package linkedservices

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebLinkedServiceTypeProperties interface {
}

// RawWebLinkedServiceTypePropertiesImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawWebLinkedServiceTypePropertiesImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalWebLinkedServiceTypePropertiesImplementation(input []byte) (WebLinkedServiceTypeProperties, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling WebLinkedServiceTypeProperties into map[string]interface: %+v", err)
	}

	value, ok := temp["authenticationType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Anonymous") {
		var out WebAnonymousAuthentication
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WebAnonymousAuthentication: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Basic") {
		var out WebBasicAuthentication
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WebBasicAuthentication: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "ClientCertificate") {
		var out WebClientCertificateAuthentication
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WebClientCertificateAuthentication: %+v", err)
		}
		return out, nil
	}

	out := RawWebLinkedServiceTypePropertiesImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
