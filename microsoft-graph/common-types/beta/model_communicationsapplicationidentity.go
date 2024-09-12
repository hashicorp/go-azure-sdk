package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Identity = CommunicationsApplicationIdentity{}

type CommunicationsApplicationIdentity struct {
	// First-party Microsoft application that presents this identity.
	ApplicationType nullable.Type[string] `json:"applicationType,omitempty"`

	// True if the participant shouldn't be shown in other participants' rosters.
	Hidden nullable.Type[bool] `json:"hidden,omitempty"`

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

func (s CommunicationsApplicationIdentity) Identity() BaseIdentityImpl {
	return BaseIdentityImpl{
		DisplayName: s.DisplayName,
		Id:          s.Id,
		ODataId:     s.ODataId,
		ODataType:   s.ODataType,
	}
}

var _ json.Marshaler = CommunicationsApplicationIdentity{}

func (s CommunicationsApplicationIdentity) MarshalJSON() ([]byte, error) {
	type wrapper CommunicationsApplicationIdentity
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CommunicationsApplicationIdentity: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CommunicationsApplicationIdentity: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.communicationsApplicationIdentity"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CommunicationsApplicationIdentity: %+v", err)
	}

	return encoded, nil
}
