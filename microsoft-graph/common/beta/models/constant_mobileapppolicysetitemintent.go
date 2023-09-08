package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppPolicySetItemIntent string

const (
	MobileAppPolicySetItemIntentavailable                  MobileAppPolicySetItemIntent = "Available"
	MobileAppPolicySetItemIntentavailableWithoutEnrollment MobileAppPolicySetItemIntent = "AvailableWithoutEnrollment"
	MobileAppPolicySetItemIntentrequired                   MobileAppPolicySetItemIntent = "Required"
	MobileAppPolicySetItemIntentuninstall                  MobileAppPolicySetItemIntent = "Uninstall"
)

func PossibleValuesForMobileAppPolicySetItemIntent() []string {
	return []string{
		string(MobileAppPolicySetItemIntentavailable),
		string(MobileAppPolicySetItemIntentavailableWithoutEnrollment),
		string(MobileAppPolicySetItemIntentrequired),
		string(MobileAppPolicySetItemIntentuninstall),
	}
}

func parseMobileAppPolicySetItemIntent(input string) (*MobileAppPolicySetItemIntent, error) {
	vals := map[string]MobileAppPolicySetItemIntent{
		"available":                  MobileAppPolicySetItemIntentavailable,
		"availablewithoutenrollment": MobileAppPolicySetItemIntentavailableWithoutEnrollment,
		"required":                   MobileAppPolicySetItemIntentrequired,
		"uninstall":                  MobileAppPolicySetItemIntentuninstall,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileAppPolicySetItemIntent(input)
	return &out, nil
}
