package securityconnectors

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CloudOffering = InformationProtectionAwsOffering{}

type InformationProtectionAwsOffering struct {
	InformationProtection *InformationProtectionAwsOfferingInformationProtection `json:"informationProtection,omitempty"`

	// Fields inherited from CloudOffering

	Description  *string      `json:"description,omitempty"`
	OfferingType OfferingType `json:"offeringType"`
}

func (s InformationProtectionAwsOffering) CloudOffering() BaseCloudOfferingImpl {
	return BaseCloudOfferingImpl{
		Description:  s.Description,
		OfferingType: s.OfferingType,
	}
}

var _ json.Marshaler = InformationProtectionAwsOffering{}

func (s InformationProtectionAwsOffering) MarshalJSON() ([]byte, error) {
	type wrapper InformationProtectionAwsOffering
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling InformationProtectionAwsOffering: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling InformationProtectionAwsOffering: %+v", err)
	}

	decoded["offeringType"] = "InformationProtectionAws"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling InformationProtectionAwsOffering: %+v", err)
	}

	return encoded, nil
}
