package dataflowdebugsession

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DatasetCompression interface {
	DatasetCompression() BaseDatasetCompressionImpl
}

var _ DatasetCompression = BaseDatasetCompressionImpl{}

type BaseDatasetCompressionImpl struct {
	Type string `json:"type"`
}

func (s BaseDatasetCompressionImpl) DatasetCompression() BaseDatasetCompressionImpl {
	return s
}

var _ DatasetCompression = RawDatasetCompressionImpl{}

// RawDatasetCompressionImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDatasetCompressionImpl struct {
	datasetCompression BaseDatasetCompressionImpl
	Type               string
	Values             map[string]interface{}
}

func (s RawDatasetCompressionImpl) DatasetCompression() BaseDatasetCompressionImpl {
	return s.datasetCompression
}

func UnmarshalDatasetCompressionImplementation(input []byte) (DatasetCompression, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DatasetCompression into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "BZip2") {
		var out DatasetBZip2Compression
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DatasetBZip2Compression: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Deflate") {
		var out DatasetDeflateCompression
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DatasetDeflateCompression: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "GZip") {
		var out DatasetGZipCompression
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DatasetGZipCompression: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Tar") {
		var out DatasetTarCompression
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DatasetTarCompression: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "TarGZip") {
		var out DatasetTarGZipCompression
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DatasetTarGZipCompression: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "ZipDeflate") {
		var out DatasetZipDeflateCompression
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DatasetZipDeflateCompression: %+v", err)
		}
		return out, nil
	}

	var parent BaseDatasetCompressionImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDatasetCompressionImpl: %+v", err)
	}

	return RawDatasetCompressionImpl{
		datasetCompression: parent,
		Type:               value,
		Values:             temp,
	}, nil

}
