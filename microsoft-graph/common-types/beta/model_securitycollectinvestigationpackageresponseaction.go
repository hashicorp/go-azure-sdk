package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityResponseAction = SecurityCollectInvestigationPackageResponseAction{}

type SecurityCollectInvestigationPackageResponseAction struct {
	Identifier *SecurityDeviceIdEntityIdentifier `json:"identifier,omitempty"`

	// Fields inherited from SecurityResponseAction

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s SecurityCollectInvestigationPackageResponseAction) SecurityResponseAction() BaseSecurityResponseActionImpl {
	return BaseSecurityResponseActionImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityCollectInvestigationPackageResponseAction{}

func (s SecurityCollectInvestigationPackageResponseAction) MarshalJSON() ([]byte, error) {
	type wrapper SecurityCollectInvestigationPackageResponseAction
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityCollectInvestigationPackageResponseAction: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityCollectInvestigationPackageResponseAction: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.security.collectInvestigationPackageResponseAction"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityCollectInvestigationPackageResponseAction: %+v", err)
	}

	return encoded, nil
}
