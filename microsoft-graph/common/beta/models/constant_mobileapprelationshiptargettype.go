package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppRelationshipTargetType string

const (
	MobileAppRelationshipTargetTypechild  MobileAppRelationshipTargetType = "Child"
	MobileAppRelationshipTargetTypeparent MobileAppRelationshipTargetType = "Parent"
)

func PossibleValuesForMobileAppRelationshipTargetType() []string {
	return []string{
		string(MobileAppRelationshipTargetTypechild),
		string(MobileAppRelationshipTargetTypeparent),
	}
}

func parseMobileAppRelationshipTargetType(input string) (*MobileAppRelationshipTargetType, error) {
	vals := map[string]MobileAppRelationshipTargetType{
		"child":  MobileAppRelationshipTargetTypechild,
		"parent": MobileAppRelationshipTargetTypeparent,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileAppRelationshipTargetType(input)
	return &out, nil
}
