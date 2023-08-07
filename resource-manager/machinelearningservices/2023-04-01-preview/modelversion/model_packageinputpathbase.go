package modelversion

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PackageInputPathBase interface {
}

// RawModeOfTransitImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawPackageInputPathBaseImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalPackageInputPathBaseImplementation(input []byte) (PackageInputPathBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling PackageInputPathBase into map[string]interface: %+v", err)
	}

	value, ok := temp["inputPathType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "PathId") {
		var out PackageInputPathId
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PackageInputPathId: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Url") {
		var out PackageInputPathUrl
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PackageInputPathUrl: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "PathVersion") {
		var out PackageInputPathVersion
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PackageInputPathVersion: %+v", err)
		}
		return out, nil
	}

	out := RawPackageInputPathBaseImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
