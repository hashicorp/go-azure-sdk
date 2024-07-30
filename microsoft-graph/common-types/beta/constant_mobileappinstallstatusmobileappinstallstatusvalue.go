package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppInstallStatusMobileAppInstallStatusValue string

const (
	MobileAppInstallStatusMobileAppInstallStatusValue_Failed          MobileAppInstallStatusMobileAppInstallStatusValue = "failed"
	MobileAppInstallStatusMobileAppInstallStatusValue_Installed       MobileAppInstallStatusMobileAppInstallStatusValue = "installed"
	MobileAppInstallStatusMobileAppInstallStatusValue_NotApplicable   MobileAppInstallStatusMobileAppInstallStatusValue = "notApplicable"
	MobileAppInstallStatusMobileAppInstallStatusValue_NotInstalled    MobileAppInstallStatusMobileAppInstallStatusValue = "notInstalled"
	MobileAppInstallStatusMobileAppInstallStatusValue_PendingInstall  MobileAppInstallStatusMobileAppInstallStatusValue = "pendingInstall"
	MobileAppInstallStatusMobileAppInstallStatusValue_UninstallFailed MobileAppInstallStatusMobileAppInstallStatusValue = "uninstallFailed"
	MobileAppInstallStatusMobileAppInstallStatusValue_Unknown         MobileAppInstallStatusMobileAppInstallStatusValue = "unknown"
)

func PossibleValuesForMobileAppInstallStatusMobileAppInstallStatusValue() []string {
	return []string{
		string(MobileAppInstallStatusMobileAppInstallStatusValue_Failed),
		string(MobileAppInstallStatusMobileAppInstallStatusValue_Installed),
		string(MobileAppInstallStatusMobileAppInstallStatusValue_NotApplicable),
		string(MobileAppInstallStatusMobileAppInstallStatusValue_NotInstalled),
		string(MobileAppInstallStatusMobileAppInstallStatusValue_PendingInstall),
		string(MobileAppInstallStatusMobileAppInstallStatusValue_UninstallFailed),
		string(MobileAppInstallStatusMobileAppInstallStatusValue_Unknown),
	}
}

func (s *MobileAppInstallStatusMobileAppInstallStatusValue) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMobileAppInstallStatusMobileAppInstallStatusValue(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMobileAppInstallStatusMobileAppInstallStatusValue(input string) (*MobileAppInstallStatusMobileAppInstallStatusValue, error) {
	vals := map[string]MobileAppInstallStatusMobileAppInstallStatusValue{
		"failed":          MobileAppInstallStatusMobileAppInstallStatusValue_Failed,
		"installed":       MobileAppInstallStatusMobileAppInstallStatusValue_Installed,
		"notapplicable":   MobileAppInstallStatusMobileAppInstallStatusValue_NotApplicable,
		"notinstalled":    MobileAppInstallStatusMobileAppInstallStatusValue_NotInstalled,
		"pendinginstall":  MobileAppInstallStatusMobileAppInstallStatusValue_PendingInstall,
		"uninstallfailed": MobileAppInstallStatusMobileAppInstallStatusValue_UninstallFailed,
		"unknown":         MobileAppInstallStatusMobileAppInstallStatusValue_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileAppInstallStatusMobileAppInstallStatusValue(input)
	return &out, nil
}
