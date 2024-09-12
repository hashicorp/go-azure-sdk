package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UserExperienceAnalyticsDeviceWithoutCloudIdentity{}

type UserExperienceAnalyticsDeviceWithoutCloudIdentity struct {
	// Azure Active Directory Device Id
	AzureAdDeviceId nullable.Type[string] `json:"azureAdDeviceId,omitempty"`

	// The tenant attach device's name.
	DeviceName nullable.Type[string] `json:"deviceName,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s UserExperienceAnalyticsDeviceWithoutCloudIdentity) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserExperienceAnalyticsDeviceWithoutCloudIdentity{}

func (s UserExperienceAnalyticsDeviceWithoutCloudIdentity) MarshalJSON() ([]byte, error) {
	type wrapper UserExperienceAnalyticsDeviceWithoutCloudIdentity
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserExperienceAnalyticsDeviceWithoutCloudIdentity: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserExperienceAnalyticsDeviceWithoutCloudIdentity: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.userExperienceAnalyticsDeviceWithoutCloudIdentity"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserExperienceAnalyticsDeviceWithoutCloudIdentity: %+v", err)
	}

	return encoded, nil
}
