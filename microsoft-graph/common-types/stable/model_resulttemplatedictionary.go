package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Dictionary = ResultTemplateDictionary{}

type ResultTemplateDictionary struct {

	// Fields inherited from Dictionary

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s ResultTemplateDictionary) Dictionary() BaseDictionaryImpl {
	return BaseDictionaryImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ResultTemplateDictionary{}

func (s ResultTemplateDictionary) MarshalJSON() ([]byte, error) {
	type wrapper ResultTemplateDictionary
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ResultTemplateDictionary: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ResultTemplateDictionary: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.resultTemplateDictionary"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ResultTemplateDictionary: %+v", err)
	}

	return encoded, nil
}
