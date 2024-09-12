package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ShiftsRoleDefinition{}

type ShiftsRoleDefinition struct {
	// The description of the role.
	Description *string `json:"description,omitempty"`

	// The display name of the role.
	DisplayName *string `json:"displayName,omitempty"`

	// The collection of role permissions within the role.
	ShiftsRolePermissions *[]ShiftsRolePermission `json:"shiftsRolePermissions,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s ShiftsRoleDefinition) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ShiftsRoleDefinition{}

func (s ShiftsRoleDefinition) MarshalJSON() ([]byte, error) {
	type wrapper ShiftsRoleDefinition
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ShiftsRoleDefinition: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ShiftsRoleDefinition: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.shiftsRoleDefinition"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ShiftsRoleDefinition: %+v", err)
	}

	return encoded, nil
}
