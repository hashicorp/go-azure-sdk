package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ManagedTenantsManagedTenantAlertRuleDefinition{}

type ManagedTenantsManagedTenantAlertRuleDefinition struct {
	AlertRules         *[]ManagedTenantsManagedTenantAlertRule    `json:"alertRules,omitempty"`
	CreatedByUserId    nullable.Type[string]                      `json:"createdByUserId,omitempty"`
	CreatedDateTime    nullable.Type[string]                      `json:"createdDateTime,omitempty"`
	DefinitionTemplate *ManagedTenantsAlertRuleDefinitionTemplate `json:"definitionTemplate,omitempty"`
	DisplayName        nullable.Type[string]                      `json:"displayName,omitempty"`
	LastActionByUserId nullable.Type[string]                      `json:"lastActionByUserId,omitempty"`
	LastActionDateTime nullable.Type[string]                      `json:"lastActionDateTime,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s ManagedTenantsManagedTenantAlertRuleDefinition) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ManagedTenantsManagedTenantAlertRuleDefinition{}

func (s ManagedTenantsManagedTenantAlertRuleDefinition) MarshalJSON() ([]byte, error) {
	type wrapper ManagedTenantsManagedTenantAlertRuleDefinition
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedTenantsManagedTenantAlertRuleDefinition: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedTenantsManagedTenantAlertRuleDefinition: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.managedTenants.managedTenantAlertRuleDefinition"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedTenantsManagedTenantAlertRuleDefinition: %+v", err)
	}

	return encoded, nil
}
