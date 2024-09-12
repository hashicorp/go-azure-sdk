package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ VerifiableCredentialRequirementStatus = VerifiableCredentialVerified{}

type VerifiableCredentialVerified struct {

	// Fields inherited from VerifiableCredentialRequirementStatus

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s VerifiableCredentialVerified) VerifiableCredentialRequirementStatus() BaseVerifiableCredentialRequirementStatusImpl {
	return BaseVerifiableCredentialRequirementStatusImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = VerifiableCredentialVerified{}

func (s VerifiableCredentialVerified) MarshalJSON() ([]byte, error) {
	type wrapper VerifiableCredentialVerified
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling VerifiableCredentialVerified: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling VerifiableCredentialVerified: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.verifiableCredentialVerified"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling VerifiableCredentialVerified: %+v", err)
	}

	return encoded, nil
}
