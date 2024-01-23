package roles

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Role = CloudEdgeManagementRole{}

type CloudEdgeManagementRole struct {
	Properties *CloudEdgeManagementRoleProperties `json:"properties,omitempty"`

	// Fields inherited from Role
	Id         *string                `json:"id,omitempty"`
	Name       *string                `json:"name,omitempty"`
	SystemData *systemdata.SystemData `json:"systemData,omitempty"`
	Type       *string                `json:"type,omitempty"`
}

var _ json.Marshaler = CloudEdgeManagementRole{}

func (s CloudEdgeManagementRole) MarshalJSON() ([]byte, error) {
	type wrapper CloudEdgeManagementRole
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CloudEdgeManagementRole: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CloudEdgeManagementRole: %+v", err)
	}
	decoded["kind"] = "CloudEdgeManagement"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CloudEdgeManagementRole: %+v", err)
	}

	return encoded, nil
}
