package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WindowsUpdatesUpdatableAssetError = WindowsUpdatesAzureADDeviceRegistrationError{}

type WindowsUpdatesAzureADDeviceRegistrationError struct {
	Reason *WindowsUpdatesAzureADDeviceRegistrationErrorReason `json:"reason,omitempty"`

	// Fields inherited from WindowsUpdatesUpdatableAssetError

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s WindowsUpdatesAzureADDeviceRegistrationError) WindowsUpdatesUpdatableAssetError() BaseWindowsUpdatesUpdatableAssetErrorImpl {
	return BaseWindowsUpdatesUpdatableAssetErrorImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsUpdatesAzureADDeviceRegistrationError{}

func (s WindowsUpdatesAzureADDeviceRegistrationError) MarshalJSON() ([]byte, error) {
	type wrapper WindowsUpdatesAzureADDeviceRegistrationError
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsUpdatesAzureADDeviceRegistrationError: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdatesAzureADDeviceRegistrationError: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.windowsUpdates.azureADDeviceRegistrationError"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsUpdatesAzureADDeviceRegistrationError: %+v", err)
	}

	return encoded, nil
}
