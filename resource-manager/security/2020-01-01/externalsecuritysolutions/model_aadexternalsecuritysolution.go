package externalsecuritysolutions

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ExternalSecuritySolution = AadExternalSecuritySolution{}

type AadExternalSecuritySolution struct {
	Properties *AadSolutionProperties `json:"properties,omitempty"`

	// Fields inherited from ExternalSecuritySolution
	Id       *string `json:"id,omitempty"`
	Location *string `json:"location,omitempty"`
	Name     *string `json:"name,omitempty"`
	Type     *string `json:"type,omitempty"`
}

var _ json.Marshaler = AadExternalSecuritySolution{}

func (s AadExternalSecuritySolution) MarshalJSON() ([]byte, error) {
	type wrapper AadExternalSecuritySolution
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AadExternalSecuritySolution: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AadExternalSecuritySolution: %+v", err)
	}
	decoded["kind"] = "AAD"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AadExternalSecuritySolution: %+v", err)
	}

	return encoded, nil
}
