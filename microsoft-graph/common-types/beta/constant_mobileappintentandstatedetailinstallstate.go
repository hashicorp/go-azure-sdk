package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppIntentAndStateDetailInstallState string

const (
	MobileAppIntentAndStateDetailInstallState_Failed          MobileAppIntentAndStateDetailInstallState = "failed"
	MobileAppIntentAndStateDetailInstallState_Installed       MobileAppIntentAndStateDetailInstallState = "installed"
	MobileAppIntentAndStateDetailInstallState_NotApplicable   MobileAppIntentAndStateDetailInstallState = "notApplicable"
	MobileAppIntentAndStateDetailInstallState_NotInstalled    MobileAppIntentAndStateDetailInstallState = "notInstalled"
	MobileAppIntentAndStateDetailInstallState_PendingInstall  MobileAppIntentAndStateDetailInstallState = "pendingInstall"
	MobileAppIntentAndStateDetailInstallState_UninstallFailed MobileAppIntentAndStateDetailInstallState = "uninstallFailed"
	MobileAppIntentAndStateDetailInstallState_Unknown         MobileAppIntentAndStateDetailInstallState = "unknown"
)

func PossibleValuesForMobileAppIntentAndStateDetailInstallState() []string {
	return []string{
		string(MobileAppIntentAndStateDetailInstallState_Failed),
		string(MobileAppIntentAndStateDetailInstallState_Installed),
		string(MobileAppIntentAndStateDetailInstallState_NotApplicable),
		string(MobileAppIntentAndStateDetailInstallState_NotInstalled),
		string(MobileAppIntentAndStateDetailInstallState_PendingInstall),
		string(MobileAppIntentAndStateDetailInstallState_UninstallFailed),
		string(MobileAppIntentAndStateDetailInstallState_Unknown),
	}
}

func (s *MobileAppIntentAndStateDetailInstallState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMobileAppIntentAndStateDetailInstallState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMobileAppIntentAndStateDetailInstallState(input string) (*MobileAppIntentAndStateDetailInstallState, error) {
	vals := map[string]MobileAppIntentAndStateDetailInstallState{
		"failed":          MobileAppIntentAndStateDetailInstallState_Failed,
		"installed":       MobileAppIntentAndStateDetailInstallState_Installed,
		"notapplicable":   MobileAppIntentAndStateDetailInstallState_NotApplicable,
		"notinstalled":    MobileAppIntentAndStateDetailInstallState_NotInstalled,
		"pendinginstall":  MobileAppIntentAndStateDetailInstallState_PendingInstall,
		"uninstallfailed": MobileAppIntentAndStateDetailInstallState_UninstallFailed,
		"unknown":         MobileAppIntentAndStateDetailInstallState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileAppIntentAndStateDetailInstallState(input)
	return &out, nil
}
