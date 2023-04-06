package artifact

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Artifact = PolicyAssignmentArtifact{}

type PolicyAssignmentArtifact struct {
	Properties PolicyAssignmentArtifactProperties `json:"properties"`

	// Fields inherited from Artifact
	Id   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

var _ json.Marshaler = PolicyAssignmentArtifact{}

func (s PolicyAssignmentArtifact) MarshalJSON() ([]byte, error) {
	type wrapper PolicyAssignmentArtifact
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PolicyAssignmentArtifact: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PolicyAssignmentArtifact: %+v", err)
	}
	decoded["kind"] = "policyAssignment"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PolicyAssignmentArtifact: %+v", err)
	}

	return encoded, nil
}
