package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhoneEASEmailProfileConfigurationUsernameSource string

const (
	WindowsPhoneEASEmailProfileConfigurationUsernameSource_PrimarySmtpAddress WindowsPhoneEASEmailProfileConfigurationUsernameSource = "primarySmtpAddress"
	WindowsPhoneEASEmailProfileConfigurationUsernameSource_UserPrincipalName  WindowsPhoneEASEmailProfileConfigurationUsernameSource = "userPrincipalName"
)

func PossibleValuesForWindowsPhoneEASEmailProfileConfigurationUsernameSource() []string {
	return []string{
		string(WindowsPhoneEASEmailProfileConfigurationUsernameSource_PrimarySmtpAddress),
		string(WindowsPhoneEASEmailProfileConfigurationUsernameSource_UserPrincipalName),
	}
}

func (s *WindowsPhoneEASEmailProfileConfigurationUsernameSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsPhoneEASEmailProfileConfigurationUsernameSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsPhoneEASEmailProfileConfigurationUsernameSource(input string) (*WindowsPhoneEASEmailProfileConfigurationUsernameSource, error) {
	vals := map[string]WindowsPhoneEASEmailProfileConfigurationUsernameSource{
		"primarysmtpaddress": WindowsPhoneEASEmailProfileConfigurationUsernameSource_PrimarySmtpAddress,
		"userprincipalname":  WindowsPhoneEASEmailProfileConfigurationUsernameSource_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhoneEASEmailProfileConfigurationUsernameSource(input)
	return &out, nil
}
