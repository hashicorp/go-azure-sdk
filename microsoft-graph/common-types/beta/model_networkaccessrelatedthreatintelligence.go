package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ NetworkaccessRelatedResource = NetworkaccessRelatedThreatIntelligence{}

type NetworkaccessRelatedThreatIntelligence struct {
	ThreatCount *int64 `json:"threatCount,omitempty"`

	// Fields inherited from NetworkaccessRelatedResource

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s NetworkaccessRelatedThreatIntelligence) NetworkaccessRelatedResource() BaseNetworkaccessRelatedResourceImpl {
	return BaseNetworkaccessRelatedResourceImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = NetworkaccessRelatedThreatIntelligence{}

func (s NetworkaccessRelatedThreatIntelligence) MarshalJSON() ([]byte, error) {
	type wrapper NetworkaccessRelatedThreatIntelligence
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling NetworkaccessRelatedThreatIntelligence: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling NetworkaccessRelatedThreatIntelligence: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.networkaccess.relatedThreatIntelligence"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling NetworkaccessRelatedThreatIntelligence: %+v", err)
	}

	return encoded, nil
}
