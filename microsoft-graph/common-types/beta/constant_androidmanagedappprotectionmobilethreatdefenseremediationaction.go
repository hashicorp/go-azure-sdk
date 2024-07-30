package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionMobileThreatDefenseRemediationAction string

const (
	AndroidManagedAppProtectionMobileThreatDefenseRemediationAction_Block AndroidManagedAppProtectionMobileThreatDefenseRemediationAction = "block"
	AndroidManagedAppProtectionMobileThreatDefenseRemediationAction_Warn  AndroidManagedAppProtectionMobileThreatDefenseRemediationAction = "warn"
	AndroidManagedAppProtectionMobileThreatDefenseRemediationAction_Wipe  AndroidManagedAppProtectionMobileThreatDefenseRemediationAction = "wipe"
)

func PossibleValuesForAndroidManagedAppProtectionMobileThreatDefenseRemediationAction() []string {
	return []string{
		string(AndroidManagedAppProtectionMobileThreatDefenseRemediationAction_Block),
		string(AndroidManagedAppProtectionMobileThreatDefenseRemediationAction_Warn),
		string(AndroidManagedAppProtectionMobileThreatDefenseRemediationAction_Wipe),
	}
}

func (s *AndroidManagedAppProtectionMobileThreatDefenseRemediationAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidManagedAppProtectionMobileThreatDefenseRemediationAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidManagedAppProtectionMobileThreatDefenseRemediationAction(input string) (*AndroidManagedAppProtectionMobileThreatDefenseRemediationAction, error) {
	vals := map[string]AndroidManagedAppProtectionMobileThreatDefenseRemediationAction{
		"block": AndroidManagedAppProtectionMobileThreatDefenseRemediationAction_Block,
		"warn":  AndroidManagedAppProtectionMobileThreatDefenseRemediationAction_Warn,
		"wipe":  AndroidManagedAppProtectionMobileThreatDefenseRemediationAction_Wipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionMobileThreatDefenseRemediationAction(input)
	return &out, nil
}
