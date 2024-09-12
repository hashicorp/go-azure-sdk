package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ UnifiedRoleManagementAlertConfiguration = RedundantAssignmentAlertConfiguration{}

type RedundantAssignmentAlertConfiguration struct {
	// The number of days without activation to look back on from current timestamp.
	Duration nullable.Type[string] `json:"duration,omitempty"`

	// Fields inherited from UnifiedRoleManagementAlertConfiguration

	// The definition of the alert that contains its description, impact, and measures to mitigate or prevent it. Supports
	// $expand.
	AlertDefinition *UnifiedRoleManagementAlertDefinition `json:"alertDefinition,omitempty"`

	// The identifier of an alert definition. Supports $filter (eq, ne).
	AlertDefinitionId nullable.Type[string] `json:"alertDefinitionId,omitempty"`

	// true if the alert is enabled. Setting it to false disables PIM scanning the tenant to identify instances that trigger
	// the alert.
	IsEnabled nullable.Type[bool] `json:"isEnabled,omitempty"`

	// The identifier of the scope to which the alert is related. Only / is supported to represent the tenant scope.
	// Supports $filter (eq, ne).
	ScopeId nullable.Type[string] `json:"scopeId,omitempty"`

	// The type of scope where the alert is created. DirectoryRole is the only currently supported scope type for Microsoft
	// Entra roles.
	ScopeType nullable.Type[string] `json:"scopeType,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s RedundantAssignmentAlertConfiguration) UnifiedRoleManagementAlertConfiguration() BaseUnifiedRoleManagementAlertConfigurationImpl {
	return BaseUnifiedRoleManagementAlertConfigurationImpl{
		AlertDefinition:   s.AlertDefinition,
		AlertDefinitionId: s.AlertDefinitionId,
		IsEnabled:         s.IsEnabled,
		ScopeId:           s.ScopeId,
		ScopeType:         s.ScopeType,
		Id:                s.Id,
		ODataId:           s.ODataId,
		ODataType:         s.ODataType,
	}
}

func (s RedundantAssignmentAlertConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = RedundantAssignmentAlertConfiguration{}

func (s RedundantAssignmentAlertConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper RedundantAssignmentAlertConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling RedundantAssignmentAlertConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling RedundantAssignmentAlertConfiguration: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.redundantAssignmentAlertConfiguration"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling RedundantAssignmentAlertConfiguration: %+v", err)
	}

	return encoded, nil
}
