package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceComplianceLocalActionBase interface {
	Entity
	AndroidDeviceComplianceLocalActionBase() BaseAndroidDeviceComplianceLocalActionBaseImpl
}

var _ AndroidDeviceComplianceLocalActionBase = BaseAndroidDeviceComplianceLocalActionBaseImpl{}

type BaseAndroidDeviceComplianceLocalActionBaseImpl struct {
	// Number of minutes to wait till a local action is enforced. Valid values 0 to 2147483647
	GracePeriodInMinutes *int64 `json:"gracePeriodInMinutes,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s BaseAndroidDeviceComplianceLocalActionBaseImpl) AndroidDeviceComplianceLocalActionBase() BaseAndroidDeviceComplianceLocalActionBaseImpl {
	return s
}

func (s BaseAndroidDeviceComplianceLocalActionBaseImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ AndroidDeviceComplianceLocalActionBase = RawAndroidDeviceComplianceLocalActionBaseImpl{}

// RawAndroidDeviceComplianceLocalActionBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAndroidDeviceComplianceLocalActionBaseImpl struct {
	androidDeviceComplianceLocalActionBase BaseAndroidDeviceComplianceLocalActionBaseImpl
	Type                                   string
	Values                                 map[string]interface{}
}

func (s RawAndroidDeviceComplianceLocalActionBaseImpl) AndroidDeviceComplianceLocalActionBase() BaseAndroidDeviceComplianceLocalActionBaseImpl {
	return s.androidDeviceComplianceLocalActionBase
}

func (s RawAndroidDeviceComplianceLocalActionBaseImpl) Entity() BaseEntityImpl {
	return s.androidDeviceComplianceLocalActionBase.Entity()
}

var _ json.Marshaler = BaseAndroidDeviceComplianceLocalActionBaseImpl{}

func (s BaseAndroidDeviceComplianceLocalActionBaseImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseAndroidDeviceComplianceLocalActionBaseImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseAndroidDeviceComplianceLocalActionBaseImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseAndroidDeviceComplianceLocalActionBaseImpl: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.androidDeviceComplianceLocalActionBase"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseAndroidDeviceComplianceLocalActionBaseImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalAndroidDeviceComplianceLocalActionBaseImplementation(input []byte) (AndroidDeviceComplianceLocalActionBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AndroidDeviceComplianceLocalActionBase into map[string]interface: %+v", err)
	}

	value, ok := temp["@odata.type"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidDeviceComplianceLocalActionLockDevice") {
		var out AndroidDeviceComplianceLocalActionLockDevice
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidDeviceComplianceLocalActionLockDevice: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidDeviceComplianceLocalActionLockDeviceWithPasscode") {
		var out AndroidDeviceComplianceLocalActionLockDeviceWithPasscode
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidDeviceComplianceLocalActionLockDeviceWithPasscode: %+v", err)
		}
		return out, nil
	}

	var parent BaseAndroidDeviceComplianceLocalActionBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAndroidDeviceComplianceLocalActionBaseImpl: %+v", err)
	}

	return RawAndroidDeviceComplianceLocalActionBaseImpl{
		androidDeviceComplianceLocalActionBase: parent,
		Type:                                   value,
		Values:                                 temp,
	}, nil

}
