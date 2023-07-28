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

	type RawExportSummaryImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawExportSummaryImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
