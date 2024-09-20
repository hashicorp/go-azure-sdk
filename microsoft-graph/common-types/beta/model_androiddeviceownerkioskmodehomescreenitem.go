package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerKioskModeHomeScreenItem interface {
	AndroidDeviceOwnerKioskModeHomeScreenItem() BaseAndroidDeviceOwnerKioskModeHomeScreenItemImpl
}

var _ AndroidDeviceOwnerKioskModeHomeScreenItem = BaseAndroidDeviceOwnerKioskModeHomeScreenItemImpl{}

type BaseAndroidDeviceOwnerKioskModeHomeScreenItemImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s BaseAndroidDeviceOwnerKioskModeHomeScreenItemImpl) AndroidDeviceOwnerKioskModeHomeScreenItem() BaseAndroidDeviceOwnerKioskModeHomeScreenItemImpl {
	return s
}

var _ AndroidDeviceOwnerKioskModeHomeScreenItem = RawAndroidDeviceOwnerKioskModeHomeScreenItemImpl{}

// RawAndroidDeviceOwnerKioskModeHomeScreenItemImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAndroidDeviceOwnerKioskModeHomeScreenItemImpl struct {
	androidDeviceOwnerKioskModeHomeScreenItem BaseAndroidDeviceOwnerKioskModeHomeScreenItemImpl
	Type                                      string
	Values                                    map[string]interface{}
}

func (s RawAndroidDeviceOwnerKioskModeHomeScreenItemImpl) AndroidDeviceOwnerKioskModeHomeScreenItem() BaseAndroidDeviceOwnerKioskModeHomeScreenItemImpl {
	return s.androidDeviceOwnerKioskModeHomeScreenItem
}

func UnmarshalAndroidDeviceOwnerKioskModeHomeScreenItemImplementation(input []byte) (AndroidDeviceOwnerKioskModeHomeScreenItem, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AndroidDeviceOwnerKioskModeHomeScreenItem into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.androidDeviceOwnerKioskModeFolderItem") {
		var out AndroidDeviceOwnerKioskModeFolderItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidDeviceOwnerKioskModeFolderItem: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidDeviceOwnerKioskModeManagedFolderReference") {
		var out AndroidDeviceOwnerKioskModeManagedFolderReference
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidDeviceOwnerKioskModeManagedFolderReference: %+v", err)
		}
		return out, nil
	}

	var parent BaseAndroidDeviceOwnerKioskModeHomeScreenItemImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAndroidDeviceOwnerKioskModeHomeScreenItemImpl: %+v", err)
	}

	return RawAndroidDeviceOwnerKioskModeHomeScreenItemImpl{
		androidDeviceOwnerKioskModeHomeScreenItem: parent,
		Type:   value,
		Values: temp,
	}, nil

}
