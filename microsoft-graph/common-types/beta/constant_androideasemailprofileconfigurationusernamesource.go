package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidEasEmailProfileConfigurationUsernameSource string

const (
	AndroidEasEmailProfileConfigurationUsernameSource_PrimarySmtpAddress AndroidEasEmailProfileConfigurationUsernameSource = "primarySmtpAddress"
	AndroidEasEmailProfileConfigurationUsernameSource_SamAccountName     AndroidEasEmailProfileConfigurationUsernameSource = "samAccountName"
	AndroidEasEmailProfileConfigurationUsernameSource_UserPrincipalName  AndroidEasEmailProfileConfigurationUsernameSource = "userPrincipalName"
	AndroidEasEmailProfileConfigurationUsernameSource_Username           AndroidEasEmailProfileConfigurationUsernameSource = "username"
)

func PossibleValuesForAndroidEasEmailProfileConfigurationUsernameSource() []string {
	return []string{
		string(AndroidEasEmailProfileConfigurationUsernameSource_PrimarySmtpAddress),
		string(AndroidEasEmailProfileConfigurationUsernameSource_SamAccountName),
		string(AndroidEasEmailProfileConfigurationUsernameSource_UserPrincipalName),
		string(AndroidEasEmailProfileConfigurationUsernameSource_Username),
	}
}

func (s *AndroidEasEmailProfileConfigurationUsernameSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidEasEmailProfileConfigurationUsernameSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidEasEmailProfileConfigurationUsernameSource(input string) (*AndroidEasEmailProfileConfigurationUsernameSource, error) {
	vals := map[string]AndroidEasEmailProfileConfigurationUsernameSource{
		"primarysmtpaddress": AndroidEasEmailProfileConfigurationUsernameSource_PrimarySmtpAddress,
		"samaccountname":     AndroidEasEmailProfileConfigurationUsernameSource_SamAccountName,
		"userprincipalname":  AndroidEasEmailProfileConfigurationUsernameSource_UserPrincipalName,
		"username":           AndroidEasEmailProfileConfigurationUsernameSource_Username,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidEasEmailProfileConfigurationUsernameSource(input)
	return &out, nil
}
