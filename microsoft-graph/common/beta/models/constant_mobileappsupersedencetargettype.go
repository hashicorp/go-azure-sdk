package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppSupersedenceTargetType string

const (
	MobileAppSupersedenceTargetTypechild  MobileAppSupersedenceTargetType = "Child"
	MobileAppSupersedenceTargetTypeparent MobileAppSupersedenceTargetType = "Parent"
)

func PossibleValuesForMobileAppSupersedenceTargetType() []string {
	return []string{
		string(MobileAppSupersedenceTargetTypechild),
		string(MobileAppSupersedenceTargetTypeparent),
	}
}

func parseMobileAppSupersedenceTargetType(input string) (*MobileAppSupersedenceTargetType, error) {
	vals := map[string]MobileAppSupersedenceTargetType{
		"child":  MobileAppSupersedenceTargetTypechild,
		"parent": MobileAppSupersedenceTargetTypeparent,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileAppSupersedenceTargetType(input)
	return &out, nil
}
