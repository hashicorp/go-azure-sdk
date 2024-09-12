package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ManagedTenantsManagedTenantTicketingEndpoint{}

type ManagedTenantsManagedTenantTicketingEndpoint struct {
	CreatedByUserId    nullable.Type[string] `json:"createdByUserId,omitempty"`
	CreatedDateTime    nullable.Type[string] `json:"createdDateTime,omitempty"`
	DisplayName        nullable.Type[string] `json:"displayName,omitempty"`
	EmailAddress       nullable.Type[string] `json:"emailAddress,omitempty"`
	LastActionByUserId nullable.Type[string] `json:"lastActionByUserId,omitempty"`
	LastActionDateTime nullable.Type[string] `json:"lastActionDateTime,omitempty"`
	PhoneNumber        nullable.Type[string] `json:"phoneNumber,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s ManagedTenantsManagedTenantTicketingEndpoint) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ManagedTenantsManagedTenantTicketingEndpoint{}

func (s ManagedTenantsManagedTenantTicketingEndpoint) MarshalJSON() ([]byte, error) {
	type wrapper ManagedTenantsManagedTenantTicketingEndpoint
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedTenantsManagedTenantTicketingEndpoint: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedTenantsManagedTenantTicketingEndpoint: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.managedTenants.managedTenantTicketingEndpoint"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedTenantsManagedTenantTicketingEndpoint: %+v", err)
	}

	return encoded, nil
}
