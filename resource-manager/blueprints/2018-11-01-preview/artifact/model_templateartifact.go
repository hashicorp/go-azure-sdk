package artifact

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Artifact = TemplateArtifact{}

type TemplateArtifact struct {
	Properties TemplateArtifactProperties `json:"properties"`

	// Fields inherited from Artifact

	Id   *string      `json:"id,omitempty"`
	Kind ArtifactKind `json:"kind"`
	Name *string      `json:"name,omitempty"`
	Type *string      `json:"type,omitempty"`
}

func (s TemplateArtifact) Artifact() BaseArtifactImpl {
	return BaseArtifactImpl{
		Id:   s.Id,
		Kind: s.Kind,
		Name: s.Name,
		Type: s.Type,
	}
}

var _ json.Marshaler = TemplateArtifact{}

func (s TemplateArtifact) MarshalJSON() ([]byte, error) {
	type wrapper TemplateArtifact
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TemplateArtifact: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TemplateArtifact: %+v", err)
	}

	decoded["kind"] = "template"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TemplateArtifact: %+v", err)
	}

	return encoded, nil
}
