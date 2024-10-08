package subscriptions

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ OutputDataSource = AzureFunctionOutputDataSource{}

type AzureFunctionOutputDataSource struct {
	Properties *AzureFunctionOutputDataSourceProperties `json:"properties,omitempty"`

	// Fields inherited from OutputDataSource

	Type string `json:"type"`
}

func (s AzureFunctionOutputDataSource) OutputDataSource() BaseOutputDataSourceImpl {
	return BaseOutputDataSourceImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = AzureFunctionOutputDataSource{}

func (s AzureFunctionOutputDataSource) MarshalJSON() ([]byte, error) {
	type wrapper AzureFunctionOutputDataSource
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureFunctionOutputDataSource: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureFunctionOutputDataSource: %+v", err)
	}

	decoded["type"] = "Microsoft.AzureFunction"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureFunctionOutputDataSource: %+v", err)
	}

	return encoded, nil
}
