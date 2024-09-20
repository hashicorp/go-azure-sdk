package alerts

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ResourceIdentifier = AzureResourceIdentifier{}

type AzureResourceIdentifier struct {
	AzureResourceId *string `json:"azureResourceId,omitempty"`

	// Fields inherited from ResourceIdentifier

	Type ResourceIdentifierType `json:"type"`
}

func (s AzureResourceIdentifier) ResourceIdentifier() BaseResourceIdentifierImpl {
	return BaseResourceIdentifierImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = AzureResourceIdentifier{}

func (s AzureResourceIdentifier) MarshalJSON() ([]byte, error) {
	type wrapper AzureResourceIdentifier
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureResourceIdentifier: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureResourceIdentifier: %+v", err)
	}

	decoded["type"] = "AzureResource"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureResourceIdentifier: %+v", err)
	}

	return encoded, nil
}
