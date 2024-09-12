package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenoteEntitySchemaObjectModel interface {
	Entity
	OnenoteEntityBaseModel
	OnenoteEntitySchemaObjectModel() BaseOnenoteEntitySchemaObjectModelImpl
}

var _ OnenoteEntitySchemaObjectModel = BaseOnenoteEntitySchemaObjectModelImpl{}

type BaseOnenoteEntitySchemaObjectModelImpl struct {
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Fields inherited from OnenoteEntityBaseModel

	Self nullable.Type[string] `json:"self,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s BaseOnenoteEntitySchemaObjectModelImpl) OnenoteEntitySchemaObjectModel() BaseOnenoteEntitySchemaObjectModelImpl {
	return s
}

func (s BaseOnenoteEntitySchemaObjectModelImpl) OnenoteEntityBaseModel() BaseOnenoteEntityBaseModelImpl {
	return BaseOnenoteEntityBaseModelImpl{
		Self:      s.Self,
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s BaseOnenoteEntitySchemaObjectModelImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ OnenoteEntitySchemaObjectModel = RawOnenoteEntitySchemaObjectModelImpl{}

// RawOnenoteEntitySchemaObjectModelImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawOnenoteEntitySchemaObjectModelImpl struct {
	onenoteEntitySchemaObjectModel BaseOnenoteEntitySchemaObjectModelImpl
	Type                           string
	Values                         map[string]interface{}
}

func (s RawOnenoteEntitySchemaObjectModelImpl) OnenoteEntitySchemaObjectModel() BaseOnenoteEntitySchemaObjectModelImpl {
	return s.onenoteEntitySchemaObjectModel
}

func (s RawOnenoteEntitySchemaObjectModelImpl) OnenoteEntityBaseModel() BaseOnenoteEntityBaseModelImpl {
	return s.onenoteEntitySchemaObjectModel.OnenoteEntityBaseModel()
}

func (s RawOnenoteEntitySchemaObjectModelImpl) Entity() BaseEntityImpl {
	return s.onenoteEntitySchemaObjectModel.Entity()
}

var _ json.Marshaler = BaseOnenoteEntitySchemaObjectModelImpl{}

func (s BaseOnenoteEntitySchemaObjectModelImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseOnenoteEntitySchemaObjectModelImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseOnenoteEntitySchemaObjectModelImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseOnenoteEntitySchemaObjectModelImpl: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.onenoteEntitySchemaObjectModel"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseOnenoteEntitySchemaObjectModelImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalOnenoteEntitySchemaObjectModelImplementation(input []byte) (OnenoteEntitySchemaObjectModel, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling OnenoteEntitySchemaObjectModel into map[string]interface: %+v", err)
	}

	value, ok := temp["@odata.type"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.onenoteEntityHierarchyModel") {
		var out OnenoteEntityHierarchyModel
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnenoteEntityHierarchyModel: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.onenotePage") {
		var out OnenotePage
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnenotePage: %+v", err)
		}
		return out, nil
	}

	var parent BaseOnenoteEntitySchemaObjectModelImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseOnenoteEntitySchemaObjectModelImpl: %+v", err)
	}

	return RawOnenoteEntitySchemaObjectModelImpl{
		onenoteEntitySchemaObjectModel: parent,
		Type:                           value,
		Values:                         temp,
	}, nil

}
