package checkdataconnectorrequirements

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataConnectorsCheckRequirements = TiTaxiiCheckRequirements{}

type TiTaxiiCheckRequirements struct {
	Properties *DataConnectorTenantId `json:"properties,omitempty"`

	// Fields inherited from DataConnectorsCheckRequirements
}

var _ json.Marshaler = TiTaxiiCheckRequirements{}

func (s TiTaxiiCheckRequirements) MarshalJSON() ([]byte, error) {
	type wrapper TiTaxiiCheckRequirements
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TiTaxiiCheckRequirements: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TiTaxiiCheckRequirements: %+v", err)
	}
	decoded["kind"] = "ThreatIntelligenceTaxii"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TiTaxiiCheckRequirements: %+v", err)
	}

	return encoded, nil
}
