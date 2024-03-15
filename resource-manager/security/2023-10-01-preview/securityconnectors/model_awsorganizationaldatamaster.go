package securityconnectors

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AwsOrganizationalData = AwsOrganizationalDataMaster{}

type AwsOrganizationalDataMaster struct {
	ExcludedAccountIds *[]string `json:"excludedAccountIds,omitempty"`
	StacksetName       *string   `json:"stacksetName,omitempty"`

	// Fields inherited from AwsOrganizationalData
}

var _ json.Marshaler = AwsOrganizationalDataMaster{}

func (s AwsOrganizationalDataMaster) MarshalJSON() ([]byte, error) {
	type wrapper AwsOrganizationalDataMaster
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AwsOrganizationalDataMaster: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AwsOrganizationalDataMaster: %+v", err)
	}
	decoded["organizationMembershipType"] = "Organization"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AwsOrganizationalDataMaster: %+v", err)
	}

	return encoded, nil
}
