package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhoneEASEmailProfileConfigurationUsernameAADSource string

const (
	WindowsPhoneEASEmailProfileConfigurationUsernameAADSource_PrimarySmtpAddress WindowsPhoneEASEmailProfileConfigurationUsernameAADSource = "primarySmtpAddress"
	WindowsPhoneEASEmailProfileConfigurationUsernameAADSource_SamAccountName     WindowsPhoneEASEmailProfileConfigurationUsernameAADSource = "samAccountName"
	WindowsPhoneEASEmailProfileConfigurationUsernameAADSource_UserPrincipalName  WindowsPhoneEASEmailProfileConfigurationUsernameAADSource = "userPrincipalName"
)

func PossibleValuesForWindowsPhoneEASEmailProfileConfigurationUsernameAADSource() []string {
	return []string{
		string(WindowsPhoneEASEmailProfileConfigurationUsernameAADSource_PrimarySmtpAddress),
		string(WindowsPhoneEASEmailProfileConfigurationUsernameAADSource_SamAccountName),
		string(WindowsPhoneEASEmailProfileConfigurationUsernameAADSource_UserPrincipalName),
	}
}

func (s *WindowsPhoneEASEmailProfileConfigurationUsernameAADSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsPhoneEASEmailProfileConfigurationUsernameAADSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsPhoneEASEmailProfileConfigurationUsernameAADSource(input string) (*WindowsPhoneEASEmailProfileConfigurationUsernameAADSource, error) {
	vals := map[string]WindowsPhoneEASEmailProfileConfigurationUsernameAADSource{
		"primarysmtpaddress": WindowsPhoneEASEmailProfileConfigurationUsernameAADSource_PrimarySmtpAddress,
		"samaccountname":     WindowsPhoneEASEmailProfileConfigurationUsernameAADSource_SamAccountName,
		"userprincipalname":  WindowsPhoneEASEmailProfileConfigurationUsernameAADSource_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhoneEASEmailProfileConfigurationUsernameAADSource(input)
	return &out, nil
}
