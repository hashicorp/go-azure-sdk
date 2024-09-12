package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SubjectSet = GroupMembers{}

type GroupMembers struct {
	// The name of the group in Microsoft Entra ID. Read only.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The ID of the group in Microsoft Entra ID.
	GroupId nullable.Type[string] `json:"groupId,omitempty"`

	// Fields inherited from SubjectSet

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s GroupMembers) SubjectSet() BaseSubjectSetImpl {
	return BaseSubjectSetImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = GroupMembers{}

func (s GroupMembers) MarshalJSON() ([]byte, error) {
	type wrapper GroupMembers
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GroupMembers: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GroupMembers: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.groupMembers"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GroupMembers: %+v", err)
	}

	return encoded, nil
}
