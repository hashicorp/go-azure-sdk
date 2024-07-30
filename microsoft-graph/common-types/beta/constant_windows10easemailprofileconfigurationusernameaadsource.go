package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EasEmailProfileConfigurationUsernameAADSource string

const (
	Windows10EasEmailProfileConfigurationUsernameAADSource_PrimarySmtpAddress Windows10EasEmailProfileConfigurationUsernameAADSource = "primarySmtpAddress"
	Windows10EasEmailProfileConfigurationUsernameAADSource_SamAccountName     Windows10EasEmailProfileConfigurationUsernameAADSource = "samAccountName"
	Windows10EasEmailProfileConfigurationUsernameAADSource_UserPrincipalName  Windows10EasEmailProfileConfigurationUsernameAADSource = "userPrincipalName"
)

func PossibleValuesForWindows10EasEmailProfileConfigurationUsernameAADSource() []string {
	return []string{
		string(Windows10EasEmailProfileConfigurationUsernameAADSource_PrimarySmtpAddress),
		string(Windows10EasEmailProfileConfigurationUsernameAADSource_SamAccountName),
		string(Windows10EasEmailProfileConfigurationUsernameAADSource_UserPrincipalName),
	}
}

func (s *Windows10EasEmailProfileConfigurationUsernameAADSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EasEmailProfileConfigurationUsernameAADSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EasEmailProfileConfigurationUsernameAADSource(input string) (*Windows10EasEmailProfileConfigurationUsernameAADSource, error) {
	vals := map[string]Windows10EasEmailProfileConfigurationUsernameAADSource{
		"primarysmtpaddress": Windows10EasEmailProfileConfigurationUsernameAADSource_PrimarySmtpAddress,
		"samaccountname":     Windows10EasEmailProfileConfigurationUsernameAADSource_SamAccountName,
		"userprincipalname":  Windows10EasEmailProfileConfigurationUsernameAADSource_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EasEmailProfileConfigurationUsernameAADSource(input)
	return &out, nil
}
