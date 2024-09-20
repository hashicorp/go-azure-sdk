package service

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PartitionSchemeDescription = NamedPartitionSchemeDescription{}

type NamedPartitionSchemeDescription struct {
	Count int64    `json:"count"`
	Names []string `json:"names"`

	// Fields inherited from PartitionSchemeDescription

	PartitionScheme PartitionScheme `json:"partitionScheme"`
}

func (s NamedPartitionSchemeDescription) PartitionSchemeDescription() BasePartitionSchemeDescriptionImpl {
	return BasePartitionSchemeDescriptionImpl{
		PartitionScheme: s.PartitionScheme,
	}
}

var _ json.Marshaler = NamedPartitionSchemeDescription{}

func (s NamedPartitionSchemeDescription) MarshalJSON() ([]byte, error) {
	type wrapper NamedPartitionSchemeDescription
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling NamedPartitionSchemeDescription: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling NamedPartitionSchemeDescription: %+v", err)
	}

	decoded["partitionScheme"] = "Named"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling NamedPartitionSchemeDescription: %+v", err)
	}

	return encoded, nil
}
