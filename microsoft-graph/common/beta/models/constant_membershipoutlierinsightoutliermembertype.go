package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MembershipOutlierInsightOutlierMemberType string

const (
	MembershipOutlierInsightOutlierMemberTypeuser MembershipOutlierInsightOutlierMemberType = "User"
)

func PossibleValuesForMembershipOutlierInsightOutlierMemberType() []string {
	return []string{
		string(MembershipOutlierInsightOutlierMemberTypeuser),
	}
}

func parseMembershipOutlierInsightOutlierMemberType(input string) (*MembershipOutlierInsightOutlierMemberType, error) {
	vals := map[string]MembershipOutlierInsightOutlierMemberType{
		"user": MembershipOutlierInsightOutlierMemberTypeuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MembershipOutlierInsightOutlierMemberType(input)
	return &out, nil
}
