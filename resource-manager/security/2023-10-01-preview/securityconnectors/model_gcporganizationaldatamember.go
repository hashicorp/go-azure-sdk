package securityconnectors

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ GcpOrganizationalData = GcpOrganizationalDataMember{}

type GcpOrganizationalDataMember struct {
	ManagementProjectNumber *string `json:"managementProjectNumber,omitempty"`
	ParentHierarchyId       *string `json:"parentHierarchyId,omitempty"`

	// Fields inherited from GcpOrganizationalData
}

var _ json.Marshaler = GcpOrganizationalDataMember{}

func (s GcpOrganizationalDataMember) MarshalJSON() ([]byte, error) {
	type wrapper GcpOrganizationalDataMember
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GcpOrganizationalDataMember: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GcpOrganizationalDataMember: %+v", err)
	}
	decoded["organizationMembershipType"] = "Member"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GcpOrganizationalDataMember: %+v", err)
	}

	return encoded, nil
}
