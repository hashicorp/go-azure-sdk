package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosManagedAppProtectionMobileThreatDefenseRemediationAction string

const (
	IosManagedAppProtectionMobileThreatDefenseRemediationAction_Block IosManagedAppProtectionMobileThreatDefenseRemediationAction = "block"
	IosManagedAppProtectionMobileThreatDefenseRemediationAction_Warn  IosManagedAppProtectionMobileThreatDefenseRemediationAction = "warn"
	IosManagedAppProtectionMobileThreatDefenseRemediationAction_Wipe  IosManagedAppProtectionMobileThreatDefenseRemediationAction = "wipe"
)

func PossibleValuesForIosManagedAppProtectionMobileThreatDefenseRemediationAction() []string {
	return []string{
		string(IosManagedAppProtectionMobileThreatDefenseRemediationAction_Block),
		string(IosManagedAppProtectionMobileThreatDefenseRemediationAction_Warn),
		string(IosManagedAppProtectionMobileThreatDefenseRemediationAction_Wipe),
	}
}

func (s *IosManagedAppProtectionMobileThreatDefenseRemediationAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosManagedAppProtectionMobileThreatDefenseRemediationAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosManagedAppProtectionMobileThreatDefenseRemediationAction(input string) (*IosManagedAppProtectionMobileThreatDefenseRemediationAction, error) {
	vals := map[string]IosManagedAppProtectionMobileThreatDefenseRemediationAction{
		"block": IosManagedAppProtectionMobileThreatDefenseRemediationAction_Block,
		"warn":  IosManagedAppProtectionMobileThreatDefenseRemediationAction_Warn,
		"wipe":  IosManagedAppProtectionMobileThreatDefenseRemediationAction_Wipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosManagedAppProtectionMobileThreatDefenseRemediationAction(input)
	return &out, nil
}
