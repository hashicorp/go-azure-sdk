package securityconnectors

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ GcpOrganizationalData = GcpOrganizationalDataOrganization{}

type GcpOrganizationalDataOrganization struct {
	ExcludedProjectNumbers     *[]string `json:"excludedProjectNumbers,omitempty"`
	OrganizationName           *string   `json:"organizationName,omitempty"`
	ServiceAccountEmailAddress *string   `json:"serviceAccountEmailAddress,omitempty"`
	WorkloadIdentityProviderId *string   `json:"workloadIdentityProviderId,omitempty"`

	// Fields inherited from GcpOrganizationalData
}

var _ json.Marshaler = GcpOrganizationalDataOrganization{}

func (s GcpOrganizationalDataOrganization) MarshalJSON() ([]byte, error) {
	type wrapper GcpOrganizationalDataOrganization
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GcpOrganizationalDataOrganization: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GcpOrganizationalDataOrganization: %+v", err)
	}
	decoded["organizationMembershipType"] = "Organization"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GcpOrganizationalDataOrganization: %+v", err)
	}

	return encoded, nil
}
