package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetedManagedAppProtectionMobileThreatDefenseRemediationAction string

const (
	TargetedManagedAppProtectionMobileThreatDefenseRemediationAction_Block TargetedManagedAppProtectionMobileThreatDefenseRemediationAction = "block"
	TargetedManagedAppProtectionMobileThreatDefenseRemediationAction_Warn  TargetedManagedAppProtectionMobileThreatDefenseRemediationAction = "warn"
	TargetedManagedAppProtectionMobileThreatDefenseRemediationAction_Wipe  TargetedManagedAppProtectionMobileThreatDefenseRemediationAction = "wipe"
)

func PossibleValuesForTargetedManagedAppProtectionMobileThreatDefenseRemediationAction() []string {
	return []string{
		string(TargetedManagedAppProtectionMobileThreatDefenseRemediationAction_Block),
		string(TargetedManagedAppProtectionMobileThreatDefenseRemediationAction_Warn),
		string(TargetedManagedAppProtectionMobileThreatDefenseRemediationAction_Wipe),
	}
}

func (s *TargetedManagedAppProtectionMobileThreatDefenseRemediationAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTargetedManagedAppProtectionMobileThreatDefenseRemediationAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTargetedManagedAppProtectionMobileThreatDefenseRemediationAction(input string) (*TargetedManagedAppProtectionMobileThreatDefenseRemediationAction, error) {
	vals := map[string]TargetedManagedAppProtectionMobileThreatDefenseRemediationAction{
		"block": TargetedManagedAppProtectionMobileThreatDefenseRemediationAction_Block,
		"warn":  TargetedManagedAppProtectionMobileThreatDefenseRemediationAction_Warn,
		"wipe":  TargetedManagedAppProtectionMobileThreatDefenseRemediationAction_Wipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetedManagedAppProtectionMobileThreatDefenseRemediationAction(input)
	return &out, nil
}
