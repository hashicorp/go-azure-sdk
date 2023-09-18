package externalsecuritysolutions

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalSecuritySolution interface {
}

// RawExternalSecuritySolutionImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawExternalSecuritySolutionImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalExternalSecuritySolutionImplementation(input []byte) (ExternalSecuritySolution, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ExternalSecuritySolution into map[string]interface: %+v", err)
	}

	value, ok := temp["kind"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "AAD") {
		var out AadExternalSecuritySolution
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AadExternalSecuritySolution: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "ATA") {
		var out AtaExternalSecuritySolution
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AtaExternalSecuritySolution: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "CEF") {
		var out CefExternalSecuritySolution
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CefExternalSecuritySolution: %+v", err)
		}
		return out, nil
	}

	out := RawExternalSecuritySolutionImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
