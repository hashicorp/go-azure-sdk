package autonomousdatabases

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutonomousDatabaseBaseProperties interface {
}

// RawAutonomousDatabaseBasePropertiesImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAutonomousDatabaseBasePropertiesImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalAutonomousDatabaseBasePropertiesImplementation(input []byte) (AutonomousDatabaseBaseProperties, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AutonomousDatabaseBaseProperties into map[string]interface: %+v", err)
	}

	value, ok := temp["dataBaseType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Clone") {
		var out AutonomousDatabaseCloneProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AutonomousDatabaseCloneProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Regular") {
		var out AutonomousDatabaseProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AutonomousDatabaseProperties: %+v", err)
		}
		return out, nil
	}

	out := RawAutonomousDatabaseBasePropertiesImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
