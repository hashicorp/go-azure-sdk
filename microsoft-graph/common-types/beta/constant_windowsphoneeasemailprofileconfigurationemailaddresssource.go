package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhoneEASEmailProfileConfigurationEmailAddressSource string

const (
	WindowsPhoneEASEmailProfileConfigurationEmailAddressSource_PrimarySmtpAddress WindowsPhoneEASEmailProfileConfigurationEmailAddressSource = "primarySmtpAddress"
	WindowsPhoneEASEmailProfileConfigurationEmailAddressSource_UserPrincipalName  WindowsPhoneEASEmailProfileConfigurationEmailAddressSource = "userPrincipalName"
)

func PossibleValuesForWindowsPhoneEASEmailProfileConfigurationEmailAddressSource() []string {
	return []string{
		string(WindowsPhoneEASEmailProfileConfigurationEmailAddressSource_PrimarySmtpAddress),
		string(WindowsPhoneEASEmailProfileConfigurationEmailAddressSource_UserPrincipalName),
	}
}

func (s *WindowsPhoneEASEmailProfileConfigurationEmailAddressSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsPhoneEASEmailProfileConfigurationEmailAddressSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsPhoneEASEmailProfileConfigurationEmailAddressSource(input string) (*WindowsPhoneEASEmailProfileConfigurationEmailAddressSource, error) {
	vals := map[string]WindowsPhoneEASEmailProfileConfigurationEmailAddressSource{
		"primarysmtpaddress": WindowsPhoneEASEmailProfileConfigurationEmailAddressSource_PrimarySmtpAddress,
		"userprincipalname":  WindowsPhoneEASEmailProfileConfigurationEmailAddressSource_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhoneEASEmailProfileConfigurationEmailAddressSource(input)
	return &out, nil
}
