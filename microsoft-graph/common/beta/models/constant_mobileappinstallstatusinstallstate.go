package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppInstallStatusInstallState string

const (
	MobileAppInstallStatusInstallStatefailed          MobileAppInstallStatusInstallState = "Failed"
	MobileAppInstallStatusInstallStateinstalled       MobileAppInstallStatusInstallState = "Installed"
	MobileAppInstallStatusInstallStatenotApplicable   MobileAppInstallStatusInstallState = "NotApplicable"
	MobileAppInstallStatusInstallStatenotInstalled    MobileAppInstallStatusInstallState = "NotInstalled"
	MobileAppInstallStatusInstallStatependingInstall  MobileAppInstallStatusInstallState = "PendingInstall"
	MobileAppInstallStatusInstallStateuninstallFailed MobileAppInstallStatusInstallState = "UninstallFailed"
	MobileAppInstallStatusInstallStateunknown         MobileAppInstallStatusInstallState = "Unknown"
)

func PossibleValuesForMobileAppInstallStatusInstallState() []string {
	return []string{
		string(MobileAppInstallStatusInstallStatefailed),
		string(MobileAppInstallStatusInstallStateinstalled),
		string(MobileAppInstallStatusInstallStatenotApplicable),
		string(MobileAppInstallStatusInstallStatenotInstalled),
		string(MobileAppInstallStatusInstallStatependingInstall),
		string(MobileAppInstallStatusInstallStateuninstallFailed),
		string(MobileAppInstallStatusInstallStateunknown),
	}
}

func parseMobileAppInstallStatusInstallState(input string) (*MobileAppInstallStatusInstallState, error) {
	vals := map[string]MobileAppInstallStatusInstallState{
		"failed":          MobileAppInstallStatusInstallStatefailed,
		"installed":       MobileAppInstallStatusInstallStateinstalled,
		"notapplicable":   MobileAppInstallStatusInstallStatenotApplicable,
		"notinstalled":    MobileAppInstallStatusInstallStatenotInstalled,
		"pendinginstall":  MobileAppInstallStatusInstallStatependingInstall,
		"uninstallfailed": MobileAppInstallStatusInstallStateuninstallFailed,
		"unknown":         MobileAppInstallStatusInstallStateunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileAppInstallStatusInstallState(input)
	return &out, nil
}
