package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = IPSecurityProfile{}

type IPSecurityProfile struct {
	ActivityGroupNames  *[]string                  `json:"activityGroupNames,omitempty"`
	Address             nullable.Type[string]      `json:"address,omitempty"`
	AzureSubscriptionId nullable.Type[string]      `json:"azureSubscriptionId,omitempty"`
	AzureTenantId       *string                    `json:"azureTenantId,omitempty"`
	CountHits           nullable.Type[int64]       `json:"countHits,omitempty"`
	CountHosts          nullable.Type[int64]       `json:"countHosts,omitempty"`
	FirstSeenDateTime   nullable.Type[string]      `json:"firstSeenDateTime,omitempty"`
	IPCategories        *[]IPCategory              `json:"ipCategories,omitempty"`
	IPReferenceData     *[]IPReferenceData         `json:"ipReferenceData,omitempty"`
	LastSeenDateTime    nullable.Type[string]      `json:"lastSeenDateTime,omitempty"`
	RiskScore           nullable.Type[string]      `json:"riskScore,omitempty"`
	Tags                *[]string                  `json:"tags,omitempty"`
	VendorInformation   *SecurityVendorInformation `json:"vendorInformation,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s IPSecurityProfile) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IPSecurityProfile{}

func (s IPSecurityProfile) MarshalJSON() ([]byte, error) {
	type wrapper IPSecurityProfile
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IPSecurityProfile: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IPSecurityProfile: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.ipSecurityProfile"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IPSecurityProfile: %+v", err)
	}

	return encoded, nil
}
