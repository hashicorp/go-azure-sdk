package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileEasEmailProfileBaseUsernameSource string

const (
	AndroidWorkProfileEasEmailProfileBaseUsernameSource_PrimarySmtpAddress AndroidWorkProfileEasEmailProfileBaseUsernameSource = "primarySmtpAddress"
	AndroidWorkProfileEasEmailProfileBaseUsernameSource_SamAccountName     AndroidWorkProfileEasEmailProfileBaseUsernameSource = "samAccountName"
	AndroidWorkProfileEasEmailProfileBaseUsernameSource_UserPrincipalName  AndroidWorkProfileEasEmailProfileBaseUsernameSource = "userPrincipalName"
	AndroidWorkProfileEasEmailProfileBaseUsernameSource_Username           AndroidWorkProfileEasEmailProfileBaseUsernameSource = "username"
)

func PossibleValuesForAndroidWorkProfileEasEmailProfileBaseUsernameSource() []string {
	return []string{
		string(AndroidWorkProfileEasEmailProfileBaseUsernameSource_PrimarySmtpAddress),
		string(AndroidWorkProfileEasEmailProfileBaseUsernameSource_SamAccountName),
		string(AndroidWorkProfileEasEmailProfileBaseUsernameSource_UserPrincipalName),
		string(AndroidWorkProfileEasEmailProfileBaseUsernameSource_Username),
	}
}

func (s *AndroidWorkProfileEasEmailProfileBaseUsernameSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWorkProfileEasEmailProfileBaseUsernameSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWorkProfileEasEmailProfileBaseUsernameSource(input string) (*AndroidWorkProfileEasEmailProfileBaseUsernameSource, error) {
	vals := map[string]AndroidWorkProfileEasEmailProfileBaseUsernameSource{
		"primarysmtpaddress": AndroidWorkProfileEasEmailProfileBaseUsernameSource_PrimarySmtpAddress,
		"samaccountname":     AndroidWorkProfileEasEmailProfileBaseUsernameSource_SamAccountName,
		"userprincipalname":  AndroidWorkProfileEasEmailProfileBaseUsernameSource_UserPrincipalName,
		"username":           AndroidWorkProfileEasEmailProfileBaseUsernameSource_Username,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileEasEmailProfileBaseUsernameSource(input)
	return &out, nil
}
