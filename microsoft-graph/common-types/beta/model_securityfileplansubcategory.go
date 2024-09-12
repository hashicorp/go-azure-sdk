package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityFilePlanDescriptorBase = SecurityFilePlanSubcategory{}

type SecurityFilePlanSubcategory struct {

	// Fields inherited from SecurityFilePlanDescriptorBase

	// Unique string that defines the name for the file plan descriptor associated with a particular retention label.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s SecurityFilePlanSubcategory) SecurityFilePlanDescriptorBase() BaseSecurityFilePlanDescriptorBaseImpl {
	return BaseSecurityFilePlanDescriptorBaseImpl{
		DisplayName: s.DisplayName,
		ODataId:     s.ODataId,
		ODataType:   s.ODataType,
	}
}

var _ json.Marshaler = SecurityFilePlanSubcategory{}

func (s SecurityFilePlanSubcategory) MarshalJSON() ([]byte, error) {
	type wrapper SecurityFilePlanSubcategory
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityFilePlanSubcategory: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityFilePlanSubcategory: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.security.filePlanSubcategory"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityFilePlanSubcategory: %+v", err)
	}

	return encoded, nil
}
