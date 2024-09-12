package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WorkbookRangeFont{}

type WorkbookRangeFont struct {
	// Represents the bold status of font.
	Bold nullable.Type[bool] `json:"bold,omitempty"`

	// HTML color code representation of the text color. for example #FF0000 represents Red.
	Color nullable.Type[string] `json:"color,omitempty"`

	// Represents the italic status of the font.
	Italic nullable.Type[bool] `json:"italic,omitempty"`

	// Font name (for example 'Calibri')
	Name nullable.Type[string] `json:"name,omitempty"`

	// Type of underline applied to the font. The possible values are: None, Single, Double, SingleAccountant,
	// DoubleAccountant.
	Underline nullable.Type[string] `json:"underline,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s WorkbookRangeFont) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WorkbookRangeFont{}

func (s WorkbookRangeFont) MarshalJSON() ([]byte, error) {
	type wrapper WorkbookRangeFont
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WorkbookRangeFont: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WorkbookRangeFont: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.workbookRangeFont"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WorkbookRangeFont: %+v", err)
	}

	return encoded, nil
}
