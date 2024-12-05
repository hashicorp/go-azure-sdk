package dataversion

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataVersionBase interface {
	DataVersionBase() BaseDataVersionBaseImpl
}

var _ DataVersionBase = BaseDataVersionBaseImpl{}

type BaseDataVersionBaseImpl struct {
	DataType    DataType           `json:"dataType"`
	DataUri     string             `json:"dataUri"`
	Description *string            `json:"description,omitempty"`
	IsAnonymous *bool              `json:"isAnonymous,omitempty"`
	IsArchived  *bool              `json:"isArchived,omitempty"`
	Properties  *map[string]string `json:"properties,omitempty"`
	Tags        *map[string]string `json:"tags,omitempty"`
}

func (s BaseDataVersionBaseImpl) DataVersionBase() BaseDataVersionBaseImpl {
	return s
}

var _ DataVersionBase = RawDataVersionBaseImpl{}

// RawDataVersionBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDataVersionBaseImpl struct {
	dataVersionBase BaseDataVersionBaseImpl
	Type            string
	Values          map[string]interface{}
}

func (s RawDataVersionBaseImpl) DataVersionBase() BaseDataVersionBaseImpl {
	return s.dataVersionBase
}

func UnmarshalDataVersionBaseImplementation(input []byte) (DataVersionBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DataVersionBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["dataType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "mltable") {
		var out MLTableData
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MLTableData: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "uri_file") {
		var out UriFileDataVersion
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UriFileDataVersion: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "uri_folder") {
		var out UriFolderDataVersion
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UriFolderDataVersion: %+v", err)
		}
		return out, nil
	}

	var parent BaseDataVersionBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDataVersionBaseImpl: %+v", err)
	}

	return RawDataVersionBaseImpl{
		dataVersionBase: parent,
		Type:            value,
		Values:          temp,
	}, nil

}
