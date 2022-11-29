package checkdataconnectorrequirements

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataConnectorsCheckRequirements = MSTICheckRequirements{}

type MSTICheckRequirements struct {
	Properties *DataConnectorTenantId `json:"properties,omitempty"`

	// Fields inherited from DataConnectorsCheckRequirements
}

var _ json.Marshaler = MSTICheckRequirements{}

func (s MSTICheckRequirements) MarshalJSON() ([]byte, error) {
	type wrapper MSTICheckRequirements
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MSTICheckRequirements: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MSTICheckRequirements: %+v", err)
	}
	decoded["kind"] = "MicrosoftThreatIntelligence"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MSTICheckRequirements: %+v", err)
	}

	return encoded, nil
}
