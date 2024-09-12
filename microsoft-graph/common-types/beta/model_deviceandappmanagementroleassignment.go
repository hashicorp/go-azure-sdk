package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ RoleAssignment = DeviceAndAppManagementRoleAssignment{}

type DeviceAndAppManagementRoleAssignment struct {
	// The list of ids of role member security groups. These are IDs from Azure Active Directory.
	Members *[]string `json:"members,omitempty"`

	// The set of Role Scope Tags defined on the Role Assignment.
	RoleScopeTags *[]RoleScopeTag `json:"roleScopeTags,omitempty"`

	// Fields inherited from RoleAssignment

	// Description of the Role Assignment.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The display or friendly name of the role Assignment.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// List of ids of role scope member security groups. These are IDs from Azure Active Directory.
	ResourceScopes *[]string `json:"resourceScopes,omitempty"`

	// Role definition this assignment is part of.
	RoleDefinition *RoleDefinition `json:"roleDefinition,omitempty"`

	// List of ids of role scope member security groups. These are IDs from Azure Active Directory.
	ScopeMembers *[]string `json:"scopeMembers,omitempty"`

	// Specifies the type of scope for a Role Assignment.
	ScopeType *RoleAssignmentScopeType `json:"scopeType,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s DeviceAndAppManagementRoleAssignment) RoleAssignment() BaseRoleAssignmentImpl {
	return BaseRoleAssignmentImpl{
		Description:    s.Description,
		DisplayName:    s.DisplayName,
		ResourceScopes: s.ResourceScopes,
		RoleDefinition: s.RoleDefinition,
		ScopeMembers:   s.ScopeMembers,
		ScopeType:      s.ScopeType,
		Id:             s.Id,
		ODataId:        s.ODataId,
		ODataType:      s.ODataType,
	}
}

func (s DeviceAndAppManagementRoleAssignment) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceAndAppManagementRoleAssignment{}

func (s DeviceAndAppManagementRoleAssignment) MarshalJSON() ([]byte, error) {
	type wrapper DeviceAndAppManagementRoleAssignment
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceAndAppManagementRoleAssignment: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceAndAppManagementRoleAssignment: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.deviceAndAppManagementRoleAssignment"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceAndAppManagementRoleAssignment: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &DeviceAndAppManagementRoleAssignment{}

func (s *DeviceAndAppManagementRoleAssignment) UnmarshalJSON(bytes []byte) error {
	type alias DeviceAndAppManagementRoleAssignment
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into DeviceAndAppManagementRoleAssignment: %+v", err)
	}

	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.Id = decoded.Id
	s.Members = decoded.Members
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.ResourceScopes = decoded.ResourceScopes
	s.RoleScopeTags = decoded.RoleScopeTags
	s.ScopeMembers = decoded.ScopeMembers
	s.ScopeType = decoded.ScopeType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DeviceAndAppManagementRoleAssignment into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["roleDefinition"]; ok {
		impl, err := UnmarshalRoleDefinitionImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'RoleDefinition' for 'DeviceAndAppManagementRoleAssignment': %+v", err)
		}
		s.RoleDefinition = &impl
	}
	return nil
}
