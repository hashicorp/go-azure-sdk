package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UnifiedRoleAssignment{}

type UnifiedRoleAssignment struct {
	// Read-only property with details of the app specific scope when the assignment scope is app specific. Containment
	// entity. Supports $expand for the entitlement provider only.
	AppScope *AppScope `json:"appScope,omitempty"`

	// Identifier of the app specific scope when the assignment scope is app specific. The scope of an assignment determines
	// the set of resources for which the principal has been granted access. App scopes are scopes that are defined and
	// understood by a resource application only. For the entitlement management provider, use this property to specify a
	// catalog. For example, /AccessPackageCatalog/beedadfe-01d5-4025-910b-84abb9369997. Supports $filter (eq, in). For
	// example, /roleManagement/entitlementManagement/roleAssignments?$filter=appScopeId eq '/AccessPackageCatalog/{catalog
	// id}'.
	AppScopeId nullable.Type[string] `json:"appScopeId,omitempty"`

	Condition nullable.Type[string] `json:"condition,omitempty"`

	// The directory object that is the scope of the assignment. Read-only. Supports $expand for the directory provider.
	DirectoryScope *DirectoryObject `json:"directoryScope,omitempty"`

	// Identifier of the directory object representing the scope of the assignment. The scope of an assignment determines
	// the set of resources for which the principal has been granted access. Directory scopes are shared scopes stored in
	// the directory that are understood by multiple applications, unlike app scopes that are defined and understood by a
	// resource application only. Supports $filter (eq, in).
	DirectoryScopeId nullable.Type[string] `json:"directoryScopeId,omitempty"`

	// OData ID for `DirectoryScope` to bind to this entity
	DirectoryScope_ODataBind *string `json:"directoryScope@odata.bind,omitempty"`

	// Referencing the assigned principal. Read-only. Supports $expand except for the Exchange provider.
	Principal *DirectoryObject `json:"principal,omitempty"`

	// Identifier of the principal to which the assignment is granted. Supported principals are users, role-assignable
	// groups, and service principals. Supports $filter (eq, in).
	PrincipalId nullable.Type[string] `json:"principalId,omitempty"`

	// OData ID for `Principal` to bind to this entity
	Principal_ODataBind *string `json:"principal@odata.bind,omitempty"`

	// The roleDefinition the assignment is for. Supports $expand.
	RoleDefinition *UnifiedRoleDefinition `json:"roleDefinition,omitempty"`

	// Identifier of the unifiedRoleDefinition the assignment is for. Read-only. Supports $filter (eq, in).
	RoleDefinitionId nullable.Type[string] `json:"roleDefinitionId,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s UnifiedRoleAssignment) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UnifiedRoleAssignment{}

func (s UnifiedRoleAssignment) MarshalJSON() ([]byte, error) {
	type wrapper UnifiedRoleAssignment
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UnifiedRoleAssignment: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UnifiedRoleAssignment: %+v", err)
	}

	delete(decoded, "directoryScope")
	delete(decoded, "principal")
	delete(decoded, "roleDefinitionId")
	decoded["@odata.type"] = "#microsoft.graph.unifiedRoleAssignment"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UnifiedRoleAssignment: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &UnifiedRoleAssignment{}

func (s *UnifiedRoleAssignment) UnmarshalJSON(bytes []byte) error {

	var decoded struct {
		AppScope                 *AppScope              `json:"appScope,omitempty"`
		AppScopeId               nullable.Type[string]  `json:"appScopeId,omitempty"`
		Condition                nullable.Type[string]  `json:"condition,omitempty"`
		DirectoryScopeId         nullable.Type[string]  `json:"directoryScopeId,omitempty"`
		DirectoryScope_ODataBind *string                `json:"directoryScope@odata.bind,omitempty"`
		PrincipalId              nullable.Type[string]  `json:"principalId,omitempty"`
		Principal_ODataBind      *string                `json:"principal@odata.bind,omitempty"`
		RoleDefinition           *UnifiedRoleDefinition `json:"roleDefinition,omitempty"`
		RoleDefinitionId         nullable.Type[string]  `json:"roleDefinitionId,omitempty"`
		Id                       *string                `json:"id,omitempty"`
		ODataId                  *string                `json:"@odata.id,omitempty"`
		ODataType                *string                `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AppScope = decoded.AppScope
	s.AppScopeId = decoded.AppScopeId
	s.Condition = decoded.Condition
	s.DirectoryScopeId = decoded.DirectoryScopeId
	s.DirectoryScope_ODataBind = decoded.DirectoryScope_ODataBind
	s.PrincipalId = decoded.PrincipalId
	s.Principal_ODataBind = decoded.Principal_ODataBind
	s.RoleDefinition = decoded.RoleDefinition
	s.RoleDefinitionId = decoded.RoleDefinitionId
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling UnifiedRoleAssignment into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["directoryScope"]; ok {
		impl, err := UnmarshalDirectoryObjectImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'DirectoryScope' for 'UnifiedRoleAssignment': %+v", err)
		}
		s.DirectoryScope = &impl
	}

	if v, ok := temp["principal"]; ok {
		impl, err := UnmarshalDirectoryObjectImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Principal' for 'UnifiedRoleAssignment': %+v", err)
		}
		s.Principal = &impl
	}

	return nil
}
