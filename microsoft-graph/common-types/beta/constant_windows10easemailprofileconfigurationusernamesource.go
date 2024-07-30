package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EasEmailProfileConfigurationUsernameSource string

const (
	Windows10EasEmailProfileConfigurationUsernameSource_PrimarySmtpAddress Windows10EasEmailProfileConfigurationUsernameSource = "primarySmtpAddress"
	Windows10EasEmailProfileConfigurationUsernameSource_UserPrincipalName  Windows10EasEmailProfileConfigurationUsernameSource = "userPrincipalName"
)

func PossibleValuesForWindows10EasEmailProfileConfigurationUsernameSource() []string {
	return []string{
		string(Windows10EasEmailProfileConfigurationUsernameSource_PrimarySmtpAddress),
		string(Windows10EasEmailProfileConfigurationUsernameSource_UserPrincipalName),
	}
}

func (s *Windows10EasEmailProfileConfigurationUsernameSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EasEmailProfileConfigurationUsernameSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EasEmailProfileConfigurationUsernameSource(input string) (*Windows10EasEmailProfileConfigurationUsernameSource, error) {
	vals := map[string]Windows10EasEmailProfileConfigurationUsernameSource{
		"primarysmtpaddress": Windows10EasEmailProfileConfigurationUsernameSource_PrimarySmtpAddress,
		"userprincipalname":  Windows10EasEmailProfileConfigurationUsernameSource_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EasEmailProfileConfigurationUsernameSource(input)
	return &out, nil
}
