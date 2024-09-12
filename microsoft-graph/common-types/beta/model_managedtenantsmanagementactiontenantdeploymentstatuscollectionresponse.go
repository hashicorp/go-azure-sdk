package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ BaseCollectionPaginationCountResponse = ManagedTenantsManagementActionTenantDeploymentStatusCollectionResponse{}

type ManagedTenantsManagementActionTenantDeploymentStatusCollectionResponse struct {
	Value *[]ManagedTenantsManagementActionTenantDeploymentStatus `json:"value,omitempty"`

	// Fields inherited from BaseCollectionPaginationCountResponse

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	ODataNextLink nullable.Type[string] `json:"@odata.nextLink,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s ManagedTenantsManagementActionTenantDeploymentStatusCollectionResponse) BaseCollectionPaginationCountResponse() BaseBaseCollectionPaginationCountResponseImpl {
	return BaseBaseCollectionPaginationCountResponseImpl{
		ODataId:       s.ODataId,
		ODataNextLink: s.ODataNextLink,
		ODataType:     s.ODataType,
	}
}

var _ json.Marshaler = ManagedTenantsManagementActionTenantDeploymentStatusCollectionResponse{}

func (s ManagedTenantsManagementActionTenantDeploymentStatusCollectionResponse) MarshalJSON() ([]byte, error) {
	type wrapper ManagedTenantsManagementActionTenantDeploymentStatusCollectionResponse
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedTenantsManagementActionTenantDeploymentStatusCollectionResponse: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedTenantsManagementActionTenantDeploymentStatusCollectionResponse: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.managedTenants.managementActionTenantDeploymentStatusCollectionResponse"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedTenantsManagementActionTenantDeploymentStatusCollectionResponse: %+v", err)
	}

	return encoded, nil
}
