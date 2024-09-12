package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AndroidDeviceOwnerKioskModeHomeScreenItem = AndroidDeviceOwnerKioskModeManagedFolderReference{}

type AndroidDeviceOwnerKioskModeManagedFolderReference struct {
	// Unique identifier for the folder
	FolderIdentifier nullable.Type[string] `json:"folderIdentifier,omitempty"`

	// Name of the folder
	FolderName *string `json:"folderName,omitempty"`

	// Fields inherited from AndroidDeviceOwnerKioskModeHomeScreenItem

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s AndroidDeviceOwnerKioskModeManagedFolderReference) AndroidDeviceOwnerKioskModeHomeScreenItem() BaseAndroidDeviceOwnerKioskModeHomeScreenItemImpl {
	return BaseAndroidDeviceOwnerKioskModeHomeScreenItemImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AndroidDeviceOwnerKioskModeManagedFolderReference{}

func (s AndroidDeviceOwnerKioskModeManagedFolderReference) MarshalJSON() ([]byte, error) {
	type wrapper AndroidDeviceOwnerKioskModeManagedFolderReference
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AndroidDeviceOwnerKioskModeManagedFolderReference: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AndroidDeviceOwnerKioskModeManagedFolderReference: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.androidDeviceOwnerKioskModeManagedFolderReference"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AndroidDeviceOwnerKioskModeManagedFolderReference: %+v", err)
	}

	return encoded, nil
}
