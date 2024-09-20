package artifact

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Artifact interface {
	Artifact() BaseArtifactImpl
}

var _ Artifact = BaseArtifactImpl{}

type BaseArtifactImpl struct {
	Id   *string      `json:"id,omitempty"`
	Kind ArtifactKind `json:"kind"`
	Name *string      `json:"name,omitempty"`
	Type *string      `json:"type,omitempty"`
}

func (s BaseArtifactImpl) Artifact() BaseArtifactImpl {
	return s
}

var _ Artifact = RawArtifactImpl{}

// RawArtifactImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawArtifactImpl struct {
	artifact BaseArtifactImpl
	Type     string
	Values   map[string]interface{}
}

func (s RawArtifactImpl) Artifact() BaseArtifactImpl {
	return s.artifact
}

func UnmarshalArtifactImplementation(input []byte) (Artifact, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling Artifact into map[string]interface: %+v", err)
	}

	value, ok := temp["kind"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "policyAssignment") {
		var out PolicyAssignmentArtifact
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PolicyAssignmentArtifact: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "roleAssignment") {
		var out RoleAssignmentArtifact
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RoleAssignmentArtifact: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "template") {
		var out TemplateArtifact
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TemplateArtifact: %+v", err)
		}
		return out, nil
	}

	var parent BaseArtifactImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseArtifactImpl: %+v", err)
	}

	return RawArtifactImpl{
		artifact: parent,
		Type:     value,
		Values:   temp,
	}, nil

}
