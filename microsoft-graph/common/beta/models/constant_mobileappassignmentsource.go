package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppAssignmentSource string

const (
	MobileAppAssignmentSourcedirect     MobileAppAssignmentSource = "Direct"
	MobileAppAssignmentSourcepolicySets MobileAppAssignmentSource = "PolicySets"
)

func PossibleValuesForMobileAppAssignmentSource() []string {
	return []string{
		string(MobileAppAssignmentSourcedirect),
		string(MobileAppAssignmentSourcepolicySets),
	}
}

func parseMobileAppAssignmentSource(input string) (*MobileAppAssignmentSource, error) {
	vals := map[string]MobileAppAssignmentSource{
		"direct":     MobileAppAssignmentSourcedirect,
		"policysets": MobileAppAssignmentSourcepolicySets,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileAppAssignmentSource(input)
	return &out, nil
}
