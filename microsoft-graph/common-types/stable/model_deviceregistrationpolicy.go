package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceRegistrationPolicy{}

type DeviceRegistrationPolicy struct {
	AzureADJoin                  *AzureADJoinPolicy            `json:"azureADJoin,omitempty"`
	AzureADRegistration          *AzureADRegistrationPolicy    `json:"azureADRegistration,omitempty"`
	Description                  nullable.Type[string]         `json:"description,omitempty"`
	DisplayName                  nullable.Type[string]         `json:"displayName,omitempty"`
	LocalAdminPassword           *LocalAdminPasswordSettings   `json:"localAdminPassword,omitempty"`
	MultiFactorAuthConfiguration *MultiFactorAuthConfiguration `json:"multiFactorAuthConfiguration,omitempty"`
	UserDeviceQuota              *int64                        `json:"userDeviceQuota,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s DeviceRegistrationPolicy) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceRegistrationPolicy{}

func (s DeviceRegistrationPolicy) MarshalJSON() ([]byte, error) {
	type wrapper DeviceRegistrationPolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceRegistrationPolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceRegistrationPolicy: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.deviceRegistrationPolicy"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceRegistrationPolicy: %+v", err)
	}

	return encoded, nil
}
