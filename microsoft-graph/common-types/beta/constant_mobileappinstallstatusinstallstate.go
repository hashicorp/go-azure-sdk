package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppInstallStatusInstallState string

const (
	MobileAppInstallStatusInstallState_Failed          MobileAppInstallStatusInstallState = "failed"
	MobileAppInstallStatusInstallState_Installed       MobileAppInstallStatusInstallState = "installed"
	MobileAppInstallStatusInstallState_NotApplicable   MobileAppInstallStatusInstallState = "notApplicable"
	MobileAppInstallStatusInstallState_NotInstalled    MobileAppInstallStatusInstallState = "notInstalled"
	MobileAppInstallStatusInstallState_PendingInstall  MobileAppInstallStatusInstallState = "pendingInstall"
	MobileAppInstallStatusInstallState_UninstallFailed MobileAppInstallStatusInstallState = "uninstallFailed"
	MobileAppInstallStatusInstallState_Unknown         MobileAppInstallStatusInstallState = "unknown"
)

func PossibleValuesForMobileAppInstallStatusInstallState() []string {
	return []string{
		string(MobileAppInstallStatusInstallState_Failed),
		string(MobileAppInstallStatusInstallState_Installed),
		string(MobileAppInstallStatusInstallState_NotApplicable),
		string(MobileAppInstallStatusInstallState_NotInstalled),
		string(MobileAppInstallStatusInstallState_PendingInstall),
		string(MobileAppInstallStatusInstallState_UninstallFailed),
		string(MobileAppInstallStatusInstallState_Unknown),
	}
}

func (s *MobileAppInstallStatusInstallState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMobileAppInstallStatusInstallState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMobileAppInstallStatusInstallState(input string) (*MobileAppInstallStatusInstallState, error) {
	vals := map[string]MobileAppInstallStatusInstallState{
		"failed":          MobileAppInstallStatusInstallState_Failed,
		"installed":       MobileAppInstallStatusInstallState_Installed,
		"notapplicable":   MobileAppInstallStatusInstallState_NotApplicable,
		"notinstalled":    MobileAppInstallStatusInstallState_NotInstalled,
		"pendinginstall":  MobileAppInstallStatusInstallState_PendingInstall,
		"uninstallfailed": MobileAppInstallStatusInstallState_UninstallFailed,
		"unknown":         MobileAppInstallStatusInstallState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileAppInstallStatusInstallState(input)
	return &out, nil
}
