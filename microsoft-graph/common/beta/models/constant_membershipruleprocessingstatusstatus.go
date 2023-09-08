package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MembershipRuleProcessingStatusStatus string

const (
	MembershipRuleProcessingStatusStatusFailed                 MembershipRuleProcessingStatusStatus = "Failed"
	MembershipRuleProcessingStatusStatusNotStarted             MembershipRuleProcessingStatusStatus = "NotStarted"
	MembershipRuleProcessingStatusStatusRunning                MembershipRuleProcessingStatusStatus = "Running"
	MembershipRuleProcessingStatusStatusSucceeded              MembershipRuleProcessingStatusStatus = "Succeeded"
	MembershipRuleProcessingStatusStatusUnsupportedFutureValue MembershipRuleProcessingStatusStatus = "UnsupportedFutureValue"
)

func PossibleValuesForMembershipRuleProcessingStatusStatus() []string {
	return []string{
		string(MembershipRuleProcessingStatusStatusFailed),
		string(MembershipRuleProcessingStatusStatusNotStarted),
		string(MembershipRuleProcessingStatusStatusRunning),
		string(MembershipRuleProcessingStatusStatusSucceeded),
		string(MembershipRuleProcessingStatusStatusUnsupportedFutureValue),
	}
}

func parseMembershipRuleProcessingStatusStatus(input string) (*MembershipRuleProcessingStatusStatus, error) {
	vals := map[string]MembershipRuleProcessingStatusStatus{
		"failed":                 MembershipRuleProcessingStatusStatusFailed,
		"notstarted":             MembershipRuleProcessingStatusStatusNotStarted,
		"running":                MembershipRuleProcessingStatusStatusRunning,
		"succeeded":              MembershipRuleProcessingStatusStatusSucceeded,
		"unsupportedfuturevalue": MembershipRuleProcessingStatusStatusUnsupportedFutureValue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MembershipRuleProcessingStatusStatus(input)
	return &out, nil
}
