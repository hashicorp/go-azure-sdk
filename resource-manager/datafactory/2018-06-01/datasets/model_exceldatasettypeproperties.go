package datasets

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExcelDatasetTypeProperties struct {
	Compression      *DatasetCompression `json:"compression,omitempty"`
	FirstRowAsHeader *interface{}        `json:"firstRowAsHeader,omitempty"`
	Location         DatasetLocation     `json:"location"`
	NullValue        *interface{}        `json:"nullValue,omitempty"`
	Range            *interface{}        `json:"range,omitempty"`
	SheetIndex       *interface{}        `json:"sheetIndex,omitempty"`
	SheetName        *interface{}        `json:"sheetName,omitempty"`
}

var _ json.Unmarshaler = &ExcelDatasetTypeProperties{}

func (s *ExcelDatasetTypeProperties) UnmarshalJSON(bytes []byte) error {
	type alias ExcelDatasetTypeProperties
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into ExcelDatasetTypeProperties: %+v", err)
	}

	s.Compression = decoded.Compression
	s.FirstRowAsHeader = decoded.FirstRowAsHeader
	s.NullValue = decoded.NullValue
	s.Range = decoded.Range
	s.SheetIndex = decoded.SheetIndex
	s.SheetName = decoded.SheetName

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ExcelDatasetTypeProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["location"]; ok {
		impl, err := unmarshalDatasetLocationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Location' for 'ExcelDatasetTypeProperties': %+v", err)
		}
		s.Location = impl
	}
	return nil
}
