package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceManagementTemplateInsightsDefinition{}

type DeviceManagementTemplateInsightsDefinition struct {
	// Setting insights in a template
	SettingInsights *[]DeviceManagementSettingInsightsDefinition `json:"settingInsights,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s DeviceManagementTemplateInsightsDefinition) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceManagementTemplateInsightsDefinition{}

func (s DeviceManagementTemplateInsightsDefinition) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementTemplateInsightsDefinition
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementTemplateInsightsDefinition: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementTemplateInsightsDefinition: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.deviceManagementTemplateInsightsDefinition"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementTemplateInsightsDefinition: %+v", err)
	}

	return encoded, nil
}
