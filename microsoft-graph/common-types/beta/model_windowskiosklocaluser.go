package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WindowsKioskUser = WindowsKioskLocalUser{}

type WindowsKioskLocalUser struct {
	// The local user that will be locked to this kiosk configuration
	UserName *string `json:"userName,omitempty"`

	// Fields inherited from WindowsKioskUser

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s WindowsKioskLocalUser) WindowsKioskUser() BaseWindowsKioskUserImpl {
	return BaseWindowsKioskUserImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsKioskLocalUser{}

func (s WindowsKioskLocalUser) MarshalJSON() ([]byte, error) {
	type wrapper WindowsKioskLocalUser
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsKioskLocalUser: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsKioskLocalUser: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.windowsKioskLocalUser"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsKioskLocalUser: %+v", err)
	}

	return encoded, nil
}
