package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceCategory{}

type DeviceCategory struct {
	// Optional description for the device category.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Display name for the device category.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Optional role scope tags for the device category.
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s DeviceCategory) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceCategory{}

func (s DeviceCategory) MarshalJSON() ([]byte, error) {
	type wrapper DeviceCategory
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceCategory: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceCategory: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.deviceCategory"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceCategory: %+v", err)
	}

	return encoded, nil
}
