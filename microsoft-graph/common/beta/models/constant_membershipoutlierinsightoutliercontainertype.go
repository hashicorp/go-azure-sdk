package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MembershipOutlierInsightOutlierContainerType string

const (
	MembershipOutlierInsightOutlierContainerTypegroup MembershipOutlierInsightOutlierContainerType = "Group"
)

func PossibleValuesForMembershipOutlierInsightOutlierContainerType() []string {
	return []string{
		string(MembershipOutlierInsightOutlierContainerTypegroup),
	}
}

func parseMembershipOutlierInsightOutlierContainerType(input string) (*MembershipOutlierInsightOutlierContainerType, error) {
	vals := map[string]MembershipOutlierInsightOutlierContainerType{
		"group": MembershipOutlierInsightOutlierContainerTypegroup,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MembershipOutlierInsightOutlierContainerType(input)
	return &out, nil
}
