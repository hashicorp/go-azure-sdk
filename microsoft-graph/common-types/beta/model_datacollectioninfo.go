package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DataCollectionInfo{}

type DataCollectionInfo struct {
	Entitlements EntitlementsDataCollectionInfo `json:"entitlements"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s DataCollectionInfo) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DataCollectionInfo{}

func (s DataCollectionInfo) MarshalJSON() ([]byte, error) {
	type wrapper DataCollectionInfo
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DataCollectionInfo: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DataCollectionInfo: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.dataCollectionInfo"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DataCollectionInfo: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &DataCollectionInfo{}

func (s *DataCollectionInfo) UnmarshalJSON(bytes []byte) error {

	var decoded struct {
		Id        *string `json:"id,omitempty"`
		ODataId   *string `json:"@odata.id,omitempty"`
		ODataType *string `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DataCollectionInfo into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["entitlements"]; ok {
		impl, err := UnmarshalEntitlementsDataCollectionInfoImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Entitlements' for 'DataCollectionInfo': %+v", err)
		}
		s.Entitlements = impl
	}

	return nil
}
