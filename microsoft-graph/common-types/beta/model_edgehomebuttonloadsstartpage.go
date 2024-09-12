package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EdgeHomeButtonConfiguration = EdgeHomeButtonLoadsStartPage{}

type EdgeHomeButtonLoadsStartPage struct {

	// Fields inherited from EdgeHomeButtonConfiguration

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s EdgeHomeButtonLoadsStartPage) EdgeHomeButtonConfiguration() BaseEdgeHomeButtonConfigurationImpl {
	return BaseEdgeHomeButtonConfigurationImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EdgeHomeButtonLoadsStartPage{}

func (s EdgeHomeButtonLoadsStartPage) MarshalJSON() ([]byte, error) {
	type wrapper EdgeHomeButtonLoadsStartPage
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EdgeHomeButtonLoadsStartPage: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EdgeHomeButtonLoadsStartPage: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.edgeHomeButtonLoadsStartPage"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EdgeHomeButtonLoadsStartPage: %+v", err)
	}

	return encoded, nil
}
