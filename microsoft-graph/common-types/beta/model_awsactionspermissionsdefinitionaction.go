package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AwsPermissionsDefinitionAction = AwsActionsPermissionsDefinitionAction{}

type AwsActionsPermissionsDefinitionAction struct {
	// Defines AWS statements.
	AssignToRoleId nullable.Type[string] `json:"assignToRoleId,omitempty"`

	Statements *[]AwsStatement `json:"statements,omitempty"`

	// Fields inherited from PermissionsDefinitionAction

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s AwsActionsPermissionsDefinitionAction) AwsPermissionsDefinitionAction() BaseAwsPermissionsDefinitionActionImpl {
	return BaseAwsPermissionsDefinitionActionImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s AwsActionsPermissionsDefinitionAction) PermissionsDefinitionAction() BasePermissionsDefinitionActionImpl {
	return BasePermissionsDefinitionActionImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AwsActionsPermissionsDefinitionAction{}

func (s AwsActionsPermissionsDefinitionAction) MarshalJSON() ([]byte, error) {
	type wrapper AwsActionsPermissionsDefinitionAction
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AwsActionsPermissionsDefinitionAction: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AwsActionsPermissionsDefinitionAction: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.awsActionsPermissionsDefinitionAction"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AwsActionsPermissionsDefinitionAction: %+v", err)
	}

	return encoded, nil
}
