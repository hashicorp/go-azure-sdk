package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IPRange = IPv4Range{}

type IPv4Range struct {
	// Lower address.
	LowerAddress *string `json:"lowerAddress,omitempty"`

	// Upper address.
	UpperAddress *string `json:"upperAddress,omitempty"`

	// Fields inherited from IPRange

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s IPv4Range) IPRange() BaseIPRangeImpl {
	return BaseIPRangeImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IPv4Range{}

func (s IPv4Range) MarshalJSON() ([]byte, error) {
	type wrapper IPv4Range
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IPv4Range: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IPv4Range: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.iPv4Range"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IPv4Range: %+v", err)
	}

	return encoded, nil
}
