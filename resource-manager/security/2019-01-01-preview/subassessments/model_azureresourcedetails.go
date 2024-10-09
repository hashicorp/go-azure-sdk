package subassessments

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ResourceDetails = AzureResourceDetails{}

type AzureResourceDetails struct {
	Id *string `json:"id,omitempty"`

	// Fields inherited from ResourceDetails

	Source Source `json:"source"`
}

func (s AzureResourceDetails) ResourceDetails() BaseResourceDetailsImpl {
	return BaseResourceDetailsImpl{
		Source: s.Source,
	}
}

var _ json.Marshaler = AzureResourceDetails{}

func (s AzureResourceDetails) MarshalJSON() ([]byte, error) {
	type wrapper AzureResourceDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureResourceDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureResourceDetails: %+v", err)
	}

	decoded["source"] = "Azure"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureResourceDetails: %+v", err)
	}

	return encoded, nil
}
