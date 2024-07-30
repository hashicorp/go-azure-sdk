package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResellerDelegatedAdminRelationshipStatus string

const (
	ResellerDelegatedAdminRelationshipStatus_Activating           ResellerDelegatedAdminRelationshipStatus = "activating"
	ResellerDelegatedAdminRelationshipStatus_Active               ResellerDelegatedAdminRelationshipStatus = "active"
	ResellerDelegatedAdminRelationshipStatus_ApprovalPending      ResellerDelegatedAdminRelationshipStatus = "approvalPending"
	ResellerDelegatedAdminRelationshipStatus_Approved             ResellerDelegatedAdminRelationshipStatus = "approved"
	ResellerDelegatedAdminRelationshipStatus_Created              ResellerDelegatedAdminRelationshipStatus = "created"
	ResellerDelegatedAdminRelationshipStatus_Expired              ResellerDelegatedAdminRelationshipStatus = "expired"
	ResellerDelegatedAdminRelationshipStatus_Expiring             ResellerDelegatedAdminRelationshipStatus = "expiring"
	ResellerDelegatedAdminRelationshipStatus_Terminated           ResellerDelegatedAdminRelationshipStatus = "terminated"
	ResellerDelegatedAdminRelationshipStatus_Terminating          ResellerDelegatedAdminRelationshipStatus = "terminating"
	ResellerDelegatedAdminRelationshipStatus_TerminationRequested ResellerDelegatedAdminRelationshipStatus = "terminationRequested"
)

func PossibleValuesForResellerDelegatedAdminRelationshipStatus() []string {
	return []string{
		string(ResellerDelegatedAdminRelationshipStatus_Activating),
		string(ResellerDelegatedAdminRelationshipStatus_Active),
		string(ResellerDelegatedAdminRelationshipStatus_ApprovalPending),
		string(ResellerDelegatedAdminRelationshipStatus_Approved),
		string(ResellerDelegatedAdminRelationshipStatus_Created),
		string(ResellerDelegatedAdminRelationshipStatus_Expired),
		string(ResellerDelegatedAdminRelationshipStatus_Expiring),
		string(ResellerDelegatedAdminRelationshipStatus_Terminated),
		string(ResellerDelegatedAdminRelationshipStatus_Terminating),
		string(ResellerDelegatedAdminRelationshipStatus_TerminationRequested),
	}
}

func (s *ResellerDelegatedAdminRelationshipStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseResellerDelegatedAdminRelationshipStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseResellerDelegatedAdminRelationshipStatus(input string) (*ResellerDelegatedAdminRelationshipStatus, error) {
	vals := map[string]ResellerDelegatedAdminRelationshipStatus{
		"activating":           ResellerDelegatedAdminRelationshipStatus_Activating,
		"active":               ResellerDelegatedAdminRelationshipStatus_Active,
		"approvalpending":      ResellerDelegatedAdminRelationshipStatus_ApprovalPending,
		"approved":             ResellerDelegatedAdminRelationshipStatus_Approved,
		"created":              ResellerDelegatedAdminRelationshipStatus_Created,
		"expired":              ResellerDelegatedAdminRelationshipStatus_Expired,
		"expiring":             ResellerDelegatedAdminRelationshipStatus_Expiring,
		"terminated":           ResellerDelegatedAdminRelationshipStatus_Terminated,
		"terminating":          ResellerDelegatedAdminRelationshipStatus_Terminating,
		"terminationrequested": ResellerDelegatedAdminRelationshipStatus_TerminationRequested,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResellerDelegatedAdminRelationshipStatus(input)
	return &out, nil
}
