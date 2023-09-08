package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppInstallStatusMobileAppInstallStatusValue string

const (
	MobileAppInstallStatusMobileAppInstallStatusValuefailed          MobileAppInstallStatusMobileAppInstallStatusValue = "Failed"
	MobileAppInstallStatusMobileAppInstallStatusValueinstalled       MobileAppInstallStatusMobileAppInstallStatusValue = "Installed"
	MobileAppInstallStatusMobileAppInstallStatusValuenotApplicable   MobileAppInstallStatusMobileAppInstallStatusValue = "NotApplicable"
	MobileAppInstallStatusMobileAppInstallStatusValuenotInstalled    MobileAppInstallStatusMobileAppInstallStatusValue = "NotInstalled"
	MobileAppInstallStatusMobileAppInstallStatusValuependingInstall  MobileAppInstallStatusMobileAppInstallStatusValue = "PendingInstall"
	MobileAppInstallStatusMobileAppInstallStatusValueuninstallFailed MobileAppInstallStatusMobileAppInstallStatusValue = "UninstallFailed"
	MobileAppInstallStatusMobileAppInstallStatusValueunknown         MobileAppInstallStatusMobileAppInstallStatusValue = "Unknown"
)

func PossibleValuesForMobileAppInstallStatusMobileAppInstallStatusValue() []string {
	return []string{
		string(MobileAppInstallStatusMobileAppInstallStatusValuefailed),
		string(MobileAppInstallStatusMobileAppInstallStatusValueinstalled),
		string(MobileAppInstallStatusMobileAppInstallStatusValuenotApplicable),
		string(MobileAppInstallStatusMobileAppInstallStatusValuenotInstalled),
		string(MobileAppInstallStatusMobileAppInstallStatusValuependingInstall),
		string(MobileAppInstallStatusMobileAppInstallStatusValueuninstallFailed),
		string(MobileAppInstallStatusMobileAppInstallStatusValueunknown),
	}
}

func parseMobileAppInstallStatusMobileAppInstallStatusValue(input string) (*MobileAppInstallStatusMobileAppInstallStatusValue, error) {
	vals := map[string]MobileAppInstallStatusMobileAppInstallStatusValue{
		"failed":          MobileAppInstallStatusMobileAppInstallStatusValuefailed,
		"installed":       MobileAppInstallStatusMobileAppInstallStatusValueinstalled,
		"notapplicable":   MobileAppInstallStatusMobileAppInstallStatusValuenotApplicable,
		"notinstalled":    MobileAppInstallStatusMobileAppInstallStatusValuenotInstalled,
		"pendinginstall":  MobileAppInstallStatusMobileAppInstallStatusValuependingInstall,
		"uninstallfailed": MobileAppInstallStatusMobileAppInstallStatusValueuninstallFailed,
		"unknown":         MobileAppInstallStatusMobileAppInstallStatusValueunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileAppInstallStatusMobileAppInstallStatusValue(input)
	return &out, nil
}
