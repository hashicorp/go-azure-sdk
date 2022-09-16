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

	type RawLinkedIntegrationRuntimeTypeImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawLinkedIntegrationRuntimeTypeImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
