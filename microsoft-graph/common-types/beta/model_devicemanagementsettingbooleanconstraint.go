package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceManagementConstraint = DeviceManagementSettingBooleanConstraint{}

type DeviceManagementSettingBooleanConstraint struct {
	// The boolean value to compare against
	Value *bool `json:"value,omitempty"`

	// Fields inherited from DeviceManagementConstraint

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s DeviceManagementSettingBooleanConstraint) DeviceManagementConstraint() BaseDeviceManagementConstraintImpl {
	return BaseDeviceManagementConstraintImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceManagementSettingBooleanConstraint{}

func (s DeviceManagementSettingBooleanConstraint) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementSettingBooleanConstraint
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementSettingBooleanConstraint: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementSettingBooleanConstraint: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.deviceManagementSettingBooleanConstraint"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementSettingBooleanConstraint: %+v", err)
	}

	return encoded, nil
}
