package commands

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommandProperties interface {
}

// RawCommandPropertiesImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawCommandPropertiesImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalCommandPropertiesImplementation(input []byte) (CommandProperties, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling CommandProperties into map[string]interface: %+v", err)
	}

	value, ok := temp["commandType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Migrate.SqlServer.AzureDbSqlMi.Complete") {
		var out MigrateMISyncCompleteCommandProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MigrateMISyncCompleteCommandProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Migrate.Sync.Complete.Database") {
		var out MigrateSyncCompleteCommandProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MigrateSyncCompleteCommandProperties: %+v", err)
		}
		return out, nil
	}

	out := RawCommandPropertiesImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
