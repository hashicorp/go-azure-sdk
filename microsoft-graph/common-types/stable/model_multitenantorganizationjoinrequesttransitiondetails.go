package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MultiTenantOrganizationJoinRequestTransitionDetails struct {
	// State of the tenant in the multitenant organization currently being processed. The possible values are: pending,
	// active, removed, unknownFutureValue. Read-only.
	DesiredMemberState *MultiTenantOrganizationMemberState `json:"desiredMemberState,omitempty"`

	// Details that explain the processing status if any. Read-only.
	Details nullable.Type[string] `json:"details,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Processing state of the asynchronous job. The possible values are: notStarted, running, succeeded, failed,
	// unknownFutureValue. Read-only.
	Status *MultiTenantOrganizationMemberProcessingStatus `json:"status,omitempty"`
}

var _ json.Marshaler = MultiTenantOrganizationJoinRequestTransitionDetails{}

func (s MultiTenantOrganizationJoinRequestTransitionDetails) MarshalJSON() ([]byte, error) {
	type wrapper MultiTenantOrganizationJoinRequestTransitionDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MultiTenantOrganizationJoinRequestTransitionDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MultiTenantOrganizationJoinRequestTransitionDetails: %+v", err)
	}

	delete(decoded, "desiredMemberState")
	delete(decoded, "details")
	delete(decoded, "status")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MultiTenantOrganizationJoinRequestTransitionDetails: %+v", err)
	}

	return encoded, nil
}
