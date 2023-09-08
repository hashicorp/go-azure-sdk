package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetedManagedAppProtectionMobileThreatDefenseRemediationAction string

const (
	TargetedManagedAppProtectionMobileThreatDefenseRemediationActionblock TargetedManagedAppProtectionMobileThreatDefenseRemediationAction = "Block"
	TargetedManagedAppProtectionMobileThreatDefenseRemediationActionwarn  TargetedManagedAppProtectionMobileThreatDefenseRemediationAction = "Warn"
	TargetedManagedAppProtectionMobileThreatDefenseRemediationActionwipe  TargetedManagedAppProtectionMobileThreatDefenseRemediationAction = "Wipe"
)

func PossibleValuesForTargetedManagedAppProtectionMobileThreatDefenseRemediationAction() []string {
	return []string{
		string(TargetedManagedAppProtectionMobileThreatDefenseRemediationActionblock),
		string(TargetedManagedAppProtectionMobileThreatDefenseRemediationActionwarn),
		string(TargetedManagedAppProtectionMobileThreatDefenseRemediationActionwipe),
	}
}

func parseTargetedManagedAppProtectionMobileThreatDefenseRemediationAction(input string) (*TargetedManagedAppProtectionMobileThreatDefenseRemediationAction, error) {
	vals := map[string]TargetedManagedAppProtectionMobileThreatDefenseRemediationAction{
		"block": TargetedManagedAppProtectionMobileThreatDefenseRemediationActionblock,
		"warn":  TargetedManagedAppProtectionMobileThreatDefenseRemediationActionwarn,
		"wipe":  TargetedManagedAppProtectionMobileThreatDefenseRemediationActionwipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetedManagedAppProtectionMobileThreatDefenseRemediationAction(input)
	return &out, nil
}
