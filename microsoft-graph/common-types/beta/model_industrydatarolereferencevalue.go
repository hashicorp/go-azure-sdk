package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IndustryDataReferenceValue = IndustryDataRoleReferenceValue{}

type IndustryDataRoleReferenceValue struct {

	// Fields inherited from IndustryDataReferenceValue

	// The code of the desired referenceDefinition entry.
	Code nullable.Type[string] `json:"code,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Value *IndustryDataReferenceDefinition `json:"value,omitempty"`
}

func (s IndustryDataRoleReferenceValue) IndustryDataReferenceValue() BaseIndustryDataReferenceValueImpl {
	return BaseIndustryDataReferenceValueImpl{
		Code:      s.Code,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
		Value:     s.Value,
	}
}

var _ json.Marshaler = IndustryDataRoleReferenceValue{}

func (s IndustryDataRoleReferenceValue) MarshalJSON() ([]byte, error) {
	type wrapper IndustryDataRoleReferenceValue
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IndustryDataRoleReferenceValue: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IndustryDataRoleReferenceValue: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.industryData.roleReferenceValue"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IndustryDataRoleReferenceValue: %+v", err)
	}

	return encoded, nil
}
