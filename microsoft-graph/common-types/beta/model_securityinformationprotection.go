package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SecurityInformationProtection{}

type SecurityInformationProtection struct {
	// Read the Microsoft Purview Information Protection policy settings for the user or organization.
	LabelPolicySettings *SecurityInformationProtectionPolicySetting `json:"labelPolicySettings,omitempty"`

	// Read the Microsoft Purview Information Protection labels for the user or organization.
	SensitivityLabels *[]SecuritySensitivityLabel `json:"sensitivityLabels,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s SecurityInformationProtection) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityInformationProtection{}

func (s SecurityInformationProtection) MarshalJSON() ([]byte, error) {
	type wrapper SecurityInformationProtection
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityInformationProtection: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityInformationProtection: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.security.informationProtection"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityInformationProtection: %+v", err)
	}

	return encoded, nil
}
