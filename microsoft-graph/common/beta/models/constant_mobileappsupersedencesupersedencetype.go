package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppSupersedenceSupersedenceType string

const (
	MobileAppSupersedenceSupersedenceTypereplace MobileAppSupersedenceSupersedenceType = "Replace"
	MobileAppSupersedenceSupersedenceTypeupdate  MobileAppSupersedenceSupersedenceType = "Update"
)

func PossibleValuesForMobileAppSupersedenceSupersedenceType() []string {
	return []string{
		string(MobileAppSupersedenceSupersedenceTypereplace),
		string(MobileAppSupersedenceSupersedenceTypeupdate),
	}
}

func parseMobileAppSupersedenceSupersedenceType(input string) (*MobileAppSupersedenceSupersedenceType, error) {
	vals := map[string]MobileAppSupersedenceSupersedenceType{
		"replace": MobileAppSupersedenceSupersedenceTypereplace,
		"update":  MobileAppSupersedenceSupersedenceTypeupdate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileAppSupersedenceSupersedenceType(input)
	return &out, nil
}
