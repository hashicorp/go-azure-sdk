package service

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PartitionSchemeDescription = UniformInt64RangePartitionSchemeDescription{}

type UniformInt64RangePartitionSchemeDescription struct {
	Count   int64  `json:"count"`
	HighKey string `json:"highKey"`
	LowKey  string `json:"lowKey"`

	// Fields inherited from PartitionSchemeDescription

	PartitionScheme PartitionScheme `json:"partitionScheme"`
}

func (s UniformInt64RangePartitionSchemeDescription) PartitionSchemeDescription() BasePartitionSchemeDescriptionImpl {
	return BasePartitionSchemeDescriptionImpl{
		PartitionScheme: s.PartitionScheme,
	}
}

var _ json.Marshaler = UniformInt64RangePartitionSchemeDescription{}

func (s UniformInt64RangePartitionSchemeDescription) MarshalJSON() ([]byte, error) {
	type wrapper UniformInt64RangePartitionSchemeDescription
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UniformInt64RangePartitionSchemeDescription: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UniformInt64RangePartitionSchemeDescription: %+v", err)
	}

	decoded["partitionScheme"] = "UniformInt64Range"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UniformInt64RangePartitionSchemeDescription: %+v", err)
	}

	return encoded, nil
}
