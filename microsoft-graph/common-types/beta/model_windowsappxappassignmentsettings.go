package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MobileAppAssignmentSettings = WindowsAppXAppAssignmentSettings{}

type WindowsAppXAppAssignmentSettings struct {
	// Whether or not to use device execution context for Windows AppX mobile app.
	UseDeviceContext *bool `json:"useDeviceContext,omitempty"`

	// Fields inherited from MobileAppAssignmentSettings

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s WindowsAppXAppAssignmentSettings) MobileAppAssignmentSettings() BaseMobileAppAssignmentSettingsImpl {
	return BaseMobileAppAssignmentSettingsImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsAppXAppAssignmentSettings{}

func (s WindowsAppXAppAssignmentSettings) MarshalJSON() ([]byte, error) {
	type wrapper WindowsAppXAppAssignmentSettings
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsAppXAppAssignmentSettings: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsAppXAppAssignmentSettings: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.windowsAppXAppAssignmentSettings"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsAppXAppAssignmentSettings: %+v", err)
	}

	return encoded, nil
}
