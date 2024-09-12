package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExactMatchJobBase interface {
	Entity
	ExactMatchJobBase() BaseExactMatchJobBaseImpl
}

var _ ExactMatchJobBase = BaseExactMatchJobBaseImpl{}

type BaseExactMatchJobBaseImpl struct {
	CompletionDateTime  nullable.Type[string] `json:"completionDateTime,omitempty"`
	CreationDateTime    nullable.Type[string] `json:"creationDateTime,omitempty"`
	Error               *ClassificationError  `json:"error,omitempty"`
	LastUpdatedDateTime nullable.Type[string] `json:"lastUpdatedDateTime,omitempty"`
	StartDateTime       nullable.Type[string] `json:"startDateTime,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s BaseExactMatchJobBaseImpl) ExactMatchJobBase() BaseExactMatchJobBaseImpl {
	return s
}

func (s BaseExactMatchJobBaseImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ ExactMatchJobBase = RawExactMatchJobBaseImpl{}

// RawExactMatchJobBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawExactMatchJobBaseImpl struct {
	exactMatchJobBase BaseExactMatchJobBaseImpl
	Type              string
	Values            map[string]interface{}
}

func (s RawExactMatchJobBaseImpl) ExactMatchJobBase() BaseExactMatchJobBaseImpl {
	return s.exactMatchJobBase
}

func (s RawExactMatchJobBaseImpl) Entity() BaseEntityImpl {
	return s.exactMatchJobBase.Entity()
}

var _ json.Marshaler = BaseExactMatchJobBaseImpl{}

func (s BaseExactMatchJobBaseImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseExactMatchJobBaseImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseExactMatchJobBaseImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseExactMatchJobBaseImpl: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.exactMatchJobBase"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseExactMatchJobBaseImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalExactMatchJobBaseImplementation(input []byte) (ExactMatchJobBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ExactMatchJobBase into map[string]interface: %+v", err)
	}

	value, ok := temp["@odata.type"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.exactMatchLookupJob") {
		var out ExactMatchLookupJob
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExactMatchLookupJob: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.exactMatchSessionBase") {
		var out ExactMatchSessionBase
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExactMatchSessionBase: %+v", err)
		}
		return out, nil
	}

	var parent BaseExactMatchJobBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseExactMatchJobBaseImpl: %+v", err)
	}

	return RawExactMatchJobBaseImpl{
		exactMatchJobBase: parent,
		Type:              value,
		Values:            temp,
	}, nil

}
