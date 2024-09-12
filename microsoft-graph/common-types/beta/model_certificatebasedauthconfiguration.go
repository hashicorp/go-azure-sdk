package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = CertificateBasedAuthConfiguration{}

type CertificateBasedAuthConfiguration struct {
	// Collection of certificate authorities which creates a trusted certificate chain.
	CertificateAuthorities *[]CertificateAuthority `json:"certificateAuthorities,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s CertificateBasedAuthConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CertificateBasedAuthConfiguration{}

func (s CertificateBasedAuthConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper CertificateBasedAuthConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CertificateBasedAuthConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CertificateBasedAuthConfiguration: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.certificateBasedAuthConfiguration"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CertificateBasedAuthConfiguration: %+v", err)
	}

	return encoded, nil
}
