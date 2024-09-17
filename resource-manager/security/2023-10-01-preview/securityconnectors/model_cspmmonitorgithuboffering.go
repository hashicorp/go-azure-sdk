package securityconnectors

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CloudOffering = CspmMonitorGithubOffering{}

type CspmMonitorGithubOffering struct {

	// Fields inherited from CloudOffering

	Description  *string      `json:"description,omitempty"`
	OfferingType OfferingType `json:"offeringType"`
}

func (s CspmMonitorGithubOffering) CloudOffering() BaseCloudOfferingImpl {
	return BaseCloudOfferingImpl{
		Description:  s.Description,
		OfferingType: s.OfferingType,
	}
}

var _ json.Marshaler = CspmMonitorGithubOffering{}

func (s CspmMonitorGithubOffering) MarshalJSON() ([]byte, error) {
	type wrapper CspmMonitorGithubOffering
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CspmMonitorGithubOffering: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CspmMonitorGithubOffering: %+v", err)
	}

	decoded["offeringType"] = "CspmMonitorGithub"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CspmMonitorGithubOffering: %+v", err)
	}

	return encoded, nil
}
