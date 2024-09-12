package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Identity = ProvisionedIdentity{}

type ProvisionedIdentity struct {
	// Details of the identity.
	Details *DetailsInfo `json:"details,omitempty"`

	// Type of identity that has been provisioned, such as 'user' or 'group.' Supports $filter (eq, contains).
	IdentityType nullable.Type[string] `json:"identityType,omitempty"`

	// Fields inherited from Identity

	// The display name of the identity. This property is read-only.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The identifier of the identity. This property is read-only.
	Id nullable.Type[string] `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s ProvisionedIdentity) Identity() BaseIdentityImpl {
	return BaseIdentityImpl{
		DisplayName: s.DisplayName,
		Id:          s.Id,
		ODataId:     s.ODataId,
		ODataType:   s.ODataType,
	}
}

var _ json.Marshaler = ProvisionedIdentity{}

func (s ProvisionedIdentity) MarshalJSON() ([]byte, error) {
	type wrapper ProvisionedIdentity
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ProvisionedIdentity: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ProvisionedIdentity: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.provisionedIdentity"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ProvisionedIdentity: %+v", err)
	}

	return encoded, nil
}
