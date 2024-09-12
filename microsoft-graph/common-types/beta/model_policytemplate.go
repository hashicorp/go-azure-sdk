package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = PolicyTemplate{}

type PolicyTemplate struct {
	// Defines an optional cross-tenant access policy template with user synchronization settings for a multi-tenant
	// organization.
	MultiTenantOrganizationIdentitySynchronization *MultiTenantOrganizationIdentitySyncPolicyTemplate `json:"multiTenantOrganizationIdentitySynchronization,omitempty"`

	// Defines an optional cross-tenant access policy template with inbound and outbound partner configuration settings for
	// a multi-tenant organization.
	MultiTenantOrganizationPartnerConfiguration *MultiTenantOrganizationPartnerConfigurationTemplate `json:"multiTenantOrganizationPartnerConfiguration,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s PolicyTemplate) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PolicyTemplate{}

func (s PolicyTemplate) MarshalJSON() ([]byte, error) {
	type wrapper PolicyTemplate
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PolicyTemplate: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PolicyTemplate: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.policyTemplate"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PolicyTemplate: %+v", err)
	}

	return encoded, nil
}
