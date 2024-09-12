package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ GroupPolicyPresentationValue = GroupPolicyPresentationValueLongDecimal{}

type GroupPolicyPresentationValueLongDecimal struct {
	// An unsigned long value for the associated presentation.
	Value *int64 `json:"value,omitempty"`

	// Fields inherited from GroupPolicyPresentationValue

	// The date and time the object was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The group policy definition value associated with the presentation value.
	DefinitionValue *GroupPolicyDefinitionValue `json:"definitionValue,omitempty"`

	// The date and time the object was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// The group policy presentation associated with the presentation value.
	Presentation *GroupPolicyPresentation `json:"presentation,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s GroupPolicyPresentationValueLongDecimal) GroupPolicyPresentationValue() BaseGroupPolicyPresentationValueImpl {
	return BaseGroupPolicyPresentationValueImpl{
		CreatedDateTime:      s.CreatedDateTime,
		DefinitionValue:      s.DefinitionValue,
		LastModifiedDateTime: s.LastModifiedDateTime,
		Presentation:         s.Presentation,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s GroupPolicyPresentationValueLongDecimal) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = GroupPolicyPresentationValueLongDecimal{}

func (s GroupPolicyPresentationValueLongDecimal) MarshalJSON() ([]byte, error) {
	type wrapper GroupPolicyPresentationValueLongDecimal
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GroupPolicyPresentationValueLongDecimal: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GroupPolicyPresentationValueLongDecimal: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.groupPolicyPresentationValueLongDecimal"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GroupPolicyPresentationValueLongDecimal: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &GroupPolicyPresentationValueLongDecimal{}

func (s *GroupPolicyPresentationValueLongDecimal) UnmarshalJSON(bytes []byte) error {
	type alias GroupPolicyPresentationValueLongDecimal
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into GroupPolicyPresentationValueLongDecimal: %+v", err)
	}

	s.CreatedDateTime = decoded.CreatedDateTime
	s.DefinitionValue = decoded.DefinitionValue
	s.Id = decoded.Id
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Value = decoded.Value

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling GroupPolicyPresentationValueLongDecimal into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["presentation"]; ok {
		impl, err := UnmarshalGroupPolicyPresentationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Presentation' for 'GroupPolicyPresentationValueLongDecimal': %+v", err)
		}
		s.Presentation = &impl
	}
	return nil
}
