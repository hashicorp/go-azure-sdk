package checkdataconnectorrequirements

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataConnectorsCheckRequirements = AATPCheckRequirements{}

type AATPCheckRequirements struct {
	Properties *DataConnectorTenantId `json:"properties,omitempty"`

	// Fields inherited from DataConnectorsCheckRequirements

	Kind DataConnectorKind `json:"kind"`
}

func (s AATPCheckRequirements) DataConnectorsCheckRequirements() BaseDataConnectorsCheckRequirementsImpl {
	return BaseDataConnectorsCheckRequirementsImpl{
		Kind: s.Kind,
	}
}

var _ json.Marshaler = AATPCheckRequirements{}

func (s AATPCheckRequirements) MarshalJSON() ([]byte, error) {
	type wrapper AATPCheckRequirements
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AATPCheckRequirements: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AATPCheckRequirements: %+v", err)
	}

	decoded["kind"] = "AzureAdvancedThreatProtection"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AATPCheckRequirements: %+v", err)
	}

	return encoded, nil
}
