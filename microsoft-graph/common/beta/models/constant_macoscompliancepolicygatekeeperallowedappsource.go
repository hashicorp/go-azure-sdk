package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSCompliancePolicyGatekeeperAllowedAppSource string

const (
	MacOSCompliancePolicyGatekeeperAllowedAppSourceanywhere                           MacOSCompliancePolicyGatekeeperAllowedAppSource = "Anywhere"
	MacOSCompliancePolicyGatekeeperAllowedAppSourcemacAppStore                        MacOSCompliancePolicyGatekeeperAllowedAppSource = "MacAppStore"
	MacOSCompliancePolicyGatekeeperAllowedAppSourcemacAppStoreAndIdentifiedDevelopers MacOSCompliancePolicyGatekeeperAllowedAppSource = "MacAppStoreAndIdentifiedDevelopers"
	MacOSCompliancePolicyGatekeeperAllowedAppSourcenotConfigured                      MacOSCompliancePolicyGatekeeperAllowedAppSource = "NotConfigured"
)

func PossibleValuesForMacOSCompliancePolicyGatekeeperAllowedAppSource() []string {
	return []string{
		string(MacOSCompliancePolicyGatekeeperAllowedAppSourceanywhere),
		string(MacOSCompliancePolicyGatekeeperAllowedAppSourcemacAppStore),
		string(MacOSCompliancePolicyGatekeeperAllowedAppSourcemacAppStoreAndIdentifiedDevelopers),
		string(MacOSCompliancePolicyGatekeeperAllowedAppSourcenotConfigured),
	}
}

func parseMacOSCompliancePolicyGatekeeperAllowedAppSource(input string) (*MacOSCompliancePolicyGatekeeperAllowedAppSource, error) {
	vals := map[string]MacOSCompliancePolicyGatekeeperAllowedAppSource{
		"anywhere":                           MacOSCompliancePolicyGatekeeperAllowedAppSourceanywhere,
		"macappstore":                        MacOSCompliancePolicyGatekeeperAllowedAppSourcemacAppStore,
		"macappstoreandidentifieddevelopers": MacOSCompliancePolicyGatekeeperAllowedAppSourcemacAppStoreAndIdentifiedDevelopers,
		"notconfigured":                      MacOSCompliancePolicyGatekeeperAllowedAppSourcenotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSCompliancePolicyGatekeeperAllowedAppSource(input)
	return &out, nil
}
