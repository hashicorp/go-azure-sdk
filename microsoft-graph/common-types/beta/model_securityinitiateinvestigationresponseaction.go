package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityResponseAction = SecurityInitiateInvestigationResponseAction{}

type SecurityInitiateInvestigationResponseAction struct {
	Identifier *SecurityDeviceIdEntityIdentifier `json:"identifier,omitempty"`

	// Fields inherited from SecurityResponseAction

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s SecurityInitiateInvestigationResponseAction) SecurityResponseAction() BaseSecurityResponseActionImpl {
	return BaseSecurityResponseActionImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityInitiateInvestigationResponseAction{}

func (s SecurityInitiateInvestigationResponseAction) MarshalJSON() ([]byte, error) {
	type wrapper SecurityInitiateInvestigationResponseAction
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityInitiateInvestigationResponseAction: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityInitiateInvestigationResponseAction: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.security.initiateInvestigationResponseAction"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityInitiateInvestigationResponseAction: %+v", err)
	}

	return encoded, nil
}
