package checkdataconnectorrequirements

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataConnectorsCheckRequirements = AADCheckRequirements{}

type AADCheckRequirements struct {
	Properties *DataConnectorTenantId `json:"properties,omitempty"`

	// Fields inherited from DataConnectorsCheckRequirements

	Kind DataConnectorKind `json:"kind"`
}

func (s AADCheckRequirements) DataConnectorsCheckRequirements() BaseDataConnectorsCheckRequirementsImpl {
	return BaseDataConnectorsCheckRequirementsImpl{
		Kind: s.Kind,
	}
}

var _ json.Marshaler = AADCheckRequirements{}

func (s AADCheckRequirements) MarshalJSON() ([]byte, error) {
	type wrapper AADCheckRequirements
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AADCheckRequirements: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AADCheckRequirements: %+v", err)
	}

	decoded["kind"] = "AzureActiveDirectory"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AADCheckRequirements: %+v", err)
	}

	return encoded, nil
}
