package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = CloudClipboardRoot{}

type CloudClipboardRoot struct {
	// Represents a collection of Cloud Clipboard items.
	Items *[]CloudClipboardItem `json:"items,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s CloudClipboardRoot) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CloudClipboardRoot{}

func (s CloudClipboardRoot) MarshalJSON() ([]byte, error) {
	type wrapper CloudClipboardRoot
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CloudClipboardRoot: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CloudClipboardRoot: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.cloudClipboardRoot"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CloudClipboardRoot: %+v", err)
	}

	return encoded, nil
}
