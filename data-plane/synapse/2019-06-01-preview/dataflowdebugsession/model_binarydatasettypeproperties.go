package dataflowdebugsession

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BinaryDatasetTypeProperties struct {
	Compression DatasetCompression `json:"compression"`
	Location    DatasetLocation    `json:"location"`
}

var _ json.Unmarshaler = &BinaryDatasetTypeProperties{}

func (s *BinaryDatasetTypeProperties) UnmarshalJSON(bytes []byte) error {

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BinaryDatasetTypeProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["compression"]; ok {
		impl, err := UnmarshalDatasetCompressionImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Compression' for 'BinaryDatasetTypeProperties': %+v", err)
		}
		s.Compression = impl
	}

	if v, ok := temp["location"]; ok {
		impl, err := UnmarshalDatasetLocationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Location' for 'BinaryDatasetTypeProperties': %+v", err)
		}
		s.Location = impl
	}

	return nil
}
