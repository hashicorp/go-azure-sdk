package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionMobileThreatDefenseRemediationAction string

const (
	AndroidManagedAppProtectionMobileThreatDefenseRemediationActionblock AndroidManagedAppProtectionMobileThreatDefenseRemediationAction = "Block"
	AndroidManagedAppProtectionMobileThreatDefenseRemediationActionwarn  AndroidManagedAppProtectionMobileThreatDefenseRemediationAction = "Warn"
	AndroidManagedAppProtectionMobileThreatDefenseRemediationActionwipe  AndroidManagedAppProtectionMobileThreatDefenseRemediationAction = "Wipe"
)

func PossibleValuesForAndroidManagedAppProtectionMobileThreatDefenseRemediationAction() []string {
	return []string{
		string(AndroidManagedAppProtectionMobileThreatDefenseRemediationActionblock),
		string(AndroidManagedAppProtectionMobileThreatDefenseRemediationActionwarn),
		string(AndroidManagedAppProtectionMobileThreatDefenseRemediationActionwipe),
	}
}

func parseAndroidManagedAppProtectionMobileThreatDefenseRemediationAction(input string) (*AndroidManagedAppProtectionMobileThreatDefenseRemediationAction, error) {
	vals := map[string]AndroidManagedAppProtectionMobileThreatDefenseRemediationAction{
		"block": AndroidManagedAppProtectionMobileThreatDefenseRemediationActionblock,
		"warn":  AndroidManagedAppProtectionMobileThreatDefenseRemediationActionwarn,
		"wipe":  AndroidManagedAppProtectionMobileThreatDefenseRemediationActionwipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionMobileThreatDefenseRemediationAction(input)
	return &out, nil
}
