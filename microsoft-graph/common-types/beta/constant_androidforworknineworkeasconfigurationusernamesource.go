package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkNineWorkEasConfigurationUsernameSource string

const (
	AndroidForWorkNineWorkEasConfigurationUsernameSource_PrimarySmtpAddress AndroidForWorkNineWorkEasConfigurationUsernameSource = "primarySmtpAddress"
	AndroidForWorkNineWorkEasConfigurationUsernameSource_SamAccountName     AndroidForWorkNineWorkEasConfigurationUsernameSource = "samAccountName"
	AndroidForWorkNineWorkEasConfigurationUsernameSource_UserPrincipalName  AndroidForWorkNineWorkEasConfigurationUsernameSource = "userPrincipalName"
	AndroidForWorkNineWorkEasConfigurationUsernameSource_Username           AndroidForWorkNineWorkEasConfigurationUsernameSource = "username"
)

func PossibleValuesForAndroidForWorkNineWorkEasConfigurationUsernameSource() []string {
	return []string{
		string(AndroidForWorkNineWorkEasConfigurationUsernameSource_PrimarySmtpAddress),
		string(AndroidForWorkNineWorkEasConfigurationUsernameSource_SamAccountName),
		string(AndroidForWorkNineWorkEasConfigurationUsernameSource_UserPrincipalName),
		string(AndroidForWorkNineWorkEasConfigurationUsernameSource_Username),
	}
}

func (s *AndroidForWorkNineWorkEasConfigurationUsernameSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkNineWorkEasConfigurationUsernameSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkNineWorkEasConfigurationUsernameSource(input string) (*AndroidForWorkNineWorkEasConfigurationUsernameSource, error) {
	vals := map[string]AndroidForWorkNineWorkEasConfigurationUsernameSource{
		"primarysmtpaddress": AndroidForWorkNineWorkEasConfigurationUsernameSource_PrimarySmtpAddress,
		"samaccountname":     AndroidForWorkNineWorkEasConfigurationUsernameSource_SamAccountName,
		"userprincipalname":  AndroidForWorkNineWorkEasConfigurationUsernameSource_UserPrincipalName,
		"username":           AndroidForWorkNineWorkEasConfigurationUsernameSource_Username,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkNineWorkEasConfigurationUsernameSource(input)
	return &out, nil
}
