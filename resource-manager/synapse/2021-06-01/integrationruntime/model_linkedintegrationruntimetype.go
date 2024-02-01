package integrationruntime

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LinkedIntegrationRuntimeType interface {
}

// RawLinkedIntegrationRuntimeTypeImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawLinkedIntegrationRuntimeTypeImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalLinkedIntegrationRuntimeTypeImplementation(input []byte) (LinkedIntegrationRuntimeType, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling LinkedIntegrationRuntimeType into map[string]interface: %+v", err)
	}

	value, ok := temp["authorizationType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Key") {
		var out LinkedIntegrationRuntimeKeyAuthorization
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LinkedIntegrationRuntimeKeyAuthorization: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "RBAC") {
		var out LinkedIntegrationRuntimeRbacAuthorization
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LinkedIntegrationRuntimeRbacAuthorization: %+v", err)
		}
		return out, nil
	}

	out := RawLinkedIntegrationRuntimeTypeImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
