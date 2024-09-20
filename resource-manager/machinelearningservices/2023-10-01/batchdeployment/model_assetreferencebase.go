package batchdeployment

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssetReferenceBase interface {
	AssetReferenceBase() BaseAssetReferenceBaseImpl
}

var _ AssetReferenceBase = BaseAssetReferenceBaseImpl{}

type BaseAssetReferenceBaseImpl struct {
	ReferenceType ReferenceType `json:"referenceType"`
}

func (s BaseAssetReferenceBaseImpl) AssetReferenceBase() BaseAssetReferenceBaseImpl {
	return s
}

var _ AssetReferenceBase = RawAssetReferenceBaseImpl{}

// RawAssetReferenceBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAssetReferenceBaseImpl struct {
	assetReferenceBase BaseAssetReferenceBaseImpl
	Type               string
	Values             map[string]interface{}
}

func (s RawAssetReferenceBaseImpl) AssetReferenceBase() BaseAssetReferenceBaseImpl {
	return s.assetReferenceBase
}

func UnmarshalAssetReferenceBaseImplementation(input []byte) (AssetReferenceBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AssetReferenceBase into map[string]interface: %+v", err)
	}

	value, ok := temp["referenceType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "DataPath") {
		var out DataPathAssetReference
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DataPathAssetReference: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Id") {
		var out IdAssetReference
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdAssetReference: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "OutputPath") {
		var out OutputPathAssetReference
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OutputPathAssetReference: %+v", err)
		}
		return out, nil
	}

	var parent BaseAssetReferenceBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAssetReferenceBaseImpl: %+v", err)
	}

	return RawAssetReferenceBaseImpl{
		assetReferenceBase: parent,
		Type:               value,
		Values:             temp,
	}, nil

}
