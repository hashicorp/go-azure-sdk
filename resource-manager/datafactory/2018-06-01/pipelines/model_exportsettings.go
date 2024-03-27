package pipelines

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExportSettings interface {
}

// RawExportSettingsImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawExportSettingsImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalExportSettingsImplementation(input []byte) (ExportSettings, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ExportSettings into map[string]interface: %+v", err)
	}

	value, ok := temp["type"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "AzureDatabricksDeltaLakeExportCommand") {
		var out AzureDatabricksDeltaLakeExportCommand
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureDatabricksDeltaLakeExportCommand: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "SnowflakeExportCopyCommand") {
		var out SnowflakeExportCopyCommand
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SnowflakeExportCopyCommand: %+v", err)
		}
		return out, nil
	}

	out := RawExportSettingsImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
