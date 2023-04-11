package externalsecuritysolutions

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ExternalSecuritySolution = CefExternalSecuritySolution{}

type CefExternalSecuritySolution struct {
	Properties *CefSolutionProperties `json:"properties,omitempty"`

	// Fields inherited from ExternalSecuritySolution
	Id       *string `json:"id,omitempty"`
	Location *string `json:"location,omitempty"`
	Name     *string `json:"name,omitempty"`
	Type     *string `json:"type,omitempty"`
}

var _ json.Marshaler = CefExternalSecuritySolution{}

func (s CefExternalSecuritySolution) MarshalJSON() ([]byte, error) {
	type wrapper CefExternalSecuritySolution
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CefExternalSecuritySolution: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CefExternalSecuritySolution: %+v", err)
	}
	decoded["kind"] = "CEF"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CefExternalSecuritySolution: %+v", err)
	}

	return encoded, nil
}
