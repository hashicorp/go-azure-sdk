package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileNineWorkEasConfigurationUsernameSource string

const (
	AndroidWorkProfileNineWorkEasConfigurationUsernameSource_PrimarySmtpAddress AndroidWorkProfileNineWorkEasConfigurationUsernameSource = "primarySmtpAddress"
	AndroidWorkProfileNineWorkEasConfigurationUsernameSource_SamAccountName     AndroidWorkProfileNineWorkEasConfigurationUsernameSource = "samAccountName"
	AndroidWorkProfileNineWorkEasConfigurationUsernameSource_UserPrincipalName  AndroidWorkProfileNineWorkEasConfigurationUsernameSource = "userPrincipalName"
	AndroidWorkProfileNineWorkEasConfigurationUsernameSource_Username           AndroidWorkProfileNineWorkEasConfigurationUsernameSource = "username"
)

func PossibleValuesForAndroidWorkProfileNineWorkEasConfigurationUsernameSource() []string {
	return []string{
		string(AndroidWorkProfileNineWorkEasConfigurationUsernameSource_PrimarySmtpAddress),
		string(AndroidWorkProfileNineWorkEasConfigurationUsernameSource_SamAccountName),
		string(AndroidWorkProfileNineWorkEasConfigurationUsernameSource_UserPrincipalName),
		string(AndroidWorkProfileNineWorkEasConfigurationUsernameSource_Username),
	}
}

func (s *AndroidWorkProfileNineWorkEasConfigurationUsernameSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWorkProfileNineWorkEasConfigurationUsernameSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWorkProfileNineWorkEasConfigurationUsernameSource(input string) (*AndroidWorkProfileNineWorkEasConfigurationUsernameSource, error) {
	vals := map[string]AndroidWorkProfileNineWorkEasConfigurationUsernameSource{
		"primarysmtpaddress": AndroidWorkProfileNineWorkEasConfigurationUsernameSource_PrimarySmtpAddress,
		"samaccountname":     AndroidWorkProfileNineWorkEasConfigurationUsernameSource_SamAccountName,
		"userprincipalname":  AndroidWorkProfileNineWorkEasConfigurationUsernameSource_UserPrincipalName,
		"username":           AndroidWorkProfileNineWorkEasConfigurationUsernameSource_Username,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileNineWorkEasConfigurationUsernameSource(input)
	return &out, nil
}
