package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ VirtualEventRegistrationQuestionBase = VirtualEventRegistrationPredefinedQuestion{}

type VirtualEventRegistrationPredefinedQuestion struct {
	Label *VirtualEventRegistrationPredefinedQuestionLabel `json:"label,omitempty"`

	// Fields inherited from VirtualEventRegistrationQuestionBase

	DisplayName nullable.Type[string] `json:"displayName,omitempty"`
	IsRequired  nullable.Type[bool]   `json:"isRequired,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s VirtualEventRegistrationPredefinedQuestion) VirtualEventRegistrationQuestionBase() BaseVirtualEventRegistrationQuestionBaseImpl {
	return BaseVirtualEventRegistrationQuestionBaseImpl{
		DisplayName: s.DisplayName,
		IsRequired:  s.IsRequired,
		Id:          s.Id,
		ODataId:     s.ODataId,
		ODataType:   s.ODataType,
	}
}

func (s VirtualEventRegistrationPredefinedQuestion) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = VirtualEventRegistrationPredefinedQuestion{}

func (s VirtualEventRegistrationPredefinedQuestion) MarshalJSON() ([]byte, error) {
	type wrapper VirtualEventRegistrationPredefinedQuestion
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling VirtualEventRegistrationPredefinedQuestion: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling VirtualEventRegistrationPredefinedQuestion: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.virtualEventRegistrationPredefinedQuestion"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling VirtualEventRegistrationPredefinedQuestion: %+v", err)
	}

	return encoded, nil
}
