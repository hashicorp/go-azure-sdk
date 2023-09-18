package labelingjob

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExportSummary interface {
}

// RawExportSummaryImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawExportSummaryImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalExportSummaryImplementation(input []byte) (ExportSummary, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ExportSummary into map[string]interface: %+v", err)
	}

	value, ok := temp["format"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Coco") {
		var out CocoExportSummary
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CocoExportSummary: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "CSV") {
		var out CsvExportSummary
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CsvExportSummary: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Dataset") {
		var out DatasetExportSummary
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DatasetExportSummary: %+v", err)
		}
		return out, nil
	}

	out := RawExportSummaryImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
