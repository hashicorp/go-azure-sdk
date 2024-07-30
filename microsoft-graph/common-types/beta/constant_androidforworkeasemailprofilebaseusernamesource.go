package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkEasEmailProfileBaseUsernameSource string

const (
	AndroidForWorkEasEmailProfileBaseUsernameSource_PrimarySmtpAddress AndroidForWorkEasEmailProfileBaseUsernameSource = "primarySmtpAddress"
	AndroidForWorkEasEmailProfileBaseUsernameSource_SamAccountName     AndroidForWorkEasEmailProfileBaseUsernameSource = "samAccountName"
	AndroidForWorkEasEmailProfileBaseUsernameSource_UserPrincipalName  AndroidForWorkEasEmailProfileBaseUsernameSource = "userPrincipalName"
	AndroidForWorkEasEmailProfileBaseUsernameSource_Username           AndroidForWorkEasEmailProfileBaseUsernameSource = "username"
)

func PossibleValuesForAndroidForWorkEasEmailProfileBaseUsernameSource() []string {
	return []string{
		string(AndroidForWorkEasEmailProfileBaseUsernameSource_PrimarySmtpAddress),
		string(AndroidForWorkEasEmailProfileBaseUsernameSource_SamAccountName),
		string(AndroidForWorkEasEmailProfileBaseUsernameSource_UserPrincipalName),
		string(AndroidForWorkEasEmailProfileBaseUsernameSource_Username),
	}
}

func (s *AndroidForWorkEasEmailProfileBaseUsernameSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkEasEmailProfileBaseUsernameSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkEasEmailProfileBaseUsernameSource(input string) (*AndroidForWorkEasEmailProfileBaseUsernameSource, error) {
	vals := map[string]AndroidForWorkEasEmailProfileBaseUsernameSource{
		"primarysmtpaddress": AndroidForWorkEasEmailProfileBaseUsernameSource_PrimarySmtpAddress,
		"samaccountname":     AndroidForWorkEasEmailProfileBaseUsernameSource_SamAccountName,
		"userprincipalname":  AndroidForWorkEasEmailProfileBaseUsernameSource_UserPrincipalName,
		"username":           AndroidForWorkEasEmailProfileBaseUsernameSource_Username,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkEasEmailProfileBaseUsernameSource(input)
	return &out, nil
}
