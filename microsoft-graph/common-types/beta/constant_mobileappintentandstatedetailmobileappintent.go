package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppIntentAndStateDetailMobileAppIntent string

const (
	MobileAppIntentAndStateDetailMobileAppIntent_Available                         MobileAppIntentAndStateDetailMobileAppIntent = "available"
	MobileAppIntentAndStateDetailMobileAppIntent_AvailableInstallWithoutEnrollment MobileAppIntentAndStateDetailMobileAppIntent = "availableInstallWithoutEnrollment"
	MobileAppIntentAndStateDetailMobileAppIntent_Exclude                           MobileAppIntentAndStateDetailMobileAppIntent = "exclude"
	MobileAppIntentAndStateDetailMobileAppIntent_NotAvailable                      MobileAppIntentAndStateDetailMobileAppIntent = "notAvailable"
	MobileAppIntentAndStateDetailMobileAppIntent_RequiredAndAvailableInstall       MobileAppIntentAndStateDetailMobileAppIntent = "requiredAndAvailableInstall"
	MobileAppIntentAndStateDetailMobileAppIntent_RequiredInstall                   MobileAppIntentAndStateDetailMobileAppIntent = "requiredInstall"
	MobileAppIntentAndStateDetailMobileAppIntent_RequiredUninstall                 MobileAppIntentAndStateDetailMobileAppIntent = "requiredUninstall"
)

func PossibleValuesForMobileAppIntentAndStateDetailMobileAppIntent() []string {
	return []string{
		string(MobileAppIntentAndStateDetailMobileAppIntent_Available),
		string(MobileAppIntentAndStateDetailMobileAppIntent_AvailableInstallWithoutEnrollment),
		string(MobileAppIntentAndStateDetailMobileAppIntent_Exclude),
		string(MobileAppIntentAndStateDetailMobileAppIntent_NotAvailable),
		string(MobileAppIntentAndStateDetailMobileAppIntent_RequiredAndAvailableInstall),
		string(MobileAppIntentAndStateDetailMobileAppIntent_RequiredInstall),
		string(MobileAppIntentAndStateDetailMobileAppIntent_RequiredUninstall),
	}
}

func (s *MobileAppIntentAndStateDetailMobileAppIntent) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMobileAppIntentAndStateDetailMobileAppIntent(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMobileAppIntentAndStateDetailMobileAppIntent(input string) (*MobileAppIntentAndStateDetailMobileAppIntent, error) {
	vals := map[string]MobileAppIntentAndStateDetailMobileAppIntent{
		"available":                         MobileAppIntentAndStateDetailMobileAppIntent_Available,
		"availableinstallwithoutenrollment": MobileAppIntentAndStateDetailMobileAppIntent_AvailableInstallWithoutEnrollment,
		"exclude":                           MobileAppIntentAndStateDetailMobileAppIntent_Exclude,
		"notavailable":                      MobileAppIntentAndStateDetailMobileAppIntent_NotAvailable,
		"requiredandavailableinstall":       MobileAppIntentAndStateDetailMobileAppIntent_RequiredAndAvailableInstall,
		"requiredinstall":                   MobileAppIntentAndStateDetailMobileAppIntent_RequiredInstall,
		"requireduninstall":                 MobileAppIntentAndStateDetailMobileAppIntent_RequiredUninstall,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileAppIntentAndStateDetailMobileAppIntent(input)
	return &out, nil
}
