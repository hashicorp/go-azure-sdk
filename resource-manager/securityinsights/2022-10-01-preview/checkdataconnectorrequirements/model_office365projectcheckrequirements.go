package checkdataconnectorrequirements

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataConnectorsCheckRequirements = Office365ProjectCheckRequirements{}

type Office365ProjectCheckRequirements struct {
	Properties *DataConnectorTenantId `json:"properties,omitempty"`

	// Fields inherited from DataConnectorsCheckRequirements

	Kind DataConnectorKind `json:"kind"`
}

func (s Office365ProjectCheckRequirements) DataConnectorsCheckRequirements() BaseDataConnectorsCheckRequirementsImpl {
	return BaseDataConnectorsCheckRequirementsImpl{
		Kind: s.Kind,
	}
}

var _ json.Marshaler = Office365ProjectCheckRequirements{}

func (s Office365ProjectCheckRequirements) MarshalJSON() ([]byte, error) {
	type wrapper Office365ProjectCheckRequirements
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Office365ProjectCheckRequirements: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Office365ProjectCheckRequirements: %+v", err)
	}

	decoded["kind"] = "Office365Project"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Office365ProjectCheckRequirements: %+v", err)
	}

	return encoded, nil
}
