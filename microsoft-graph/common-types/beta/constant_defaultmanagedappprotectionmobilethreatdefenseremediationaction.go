package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionMobileThreatDefenseRemediationAction string

const (
	DefaultManagedAppProtectionMobileThreatDefenseRemediationAction_Block DefaultManagedAppProtectionMobileThreatDefenseRemediationAction = "block"
	DefaultManagedAppProtectionMobileThreatDefenseRemediationAction_Warn  DefaultManagedAppProtectionMobileThreatDefenseRemediationAction = "warn"
	DefaultManagedAppProtectionMobileThreatDefenseRemediationAction_Wipe  DefaultManagedAppProtectionMobileThreatDefenseRemediationAction = "wipe"
)

func PossibleValuesForDefaultManagedAppProtectionMobileThreatDefenseRemediationAction() []string {
	return []string{
		string(DefaultManagedAppProtectionMobileThreatDefenseRemediationAction_Block),
		string(DefaultManagedAppProtectionMobileThreatDefenseRemediationAction_Warn),
		string(DefaultManagedAppProtectionMobileThreatDefenseRemediationAction_Wipe),
	}
}

func (s *DefaultManagedAppProtectionMobileThreatDefenseRemediationAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDefaultManagedAppProtectionMobileThreatDefenseRemediationAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDefaultManagedAppProtectionMobileThreatDefenseRemediationAction(input string) (*DefaultManagedAppProtectionMobileThreatDefenseRemediationAction, error) {
	vals := map[string]DefaultManagedAppProtectionMobileThreatDefenseRemediationAction{
		"block": DefaultManagedAppProtectionMobileThreatDefenseRemediationAction_Block,
		"warn":  DefaultManagedAppProtectionMobileThreatDefenseRemediationAction_Warn,
		"wipe":  DefaultManagedAppProtectionMobileThreatDefenseRemediationAction_Wipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionMobileThreatDefenseRemediationAction(input)
	return &out, nil
}
