package checkdataconnectorrequirements

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataConnectorsCheckRequirements = Dynamics365CheckRequirements{}

type Dynamics365CheckRequirements struct {
	Properties *DataConnectorTenantId `json:"properties,omitempty"`

	// Fields inherited from DataConnectorsCheckRequirements

	Kind DataConnectorKind `json:"kind"`
}

func (s Dynamics365CheckRequirements) DataConnectorsCheckRequirements() BaseDataConnectorsCheckRequirementsImpl {
	return BaseDataConnectorsCheckRequirementsImpl{
		Kind: s.Kind,
	}
}

var _ json.Marshaler = Dynamics365CheckRequirements{}

func (s Dynamics365CheckRequirements) MarshalJSON() ([]byte, error) {
	type wrapper Dynamics365CheckRequirements
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Dynamics365CheckRequirements: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Dynamics365CheckRequirements: %+v", err)
	}

	decoded["kind"] = "Dynamics365"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Dynamics365CheckRequirements: %+v", err)
	}

	return encoded, nil
}
