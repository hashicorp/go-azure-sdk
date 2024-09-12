package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ BaseCollectionPaginationCountResponse = TargetedManagedAppConfigurationPolicySetItemCollectionResponse{}

type TargetedManagedAppConfigurationPolicySetItemCollectionResponse struct {
	Value *[]TargetedManagedAppConfigurationPolicySetItem `json:"value,omitempty"`

	// Fields inherited from BaseCollectionPaginationCountResponse

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	ODataNextLink nullable.Type[string] `json:"@odata.nextLink,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s TargetedManagedAppConfigurationPolicySetItemCollectionResponse) BaseCollectionPaginationCountResponse() BaseBaseCollectionPaginationCountResponseImpl {
	return BaseBaseCollectionPaginationCountResponseImpl{
		ODataId:       s.ODataId,
		ODataNextLink: s.ODataNextLink,
		ODataType:     s.ODataType,
	}
}

var _ json.Marshaler = TargetedManagedAppConfigurationPolicySetItemCollectionResponse{}

func (s TargetedManagedAppConfigurationPolicySetItemCollectionResponse) MarshalJSON() ([]byte, error) {
	type wrapper TargetedManagedAppConfigurationPolicySetItemCollectionResponse
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TargetedManagedAppConfigurationPolicySetItemCollectionResponse: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TargetedManagedAppConfigurationPolicySetItemCollectionResponse: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.targetedManagedAppConfigurationPolicySetItemCollectionResponse"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TargetedManagedAppConfigurationPolicySetItemCollectionResponse: %+v", err)
	}

	return encoded, nil
}
