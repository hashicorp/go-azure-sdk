package service

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PartitionSchemeDescription = SingletonPartitionSchemeDescription{}

type SingletonPartitionSchemeDescription struct {

	// Fields inherited from PartitionSchemeDescription

	PartitionScheme PartitionScheme `json:"partitionScheme"`
}

func (s SingletonPartitionSchemeDescription) PartitionSchemeDescription() BasePartitionSchemeDescriptionImpl {
	return BasePartitionSchemeDescriptionImpl{
		PartitionScheme: s.PartitionScheme,
	}
}

var _ json.Marshaler = SingletonPartitionSchemeDescription{}

func (s SingletonPartitionSchemeDescription) MarshalJSON() ([]byte, error) {
	type wrapper SingletonPartitionSchemeDescription
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SingletonPartitionSchemeDescription: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SingletonPartitionSchemeDescription: %+v", err)
	}

	decoded["partitionScheme"] = "Singleton"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SingletonPartitionSchemeDescription: %+v", err)
	}

	return encoded, nil
}
