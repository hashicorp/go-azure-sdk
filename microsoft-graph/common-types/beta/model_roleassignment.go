package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleAssignment interface {
	Entity
	RoleAssignment() BaseRoleAssignmentImpl
}

var _ RoleAssignment = BaseRoleAssignmentImpl{}

type BaseRoleAssignmentImpl struct {
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

func (s BaseRoleAssignmentImpl) RoleAssignment() BaseRoleAssignmentImpl {
	return s
}

func (s BaseRoleAssignmentImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ RoleAssignment = RawRoleAssignmentImpl{}

// RawRoleAssignmentImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawRoleAssignmentImpl struct {
	roleAssignment BaseRoleAssignmentImpl
	Type           string
	Values         map[string]interface{}
}

func (s RawRoleAssignmentImpl) RoleAssignment() BaseRoleAssignmentImpl {
	return s.roleAssignment
}

func (s RawRoleAssignmentImpl) Entity() BaseEntityImpl {
	return s.roleAssignment.Entity()
}

var _ json.Marshaler = BaseRoleAssignmentImpl{}

func (s BaseRoleAssignmentImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseRoleAssignmentImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseRoleAssignmentImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseRoleAssignmentImpl: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.roleAssignment"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseRoleAssignmentImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseRoleAssignmentImpl{}

func (s *BaseRoleAssignmentImpl) UnmarshalJSON(bytes []byte) error {
	type alias BaseRoleAssignmentImpl
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into BaseRoleAssignmentImpl: %+v", err)
	}

	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.ResourceScopes = decoded.ResourceScopes
	s.ScopeMembers = decoded.ScopeMembers
	s.ScopeType = decoded.ScopeType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseRoleAssignmentImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["roleDefinition"]; ok {
		impl, err := UnmarshalRoleDefinitionImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'RoleDefinition' for 'BaseRoleAssignmentImpl': %+v", err)
		}
		s.RoleDefinition = &impl
	}
	return nil
}

func UnmarshalRoleAssignmentImplementation(input []byte) (RoleAssignment, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling RoleAssignment into map[string]interface: %+v", err)
	}

	value, ok := temp["@odata.type"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceAndAppManagementRoleAssignment") {
		var out DeviceAndAppManagementRoleAssignment
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceAndAppManagementRoleAssignment: %+v", err)
		}
		return out, nil
	}

	var parent BaseRoleAssignmentImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseRoleAssignmentImpl: %+v", err)
	}

	return RawRoleAssignmentImpl{
		roleAssignment: parent,
		Type:           value,
		Values:         temp,
	}, nil

}
