package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PermissionsDefinitionIdentitySource = SamlIdentitySource{}

type SamlIdentitySource struct {

	// Fields inherited from PermissionsDefinitionIdentitySource

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s SamlIdentitySource) PermissionsDefinitionIdentitySource() BasePermissionsDefinitionIdentitySourceImpl {
	return BasePermissionsDefinitionIdentitySourceImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SamlIdentitySource{}

func (s SamlIdentitySource) MarshalJSON() ([]byte, error) {
	type wrapper SamlIdentitySource
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SamlIdentitySource: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SamlIdentitySource: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.samlIdentitySource"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SamlIdentitySource: %+v", err)
	}

	return encoded, nil
}
