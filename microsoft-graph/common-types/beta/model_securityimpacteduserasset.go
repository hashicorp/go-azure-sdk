package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityImpactedAsset = SecurityImpactedUserAsset{}

type SecurityImpactedUserAsset struct {
	Identifier *SecurityUserAssetIdentifier `json:"identifier,omitempty"`

	// Fields inherited from SecurityImpactedAsset

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s SecurityImpactedUserAsset) SecurityImpactedAsset() BaseSecurityImpactedAssetImpl {
	return BaseSecurityImpactedAssetImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityImpactedUserAsset{}

func (s SecurityImpactedUserAsset) MarshalJSON() ([]byte, error) {
	type wrapper SecurityImpactedUserAsset
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityImpactedUserAsset: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityImpactedUserAsset: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.security.impactedUserAsset"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityImpactedUserAsset: %+v", err)
	}

	return encoded, nil
}
