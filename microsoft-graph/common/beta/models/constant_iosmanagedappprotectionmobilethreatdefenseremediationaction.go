package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosManagedAppProtectionMobileThreatDefenseRemediationAction string

const (
	IosManagedAppProtectionMobileThreatDefenseRemediationActionblock IosManagedAppProtectionMobileThreatDefenseRemediationAction = "Block"
	IosManagedAppProtectionMobileThreatDefenseRemediationActionwarn  IosManagedAppProtectionMobileThreatDefenseRemediationAction = "Warn"
	IosManagedAppProtectionMobileThreatDefenseRemediationActionwipe  IosManagedAppProtectionMobileThreatDefenseRemediationAction = "Wipe"
)

func PossibleValuesForIosManagedAppProtectionMobileThreatDefenseRemediationAction() []string {
	return []string{
		string(IosManagedAppProtectionMobileThreatDefenseRemediationActionblock),
		string(IosManagedAppProtectionMobileThreatDefenseRemediationActionwarn),
		string(IosManagedAppProtectionMobileThreatDefenseRemediationActionwipe),
	}
}

func parseIosManagedAppProtectionMobileThreatDefenseRemediationAction(input string) (*IosManagedAppProtectionMobileThreatDefenseRemediationAction, error) {
	vals := map[string]IosManagedAppProtectionMobileThreatDefenseRemediationAction{
		"block": IosManagedAppProtectionMobileThreatDefenseRemediationActionblock,
		"warn":  IosManagedAppProtectionMobileThreatDefenseRemediationActionwarn,
		"wipe":  IosManagedAppProtectionMobileThreatDefenseRemediationActionwipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosManagedAppProtectionMobileThreatDefenseRemediationAction(input)
	return &out, nil
}
