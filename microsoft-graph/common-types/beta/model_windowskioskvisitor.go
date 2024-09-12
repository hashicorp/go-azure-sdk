package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WindowsKioskUser = WindowsKioskVisitor{}

type WindowsKioskVisitor struct {

	// Fields inherited from WindowsKioskUser

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s WindowsKioskVisitor) WindowsKioskUser() BaseWindowsKioskUserImpl {
	return BaseWindowsKioskUserImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsKioskVisitor{}

func (s WindowsKioskVisitor) MarshalJSON() ([]byte, error) {
	type wrapper WindowsKioskVisitor
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsKioskVisitor: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsKioskVisitor: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.windowsKioskVisitor"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsKioskVisitor: %+v", err)
	}

	return encoded, nil
}
