package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionMobileThreatDefenseRemediationAction string

const (
	DefaultManagedAppProtectionMobileThreatDefenseRemediationActionblock DefaultManagedAppProtectionMobileThreatDefenseRemediationAction = "Block"
	DefaultManagedAppProtectionMobileThreatDefenseRemediationActionwarn  DefaultManagedAppProtectionMobileThreatDefenseRemediationAction = "Warn"
	DefaultManagedAppProtectionMobileThreatDefenseRemediationActionwipe  DefaultManagedAppProtectionMobileThreatDefenseRemediationAction = "Wipe"
)

func PossibleValuesForDefaultManagedAppProtectionMobileThreatDefenseRemediationAction() []string {
	return []string{
		string(DefaultManagedAppProtectionMobileThreatDefenseRemediationActionblock),
		string(DefaultManagedAppProtectionMobileThreatDefenseRemediationActionwarn),
		string(DefaultManagedAppProtectionMobileThreatDefenseRemediationActionwipe),
	}
}

func parseDefaultManagedAppProtectionMobileThreatDefenseRemediationAction(input string) (*DefaultManagedAppProtectionMobileThreatDefenseRemediationAction, error) {
	vals := map[string]DefaultManagedAppProtectionMobileThreatDefenseRemediationAction{
		"block": DefaultManagedAppProtectionMobileThreatDefenseRemediationActionblock,
		"warn":  DefaultManagedAppProtectionMobileThreatDefenseRemediationActionwarn,
		"wipe":  DefaultManagedAppProtectionMobileThreatDefenseRemediationActionwipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionMobileThreatDefenseRemediationAction(input)
	return &out, nil
}
