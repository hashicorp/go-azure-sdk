package checkdataconnectorrequirements

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataConnectorsCheckRequirements = OfficeIRMCheckRequirements{}

type OfficeIRMCheckRequirements struct {
	Properties *DataConnectorTenantId `json:"properties,omitempty"`

	// Fields inherited from DataConnectorsCheckRequirements
}

var _ json.Marshaler = OfficeIRMCheckRequirements{}

func (s OfficeIRMCheckRequirements) MarshalJSON() ([]byte, error) {
	type wrapper OfficeIRMCheckRequirements
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OfficeIRMCheckRequirements: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OfficeIRMCheckRequirements: %+v", err)
	}
	decoded["kind"] = "OfficeIRM"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OfficeIRMCheckRequirements: %+v", err)
	}

	return encoded, nil
}
