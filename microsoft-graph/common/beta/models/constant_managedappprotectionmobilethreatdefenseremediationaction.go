package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppProtectionMobileThreatDefenseRemediationAction string

const (
	ManagedAppProtectionMobileThreatDefenseRemediationActionblock ManagedAppProtectionMobileThreatDefenseRemediationAction = "Block"
	ManagedAppProtectionMobileThreatDefenseRemediationActionwarn  ManagedAppProtectionMobileThreatDefenseRemediationAction = "Warn"
	ManagedAppProtectionMobileThreatDefenseRemediationActionwipe  ManagedAppProtectionMobileThreatDefenseRemediationAction = "Wipe"
)

func PossibleValuesForManagedAppProtectionMobileThreatDefenseRemediationAction() []string {
	return []string{
		string(ManagedAppProtectionMobileThreatDefenseRemediationActionblock),
		string(ManagedAppProtectionMobileThreatDefenseRemediationActionwarn),
		string(ManagedAppProtectionMobileThreatDefenseRemediationActionwipe),
	}
}

func parseManagedAppProtectionMobileThreatDefenseRemediationAction(input string) (*ManagedAppProtectionMobileThreatDefenseRemediationAction, error) {
	vals := map[string]ManagedAppProtectionMobileThreatDefenseRemediationAction{
		"block": ManagedAppProtectionMobileThreatDefenseRemediationActionblock,
		"warn":  ManagedAppProtectionMobileThreatDefenseRemediationActionwarn,
		"wipe":  ManagedAppProtectionMobileThreatDefenseRemediationActionwipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppProtectionMobileThreatDefenseRemediationAction(input)
	return &out, nil
}
