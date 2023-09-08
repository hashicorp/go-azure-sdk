package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppIntentAndStateDetailInstallState string

const (
	MobileAppIntentAndStateDetailInstallStatefailed          MobileAppIntentAndStateDetailInstallState = "Failed"
	MobileAppIntentAndStateDetailInstallStateinstalled       MobileAppIntentAndStateDetailInstallState = "Installed"
	MobileAppIntentAndStateDetailInstallStatenotApplicable   MobileAppIntentAndStateDetailInstallState = "NotApplicable"
	MobileAppIntentAndStateDetailInstallStatenotInstalled    MobileAppIntentAndStateDetailInstallState = "NotInstalled"
	MobileAppIntentAndStateDetailInstallStatependingInstall  MobileAppIntentAndStateDetailInstallState = "PendingInstall"
	MobileAppIntentAndStateDetailInstallStateuninstallFailed MobileAppIntentAndStateDetailInstallState = "UninstallFailed"
	MobileAppIntentAndStateDetailInstallStateunknown         MobileAppIntentAndStateDetailInstallState = "Unknown"
)

func PossibleValuesForMobileAppIntentAndStateDetailInstallState() []string {
	return []string{
		string(MobileAppIntentAndStateDetailInstallStatefailed),
		string(MobileAppIntentAndStateDetailInstallStateinstalled),
		string(MobileAppIntentAndStateDetailInstallStatenotApplicable),
		string(MobileAppIntentAndStateDetailInstallStatenotInstalled),
		string(MobileAppIntentAndStateDetailInstallStatependingInstall),
		string(MobileAppIntentAndStateDetailInstallStateuninstallFailed),
		string(MobileAppIntentAndStateDetailInstallStateunknown),
	}
}

func parseMobileAppIntentAndStateDetailInstallState(input string) (*MobileAppIntentAndStateDetailInstallState, error) {
	vals := map[string]MobileAppIntentAndStateDetailInstallState{
		"failed":          MobileAppIntentAndStateDetailInstallStatefailed,
		"installed":       MobileAppIntentAndStateDetailInstallStateinstalled,
		"notapplicable":   MobileAppIntentAndStateDetailInstallStatenotApplicable,
		"notinstalled":    MobileAppIntentAndStateDetailInstallStatenotInstalled,
		"pendinginstall":  MobileAppIntentAndStateDetailInstallStatependingInstall,
		"uninstallfailed": MobileAppIntentAndStateDetailInstallStateuninstallFailed,
		"unknown":         MobileAppIntentAndStateDetailInstallStateunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileAppIntentAndStateDetailInstallState(input)
	return &out, nil
}
