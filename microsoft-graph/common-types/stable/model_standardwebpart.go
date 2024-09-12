package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WebPart = StandardWebPart{}

type StandardWebPart struct {
	// The instance identifier of the container text webPart. It only works for inline standard webPart in rich text
	// webParts.
	ContainerTextWebPartId nullable.Type[string] `json:"containerTextWebPartId,omitempty"`

	// Data of the webPart.
	Data *WebPartData `json:"data,omitempty"`

	// A Guid that indicates the webPart type.
	WebPartType nullable.Type[string] `json:"webPartType,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s StandardWebPart) WebPart() BaseWebPartImpl {
	return BaseWebPartImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s StandardWebPart) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = StandardWebPart{}

func (s StandardWebPart) MarshalJSON() ([]byte, error) {
	type wrapper StandardWebPart
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling StandardWebPart: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling StandardWebPart: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.standardWebPart"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling StandardWebPart: %+v", err)
	}

	return encoded, nil
}
