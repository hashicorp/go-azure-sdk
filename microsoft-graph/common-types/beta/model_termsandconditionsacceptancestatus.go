package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = TermsAndConditionsAcceptanceStatus{}

type TermsAndConditionsAcceptanceStatus struct {
	// DateTime when the terms were last accepted by the user.
	AcceptedDateTime *string `json:"acceptedDateTime,omitempty"`

	// Most recent version number of the T&C accepted by the user.
	AcceptedVersion *int64 `json:"acceptedVersion,omitempty"`

	// Navigation link to the terms and conditions that are assigned.
	TermsAndConditions *TermsAndConditions `json:"termsAndConditions,omitempty"`

	// Display name of the user whose acceptance the entity represents.
	UserDisplayName nullable.Type[string] `json:"userDisplayName,omitempty"`

	// The userPrincipalName of the User that accepted the term.
	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s TermsAndConditionsAcceptanceStatus) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TermsAndConditionsAcceptanceStatus{}

func (s TermsAndConditionsAcceptanceStatus) MarshalJSON() ([]byte, error) {
	type wrapper TermsAndConditionsAcceptanceStatus
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TermsAndConditionsAcceptanceStatus: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TermsAndConditionsAcceptanceStatus: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.termsAndConditionsAcceptanceStatus"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TermsAndConditionsAcceptanceStatus: %+v", err)
	}

	return encoded, nil
}
