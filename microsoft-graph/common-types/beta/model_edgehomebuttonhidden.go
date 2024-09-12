package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EdgeHomeButtonConfiguration = EdgeHomeButtonHidden{}

type EdgeHomeButtonHidden struct {

	// Fields inherited from EdgeHomeButtonConfiguration

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s EdgeHomeButtonHidden) EdgeHomeButtonConfiguration() BaseEdgeHomeButtonConfigurationImpl {
	return BaseEdgeHomeButtonConfigurationImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EdgeHomeButtonHidden{}

func (s EdgeHomeButtonHidden) MarshalJSON() ([]byte, error) {
	type wrapper EdgeHomeButtonHidden
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EdgeHomeButtonHidden: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EdgeHomeButtonHidden: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.edgeHomeButtonHidden"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EdgeHomeButtonHidden: %+v", err)
	}

	return encoded, nil
}
