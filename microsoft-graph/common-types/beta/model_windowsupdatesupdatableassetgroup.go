package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WindowsUpdatesUpdatableAsset = WindowsUpdatesUpdatableAssetGroup{}

type WindowsUpdatesUpdatableAssetGroup struct {
	// Members of the group. Read-only.
	Members *[]WindowsUpdatesUpdatableAsset `json:"members,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s WindowsUpdatesUpdatableAssetGroup) WindowsUpdatesUpdatableAsset() BaseWindowsUpdatesUpdatableAssetImpl {
	return BaseWindowsUpdatesUpdatableAssetImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s WindowsUpdatesUpdatableAssetGroup) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsUpdatesUpdatableAssetGroup{}

func (s WindowsUpdatesUpdatableAssetGroup) MarshalJSON() ([]byte, error) {
	type wrapper WindowsUpdatesUpdatableAssetGroup
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsUpdatesUpdatableAssetGroup: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdatesUpdatableAssetGroup: %+v", err)
	}

	delete(decoded, "members")
	decoded["@odata.type"] = "#microsoft.graph.windowsUpdates.updatableAssetGroup"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsUpdatesUpdatableAssetGroup: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &WindowsUpdatesUpdatableAssetGroup{}

func (s *WindowsUpdatesUpdatableAssetGroup) UnmarshalJSON(bytes []byte) error {
	type alias WindowsUpdatesUpdatableAssetGroup
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into WindowsUpdatesUpdatableAssetGroup: %+v", err)
	}

	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling WindowsUpdatesUpdatableAssetGroup into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["members"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Members into list []json.RawMessage: %+v", err)
		}

		output := make([]WindowsUpdatesUpdatableAsset, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalWindowsUpdatesUpdatableAssetImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Members' for 'WindowsUpdatesUpdatableAssetGroup': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Members = &output
	}
	return nil
}
