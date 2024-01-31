package migratessistasks

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MigrateSsisTaskOutput interface {
}

// RawMigrateSsisTaskOutputImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawMigrateSsisTaskOutputImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalMigrateSsisTaskOutputImplementation(input []byte) (MigrateSsisTaskOutput, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling MigrateSsisTaskOutput into map[string]interface: %+v", err)
	}

	value, ok := temp["resultType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "MigrationLevelOutput") {
		var out MigrateSsisTaskOutputMigrationLevel
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MigrateSsisTaskOutputMigrationLevel: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "SsisProjectLevelOutput") {
		var out MigrateSsisTaskOutputProjectLevel
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MigrateSsisTaskOutputProjectLevel: %+v", err)
		}
		return out, nil
	}

	out := RawMigrateSsisTaskOutputImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
