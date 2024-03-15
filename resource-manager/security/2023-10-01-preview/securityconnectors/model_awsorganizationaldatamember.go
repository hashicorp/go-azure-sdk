package securityconnectors

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AwsOrganizationalData = AwsOrganizationalDataMember{}

type AwsOrganizationalDataMember struct {
	ParentHierarchyId *string `json:"parentHierarchyId,omitempty"`

	// Fields inherited from AwsOrganizationalData
}

var _ json.Marshaler = AwsOrganizationalDataMember{}

func (s AwsOrganizationalDataMember) MarshalJSON() ([]byte, error) {
	type wrapper AwsOrganizationalDataMember
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AwsOrganizationalDataMember: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AwsOrganizationalDataMember: %+v", err)
	}
	decoded["organizationMembershipType"] = "Member"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AwsOrganizationalDataMember: %+v", err)
	}

	return encoded, nil
}
