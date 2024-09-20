package externalsecuritysolutions

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ExternalSecuritySolution = AtaExternalSecuritySolution{}

type AtaExternalSecuritySolution struct {
	Properties *AtaSolutionProperties `json:"properties,omitempty"`

	// Fields inherited from ExternalSecuritySolution

	Id       *string                       `json:"id,omitempty"`
	Kind     *ExternalSecuritySolutionKind `json:"kind,omitempty"`
	Location *string                       `json:"location,omitempty"`
	Name     *string                       `json:"name,omitempty"`
	Type     *string                       `json:"type,omitempty"`
}

func (s AtaExternalSecuritySolution) ExternalSecuritySolution() BaseExternalSecuritySolutionImpl {
	return BaseExternalSecuritySolutionImpl{
		Id:       s.Id,
		Kind:     s.Kind,
		Location: s.Location,
		Name:     s.Name,
		Type:     s.Type,
	}
}

var _ json.Marshaler = AtaExternalSecuritySolution{}

func (s AtaExternalSecuritySolution) MarshalJSON() ([]byte, error) {
	type wrapper AtaExternalSecuritySolution
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AtaExternalSecuritySolution: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AtaExternalSecuritySolution: %+v", err)
	}

	decoded["kind"] = "ATA"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AtaExternalSecuritySolution: %+v", err)
	}

	return encoded, nil
}
