package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AzurePermissionsDefinitionAction = AzureRolePermissionsDefinitionAction{}

type AzureRolePermissionsDefinitionAction struct {
	Roles *[]PermissionsDefinitionAzureRole `json:"roles,omitempty"`

	// Fields inherited from PermissionsDefinitionAction

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s AzureRolePermissionsDefinitionAction) AzurePermissionsDefinitionAction() BaseAzurePermissionsDefinitionActionImpl {
	return BaseAzurePermissionsDefinitionActionImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s AzureRolePermissionsDefinitionAction) PermissionsDefinitionAction() BasePermissionsDefinitionActionImpl {
	return BasePermissionsDefinitionActionImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AzureRolePermissionsDefinitionAction{}

func (s AzureRolePermissionsDefinitionAction) MarshalJSON() ([]byte, error) {
	type wrapper AzureRolePermissionsDefinitionAction
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureRolePermissionsDefinitionAction: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureRolePermissionsDefinitionAction: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.azureRolePermissionsDefinitionAction"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureRolePermissionsDefinitionAction: %+v", err)
	}

	return encoded, nil
}
