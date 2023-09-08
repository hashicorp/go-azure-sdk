package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppIntentAndStateDetailMobileAppIntent string

const (
	MobileAppIntentAndStateDetailMobileAppIntentavailable                         MobileAppIntentAndStateDetailMobileAppIntent = "Available"
	MobileAppIntentAndStateDetailMobileAppIntentavailableInstallWithoutEnrollment MobileAppIntentAndStateDetailMobileAppIntent = "AvailableInstallWithoutEnrollment"
	MobileAppIntentAndStateDetailMobileAppIntentexclude                           MobileAppIntentAndStateDetailMobileAppIntent = "Exclude"
	MobileAppIntentAndStateDetailMobileAppIntentnotAvailable                      MobileAppIntentAndStateDetailMobileAppIntent = "NotAvailable"
	MobileAppIntentAndStateDetailMobileAppIntentrequiredAndAvailableInstall       MobileAppIntentAndStateDetailMobileAppIntent = "RequiredAndAvailableInstall"
	MobileAppIntentAndStateDetailMobileAppIntentrequiredInstall                   MobileAppIntentAndStateDetailMobileAppIntent = "RequiredInstall"
	MobileAppIntentAndStateDetailMobileAppIntentrequiredUninstall                 MobileAppIntentAndStateDetailMobileAppIntent = "RequiredUninstall"
)

func PossibleValuesForMobileAppIntentAndStateDetailMobileAppIntent() []string {
	return []string{
		string(MobileAppIntentAndStateDetailMobileAppIntentavailable),
		string(MobileAppIntentAndStateDetailMobileAppIntentavailableInstallWithoutEnrollment),
		string(MobileAppIntentAndStateDetailMobileAppIntentexclude),
		string(MobileAppIntentAndStateDetailMobileAppIntentnotAvailable),
		string(MobileAppIntentAndStateDetailMobileAppIntentrequiredAndAvailableInstall),
		string(MobileAppIntentAndStateDetailMobileAppIntentrequiredInstall),
		string(MobileAppIntentAndStateDetailMobileAppIntentrequiredUninstall),
	}
}

func parseMobileAppIntentAndStateDetailMobileAppIntent(input string) (*MobileAppIntentAndStateDetailMobileAppIntent, error) {
	vals := map[string]MobileAppIntentAndStateDetailMobileAppIntent{
		"available":                         MobileAppIntentAndStateDetailMobileAppIntentavailable,
		"availableinstallwithoutenrollment": MobileAppIntentAndStateDetailMobileAppIntentavailableInstallWithoutEnrollment,
		"exclude":                           MobileAppIntentAndStateDetailMobileAppIntentexclude,
		"notavailable":                      MobileAppIntentAndStateDetailMobileAppIntentnotAvailable,
		"requiredandavailableinstall":       MobileAppIntentAndStateDetailMobileAppIntentrequiredAndAvailableInstall,
		"requiredinstall":                   MobileAppIntentAndStateDetailMobileAppIntentrequiredInstall,
		"requireduninstall":                 MobileAppIntentAndStateDetailMobileAppIntentrequiredUninstall,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileAppIntentAndStateDetailMobileAppIntent(input)
	return &out, nil
}
