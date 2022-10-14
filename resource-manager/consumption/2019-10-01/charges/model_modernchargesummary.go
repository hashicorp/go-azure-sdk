package charges

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ChargeSummary = ModernChargeSummary{}

type ModernChargeSummary struct {
	Properties ModernChargeSummaryProperties `json:"properties"`

	// Fields inherited from ChargeSummary
	ETag *string            `json:"eTag,omitempty"`
	Id   *string            `json:"id,omitempty"`
	Name *string            `json:"name,omitempty"`
	Tags *map[string]string `json:"tags,omitempty"`
	Type *string            `json:"type,omitempty"`
}

var _ json.Marshaler = ModernChargeSummary{}

func (s ModernChargeSummary) MarshalJSON() ([]byte, error) {
	type wrapper ModernChargeSummary
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ModernChargeSummary: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ModernChargeSummary: %+v", err)
	}
	decoded["kind"] = "modern"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ModernChargeSummary: %+v", err)
	}

	return encoded, nil
}
