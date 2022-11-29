package checkdataconnectorrequirements

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataConnectorsCheckRequirements = MtpCheckRequirements{}

type MtpCheckRequirements struct {
	Properties *DataConnectorTenantId `json:"properties,omitempty"`

	// Fields inherited from DataConnectorsCheckRequirements
}

var _ json.Marshaler = MtpCheckRequirements{}

func (s MtpCheckRequirements) MarshalJSON() ([]byte, error) {
	type wrapper MtpCheckRequirements
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MtpCheckRequirements: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MtpCheckRequirements: %+v", err)
	}
	decoded["kind"] = "MicrosoftThreatProtection"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MtpCheckRequirements: %+v", err)
	}

	return encoded, nil
}
