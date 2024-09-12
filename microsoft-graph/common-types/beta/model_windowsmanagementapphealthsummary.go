package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WindowsManagementAppHealthSummary{}

type WindowsManagementAppHealthSummary struct {
	// Healthy device count.
	HealthyDeviceCount *int64 `json:"healthyDeviceCount,omitempty"`

	// Unhealthy device count.
	UnhealthyDeviceCount *int64 `json:"unhealthyDeviceCount,omitempty"`

	// Unknown device count.
	UnknownDeviceCount *int64 `json:"unknownDeviceCount,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s WindowsManagementAppHealthSummary) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsManagementAppHealthSummary{}

func (s WindowsManagementAppHealthSummary) MarshalJSON() ([]byte, error) {
	type wrapper WindowsManagementAppHealthSummary
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsManagementAppHealthSummary: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsManagementAppHealthSummary: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.windowsManagementAppHealthSummary"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsManagementAppHealthSummary: %+v", err)
	}

	return encoded, nil
}
