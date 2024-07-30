package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkGmailEasConfigurationUsernameSource string

const (
	AndroidForWorkGmailEasConfigurationUsernameSource_PrimarySmtpAddress AndroidForWorkGmailEasConfigurationUsernameSource = "primarySmtpAddress"
	AndroidForWorkGmailEasConfigurationUsernameSource_SamAccountName     AndroidForWorkGmailEasConfigurationUsernameSource = "samAccountName"
	AndroidForWorkGmailEasConfigurationUsernameSource_UserPrincipalName  AndroidForWorkGmailEasConfigurationUsernameSource = "userPrincipalName"
	AndroidForWorkGmailEasConfigurationUsernameSource_Username           AndroidForWorkGmailEasConfigurationUsernameSource = "username"
)

func PossibleValuesForAndroidForWorkGmailEasConfigurationUsernameSource() []string {
	return []string{
		string(AndroidForWorkGmailEasConfigurationUsernameSource_PrimarySmtpAddress),
		string(AndroidForWorkGmailEasConfigurationUsernameSource_SamAccountName),
		string(AndroidForWorkGmailEasConfigurationUsernameSource_UserPrincipalName),
		string(AndroidForWorkGmailEasConfigurationUsernameSource_Username),
	}
}

func (s *AndroidForWorkGmailEasConfigurationUsernameSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkGmailEasConfigurationUsernameSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkGmailEasConfigurationUsernameSource(input string) (*AndroidForWorkGmailEasConfigurationUsernameSource, error) {
	vals := map[string]AndroidForWorkGmailEasConfigurationUsernameSource{
		"primarysmtpaddress": AndroidForWorkGmailEasConfigurationUsernameSource_PrimarySmtpAddress,
		"samaccountname":     AndroidForWorkGmailEasConfigurationUsernameSource_SamAccountName,
		"userprincipalname":  AndroidForWorkGmailEasConfigurationUsernameSource_UserPrincipalName,
		"username":           AndroidForWorkGmailEasConfigurationUsernameSource_Username,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkGmailEasConfigurationUsernameSource(input)
	return &out, nil
}
