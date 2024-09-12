package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UnifiedRbacResourceAction{}

type UnifiedRbacResourceAction struct {
	ActionVerb                      nullable.Type[string] `json:"actionVerb,omitempty"`
	AuthenticationContextId         nullable.Type[string] `json:"authenticationContextId,omitempty"`
	Description                     nullable.Type[string] `json:"description,omitempty"`
	IsAuthenticationContextSettable nullable.Type[bool]   `json:"isAuthenticationContextSettable,omitempty"`
	Name                            *string               `json:"name,omitempty"`
	ResourceScopeId                 nullable.Type[string] `json:"resourceScopeId,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s UnifiedRbacResourceAction) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UnifiedRbacResourceAction{}

func (s UnifiedRbacResourceAction) MarshalJSON() ([]byte, error) {
	type wrapper UnifiedRbacResourceAction
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UnifiedRbacResourceAction: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UnifiedRbacResourceAction: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.unifiedRbacResourceAction"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UnifiedRbacResourceAction: %+v", err)
	}

	return encoded, nil
}
