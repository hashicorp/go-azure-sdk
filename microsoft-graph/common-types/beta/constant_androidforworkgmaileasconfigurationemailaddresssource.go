package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkGmailEasConfigurationEmailAddressSource string

const (
	AndroidForWorkGmailEasConfigurationEmailAddressSource_PrimarySmtpAddress AndroidForWorkGmailEasConfigurationEmailAddressSource = "primarySmtpAddress"
	AndroidForWorkGmailEasConfigurationEmailAddressSource_UserPrincipalName  AndroidForWorkGmailEasConfigurationEmailAddressSource = "userPrincipalName"
)

func PossibleValuesForAndroidForWorkGmailEasConfigurationEmailAddressSource() []string {
	return []string{
		string(AndroidForWorkGmailEasConfigurationEmailAddressSource_PrimarySmtpAddress),
		string(AndroidForWorkGmailEasConfigurationEmailAddressSource_UserPrincipalName),
	}
}

func (s *AndroidForWorkGmailEasConfigurationEmailAddressSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkGmailEasConfigurationEmailAddressSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkGmailEasConfigurationEmailAddressSource(input string) (*AndroidForWorkGmailEasConfigurationEmailAddressSource, error) {
	vals := map[string]AndroidForWorkGmailEasConfigurationEmailAddressSource{
		"primarysmtpaddress": AndroidForWorkGmailEasConfigurationEmailAddressSource_PrimarySmtpAddress,
		"userprincipalname":  AndroidForWorkGmailEasConfigurationEmailAddressSource_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkGmailEasConfigurationEmailAddressSource(input)
	return &out, nil
}
