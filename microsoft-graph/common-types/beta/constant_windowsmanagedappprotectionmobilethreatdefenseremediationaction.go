package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsManagedAppProtectionMobileThreatDefenseRemediationAction string

const (
	WindowsManagedAppProtectionMobileThreatDefenseRemediationAction_Block WindowsManagedAppProtectionMobileThreatDefenseRemediationAction = "block"
	WindowsManagedAppProtectionMobileThreatDefenseRemediationAction_Warn  WindowsManagedAppProtectionMobileThreatDefenseRemediationAction = "warn"
	WindowsManagedAppProtectionMobileThreatDefenseRemediationAction_Wipe  WindowsManagedAppProtectionMobileThreatDefenseRemediationAction = "wipe"
)

func PossibleValuesForWindowsManagedAppProtectionMobileThreatDefenseRemediationAction() []string {
	return []string{
		string(WindowsManagedAppProtectionMobileThreatDefenseRemediationAction_Block),
		string(WindowsManagedAppProtectionMobileThreatDefenseRemediationAction_Warn),
		string(WindowsManagedAppProtectionMobileThreatDefenseRemediationAction_Wipe),
	}
}

func (s *WindowsManagedAppProtectionMobileThreatDefenseRemediationAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsManagedAppProtectionMobileThreatDefenseRemediationAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsManagedAppProtectionMobileThreatDefenseRemediationAction(input string) (*WindowsManagedAppProtectionMobileThreatDefenseRemediationAction, error) {
	vals := map[string]WindowsManagedAppProtectionMobileThreatDefenseRemediationAction{
		"block": WindowsManagedAppProtectionMobileThreatDefenseRemediationAction_Block,
		"warn":  WindowsManagedAppProtectionMobileThreatDefenseRemediationAction_Warn,
		"wipe":  WindowsManagedAppProtectionMobileThreatDefenseRemediationAction_Wipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsManagedAppProtectionMobileThreatDefenseRemediationAction(input)
	return &out, nil
}
