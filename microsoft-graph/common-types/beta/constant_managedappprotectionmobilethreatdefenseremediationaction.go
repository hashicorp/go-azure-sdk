package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppProtectionMobileThreatDefenseRemediationAction string

const (
	ManagedAppProtectionMobileThreatDefenseRemediationAction_Block ManagedAppProtectionMobileThreatDefenseRemediationAction = "block"
	ManagedAppProtectionMobileThreatDefenseRemediationAction_Warn  ManagedAppProtectionMobileThreatDefenseRemediationAction = "warn"
	ManagedAppProtectionMobileThreatDefenseRemediationAction_Wipe  ManagedAppProtectionMobileThreatDefenseRemediationAction = "wipe"
)

func PossibleValuesForManagedAppProtectionMobileThreatDefenseRemediationAction() []string {
	return []string{
		string(ManagedAppProtectionMobileThreatDefenseRemediationAction_Block),
		string(ManagedAppProtectionMobileThreatDefenseRemediationAction_Warn),
		string(ManagedAppProtectionMobileThreatDefenseRemediationAction_Wipe),
	}
}

func (s *ManagedAppProtectionMobileThreatDefenseRemediationAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedAppProtectionMobileThreatDefenseRemediationAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedAppProtectionMobileThreatDefenseRemediationAction(input string) (*ManagedAppProtectionMobileThreatDefenseRemediationAction, error) {
	vals := map[string]ManagedAppProtectionMobileThreatDefenseRemediationAction{
		"block": ManagedAppProtectionMobileThreatDefenseRemediationAction_Block,
		"warn":  ManagedAppProtectionMobileThreatDefenseRemediationAction_Warn,
		"wipe":  ManagedAppProtectionMobileThreatDefenseRemediationAction_Wipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppProtectionMobileThreatDefenseRemediationAction(input)
	return &out, nil
}
