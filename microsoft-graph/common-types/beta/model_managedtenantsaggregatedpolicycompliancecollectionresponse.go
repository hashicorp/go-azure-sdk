package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ BaseCollectionPaginationCountResponse = ManagedTenantsAggregatedPolicyComplianceCollectionResponse{}

type ManagedTenantsAggregatedPolicyComplianceCollectionResponse struct {
	Value *[]ManagedTenantsAggregatedPolicyCompliance `json:"value,omitempty"`

	// Fields inherited from BaseCollectionPaginationCountResponse

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	ODataNextLink nullable.Type[string] `json:"@odata.nextLink,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s ManagedTenantsAggregatedPolicyComplianceCollectionResponse) BaseCollectionPaginationCountResponse() BaseBaseCollectionPaginationCountResponseImpl {
	return BaseBaseCollectionPaginationCountResponseImpl{
		ODataId:       s.ODataId,
		ODataNextLink: s.ODataNextLink,
		ODataType:     s.ODataType,
	}
}

var _ json.Marshaler = ManagedTenantsAggregatedPolicyComplianceCollectionResponse{}

func (s ManagedTenantsAggregatedPolicyComplianceCollectionResponse) MarshalJSON() ([]byte, error) {
	type wrapper ManagedTenantsAggregatedPolicyComplianceCollectionResponse
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedTenantsAggregatedPolicyComplianceCollectionResponse: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedTenantsAggregatedPolicyComplianceCollectionResponse: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.managedTenants.aggregatedPolicyComplianceCollectionResponse"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedTenantsAggregatedPolicyComplianceCollectionResponse: %+v", err)
	}

	return encoded, nil
}
