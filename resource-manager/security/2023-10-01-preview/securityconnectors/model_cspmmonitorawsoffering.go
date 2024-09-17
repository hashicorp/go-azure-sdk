package securityconnectors

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CloudOffering = CspmMonitorAwsOffering{}

type CspmMonitorAwsOffering struct {
	NativeCloudConnection *CspmMonitorAwsOfferingNativeCloudConnection `json:"nativeCloudConnection,omitempty"`

	// Fields inherited from CloudOffering

	Description  *string      `json:"description,omitempty"`
	OfferingType OfferingType `json:"offeringType"`
}

func (s CspmMonitorAwsOffering) CloudOffering() BaseCloudOfferingImpl {
	return BaseCloudOfferingImpl{
		Description:  s.Description,
		OfferingType: s.OfferingType,
	}
}

var _ json.Marshaler = CspmMonitorAwsOffering{}

func (s CspmMonitorAwsOffering) MarshalJSON() ([]byte, error) {
	type wrapper CspmMonitorAwsOffering
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CspmMonitorAwsOffering: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CspmMonitorAwsOffering: %+v", err)
	}

	decoded["offeringType"] = "CspmMonitorAws"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CspmMonitorAwsOffering: %+v", err)
	}

	return encoded, nil
}
