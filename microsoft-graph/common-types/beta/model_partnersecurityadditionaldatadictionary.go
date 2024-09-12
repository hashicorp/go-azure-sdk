package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Dictionary = PartnerSecurityAdditionalDataDictionary{}

type PartnerSecurityAdditionalDataDictionary struct {

	// Fields inherited from Dictionary

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s PartnerSecurityAdditionalDataDictionary) Dictionary() BaseDictionaryImpl {
	return BaseDictionaryImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PartnerSecurityAdditionalDataDictionary{}

func (s PartnerSecurityAdditionalDataDictionary) MarshalJSON() ([]byte, error) {
	type wrapper PartnerSecurityAdditionalDataDictionary
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PartnerSecurityAdditionalDataDictionary: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PartnerSecurityAdditionalDataDictionary: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.partner.security.additionalDataDictionary"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PartnerSecurityAdditionalDataDictionary: %+v", err)
	}

	return encoded, nil
}
