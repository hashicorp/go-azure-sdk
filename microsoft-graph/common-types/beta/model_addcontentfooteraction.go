package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ InformationProtectionAction = AddContentFooterAction{}

type AddContentFooterAction struct {
	Alignment *ContentAlignment `json:"alignment,omitempty"`

	// Color of the font to use for the footer.
	FontColor nullable.Type[string] `json:"fontColor,omitempty"`

	// Name of the font to use for the footer.
	FontName nullable.Type[string] `json:"fontName,omitempty"`

	// Font size to use for the footer.
	FontSize *int64 `json:"fontSize,omitempty"`

	// The margin of the header from the bottom of the document.
	Margin *int64 `json:"margin,omitempty"`

	// The contents of the footer itself.
	Text nullable.Type[string] `json:"text,omitempty"`

	// The name of the UI element where the footer should be placed.
	UiElementName nullable.Type[string] `json:"uiElementName,omitempty"`

	// Fields inherited from InformationProtectionAction

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s AddContentFooterAction) InformationProtectionAction() BaseInformationProtectionActionImpl {
	return BaseInformationProtectionActionImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AddContentFooterAction{}

func (s AddContentFooterAction) MarshalJSON() ([]byte, error) {
	type wrapper AddContentFooterAction
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AddContentFooterAction: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AddContentFooterAction: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.addContentFooterAction"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AddContentFooterAction: %+v", err)
	}

	return encoded, nil
}
