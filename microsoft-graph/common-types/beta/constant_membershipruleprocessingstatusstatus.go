package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MembershipRuleProcessingStatusStatus string

const (
	MembershipRuleProcessingStatusStatus_Failed                 MembershipRuleProcessingStatusStatus = "Failed"
	MembershipRuleProcessingStatusStatus_NotStarted             MembershipRuleProcessingStatusStatus = "NotStarted"
	MembershipRuleProcessingStatusStatus_Running                MembershipRuleProcessingStatusStatus = "Running"
	MembershipRuleProcessingStatusStatus_Succeeded              MembershipRuleProcessingStatusStatus = "Succeeded"
	MembershipRuleProcessingStatusStatus_UnsupportedFutureValue MembershipRuleProcessingStatusStatus = "UnsupportedFutureValue"
)

func PossibleValuesForMembershipRuleProcessingStatusStatus() []string {
	return []string{
		string(MembershipRuleProcessingStatusStatus_Failed),
		string(MembershipRuleProcessingStatusStatus_NotStarted),
		string(MembershipRuleProcessingStatusStatus_Running),
		string(MembershipRuleProcessingStatusStatus_Succeeded),
		string(MembershipRuleProcessingStatusStatus_UnsupportedFutureValue),
	}
}

func (s *MembershipRuleProcessingStatusStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMembershipRuleProcessingStatusStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMembershipRuleProcessingStatusStatus(input string) (*MembershipRuleProcessingStatusStatus, error) {
	vals := map[string]MembershipRuleProcessingStatusStatus{
		"failed":                 MembershipRuleProcessingStatusStatus_Failed,
		"notstarted":             MembershipRuleProcessingStatusStatus_NotStarted,
		"running":                MembershipRuleProcessingStatusStatus_Running,
		"succeeded":              MembershipRuleProcessingStatusStatus_Succeeded,
		"unsupportedfuturevalue": MembershipRuleProcessingStatusStatus_UnsupportedFutureValue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MembershipRuleProcessingStatusStatus(input)
	return &out, nil
}
