package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileGmailEasConfigurationEmailAddressSource string

const (
	AndroidWorkProfileGmailEasConfigurationEmailAddressSource_PrimarySmtpAddress AndroidWorkProfileGmailEasConfigurationEmailAddressSource = "primarySmtpAddress"
	AndroidWorkProfileGmailEasConfigurationEmailAddressSource_UserPrincipalName  AndroidWorkProfileGmailEasConfigurationEmailAddressSource = "userPrincipalName"
)

func PossibleValuesForAndroidWorkProfileGmailEasConfigurationEmailAddressSource() []string {
	return []string{
		string(AndroidWorkProfileGmailEasConfigurationEmailAddressSource_PrimarySmtpAddress),
		string(AndroidWorkProfileGmailEasConfigurationEmailAddressSource_UserPrincipalName),
	}
}

func (s *AndroidWorkProfileGmailEasConfigurationEmailAddressSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWorkProfileGmailEasConfigurationEmailAddressSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWorkProfileGmailEasConfigurationEmailAddressSource(input string) (*AndroidWorkProfileGmailEasConfigurationEmailAddressSource, error) {
	vals := map[string]AndroidWorkProfileGmailEasConfigurationEmailAddressSource{
		"primarysmtpaddress": AndroidWorkProfileGmailEasConfigurationEmailAddressSource_PrimarySmtpAddress,
		"userprincipalname":  AndroidWorkProfileGmailEasConfigurationEmailAddressSource_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileGmailEasConfigurationEmailAddressSource(input)
	return &out, nil
}
