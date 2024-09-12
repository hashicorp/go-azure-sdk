package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = PermissionsDefinitionAzureRole{}

type PermissionsDefinitionAzureRole struct {

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s PermissionsDefinitionAzureRole) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PermissionsDefinitionAzureRole{}

func (s PermissionsDefinitionAzureRole) MarshalJSON() ([]byte, error) {
	type wrapper PermissionsDefinitionAzureRole
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PermissionsDefinitionAzureRole: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PermissionsDefinitionAzureRole: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.permissionsDefinitionAzureRole"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PermissionsDefinitionAzureRole: %+v", err)
	}

	return encoded, nil
}
