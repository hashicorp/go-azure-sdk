package checkdataconnectorrequirements

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataConnectorsCheckRequirements = OfficeATPCheckRequirements{}

type OfficeATPCheckRequirements struct {
	Properties *DataConnectorTenantId `json:"properties,omitempty"`

	// Fields inherited from DataConnectorsCheckRequirements
}

var _ json.Marshaler = OfficeATPCheckRequirements{}

func (s OfficeATPCheckRequirements) MarshalJSON() ([]byte, error) {
	type wrapper OfficeATPCheckRequirements
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OfficeATPCheckRequirements: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OfficeATPCheckRequirements: %+v", err)
	}
	decoded["kind"] = "OfficeATP"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OfficeATPCheckRequirements: %+v", err)
	}

	return encoded, nil
}
