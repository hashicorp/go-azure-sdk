package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EasEmailProfileConfigurationEmailAddressSource string

const (
	Windows10EasEmailProfileConfigurationEmailAddressSource_PrimarySmtpAddress Windows10EasEmailProfileConfigurationEmailAddressSource = "primarySmtpAddress"
	Windows10EasEmailProfileConfigurationEmailAddressSource_UserPrincipalName  Windows10EasEmailProfileConfigurationEmailAddressSource = "userPrincipalName"
)

func PossibleValuesForWindows10EasEmailProfileConfigurationEmailAddressSource() []string {
	return []string{
		string(Windows10EasEmailProfileConfigurationEmailAddressSource_PrimarySmtpAddress),
		string(Windows10EasEmailProfileConfigurationEmailAddressSource_UserPrincipalName),
	}
}

func (s *Windows10EasEmailProfileConfigurationEmailAddressSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EasEmailProfileConfigurationEmailAddressSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EasEmailProfileConfigurationEmailAddressSource(input string) (*Windows10EasEmailProfileConfigurationEmailAddressSource, error) {
	vals := map[string]Windows10EasEmailProfileConfigurationEmailAddressSource{
		"primarysmtpaddress": Windows10EasEmailProfileConfigurationEmailAddressSource_PrimarySmtpAddress,
		"userprincipalname":  Windows10EasEmailProfileConfigurationEmailAddressSource_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EasEmailProfileConfigurationEmailAddressSource(input)
	return &out, nil
}
