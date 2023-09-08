package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsManagedAppProtectionMobileThreatDefenseRemediationAction string

const (
	WindowsManagedAppProtectionMobileThreatDefenseRemediationActionblock WindowsManagedAppProtectionMobileThreatDefenseRemediationAction = "Block"
	WindowsManagedAppProtectionMobileThreatDefenseRemediationActionwarn  WindowsManagedAppProtectionMobileThreatDefenseRemediationAction = "Warn"
	WindowsManagedAppProtectionMobileThreatDefenseRemediationActionwipe  WindowsManagedAppProtectionMobileThreatDefenseRemediationAction = "Wipe"
)

func PossibleValuesForWindowsManagedAppProtectionMobileThreatDefenseRemediationAction() []string {
	return []string{
		string(WindowsManagedAppProtectionMobileThreatDefenseRemediationActionblock),
		string(WindowsManagedAppProtectionMobileThreatDefenseRemediationActionwarn),
		string(WindowsManagedAppProtectionMobileThreatDefenseRemediationActionwipe),
	}
}

func parseWindowsManagedAppProtectionMobileThreatDefenseRemediationAction(input string) (*WindowsManagedAppProtectionMobileThreatDefenseRemediationAction, error) {
	vals := map[string]WindowsManagedAppProtectionMobileThreatDefenseRemediationAction{
		"block": WindowsManagedAppProtectionMobileThreatDefenseRemediationActionblock,
		"warn":  WindowsManagedAppProtectionMobileThreatDefenseRemediationActionwarn,
		"wipe":  WindowsManagedAppProtectionMobileThreatDefenseRemediationActionwipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsManagedAppProtectionMobileThreatDefenseRemediationAction(input)
	return &out, nil
}
