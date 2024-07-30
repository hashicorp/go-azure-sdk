package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileGmailEasConfigurationUsernameSource string

const (
	AndroidWorkProfileGmailEasConfigurationUsernameSource_PrimarySmtpAddress AndroidWorkProfileGmailEasConfigurationUsernameSource = "primarySmtpAddress"
	AndroidWorkProfileGmailEasConfigurationUsernameSource_SamAccountName     AndroidWorkProfileGmailEasConfigurationUsernameSource = "samAccountName"
	AndroidWorkProfileGmailEasConfigurationUsernameSource_UserPrincipalName  AndroidWorkProfileGmailEasConfigurationUsernameSource = "userPrincipalName"
	AndroidWorkProfileGmailEasConfigurationUsernameSource_Username           AndroidWorkProfileGmailEasConfigurationUsernameSource = "username"
)

func PossibleValuesForAndroidWorkProfileGmailEasConfigurationUsernameSource() []string {
	return []string{
		string(AndroidWorkProfileGmailEasConfigurationUsernameSource_PrimarySmtpAddress),
		string(AndroidWorkProfileGmailEasConfigurationUsernameSource_SamAccountName),
		string(AndroidWorkProfileGmailEasConfigurationUsernameSource_UserPrincipalName),
		string(AndroidWorkProfileGmailEasConfigurationUsernameSource_Username),
	}
}

func (s *AndroidWorkProfileGmailEasConfigurationUsernameSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWorkProfileGmailEasConfigurationUsernameSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWorkProfileGmailEasConfigurationUsernameSource(input string) (*AndroidWorkProfileGmailEasConfigurationUsernameSource, error) {
	vals := map[string]AndroidWorkProfileGmailEasConfigurationUsernameSource{
		"primarysmtpaddress": AndroidWorkProfileGmailEasConfigurationUsernameSource_PrimarySmtpAddress,
		"samaccountname":     AndroidWorkProfileGmailEasConfigurationUsernameSource_SamAccountName,
		"userprincipalname":  AndroidWorkProfileGmailEasConfigurationUsernameSource_UserPrincipalName,
		"username":           AndroidWorkProfileGmailEasConfigurationUsernameSource_Username,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileGmailEasConfigurationUsernameSource(input)
	return &out, nil
}
