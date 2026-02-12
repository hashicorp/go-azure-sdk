package datasets

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type XmlDatasetTypeProperties struct {
	Compression  DatasetCompression `json:"compression"`
	EncodingName *interface{}       `json:"encodingName,omitempty"`
	Location     DatasetLocation    `json:"location"`
	NullValue    *interface{}       `json:"nullValue,omitempty"`
}

var _ json.Unmarshaler = &XmlDatasetTypeProperties{}

func (s *XmlDatasetTypeProperties) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		EncodingName *interface{} `json:"encodingName,omitempty"`
		NullValue    *interface{} `json:"nullValue,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.EncodingName = decoded.EncodingName
	s.NullValue = decoded.NullValue

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling XmlDatasetTypeProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["compression"]; ok {
		impl, err := UnmarshalDatasetCompressionImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Compression' for 'XmlDatasetTypeProperties': %+v", err)
		}
		s.Compression = impl
	}

	if v, ok := temp["location"]; ok {
		impl, err := UnmarshalDatasetLocationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Location' for 'XmlDatasetTypeProperties': %+v", err)
		}
		s.Location = impl
	}

	return nil
}
