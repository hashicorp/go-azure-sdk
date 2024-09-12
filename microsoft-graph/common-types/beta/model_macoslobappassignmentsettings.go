package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MobileAppAssignmentSettings = MacOsLobAppAssignmentSettings{}

type MacOsLobAppAssignmentSettings struct {
	// Whether or not to uninstall the app when device is removed from Intune.
	UninstallOnDeviceRemoval nullable.Type[bool] `json:"uninstallOnDeviceRemoval,omitempty"`

	// Fields inherited from MobileAppAssignmentSettings

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s MacOsLobAppAssignmentSettings) MobileAppAssignmentSettings() BaseMobileAppAssignmentSettingsImpl {
	return BaseMobileAppAssignmentSettingsImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MacOsLobAppAssignmentSettings{}

func (s MacOsLobAppAssignmentSettings) MarshalJSON() ([]byte, error) {
	type wrapper MacOsLobAppAssignmentSettings
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MacOsLobAppAssignmentSettings: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MacOsLobAppAssignmentSettings: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.macOsLobAppAssignmentSettings"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MacOsLobAppAssignmentSettings: %+v", err)
	}

	return encoded, nil
}
