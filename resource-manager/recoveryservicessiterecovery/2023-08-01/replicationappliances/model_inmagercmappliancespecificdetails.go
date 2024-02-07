package replicationappliances

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ApplianceSpecificDetails = InMageRcmApplianceSpecificDetails{}

type InMageRcmApplianceSpecificDetails struct {
	Appliances *[]InMageRcmApplianceDetails `json:"appliances,omitempty"`

	// Fields inherited from ApplianceSpecificDetails
}

var _ json.Marshaler = InMageRcmApplianceSpecificDetails{}

func (s InMageRcmApplianceSpecificDetails) MarshalJSON() ([]byte, error) {
	type wrapper InMageRcmApplianceSpecificDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling InMageRcmApplianceSpecificDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling InMageRcmApplianceSpecificDetails: %+v", err)
	}
	decoded["instanceType"] = "InMageRcm"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling InMageRcmApplianceSpecificDetails: %+v", err)
	}

	return encoded, nil
}
