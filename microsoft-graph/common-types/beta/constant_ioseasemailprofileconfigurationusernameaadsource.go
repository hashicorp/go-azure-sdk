package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosEasEmailProfileConfigurationUsernameAADSource string

const (
	IosEasEmailProfileConfigurationUsernameAADSource_PrimarySmtpAddress IosEasEmailProfileConfigurationUsernameAADSource = "primarySmtpAddress"
	IosEasEmailProfileConfigurationUsernameAADSource_SamAccountName     IosEasEmailProfileConfigurationUsernameAADSource = "samAccountName"
	IosEasEmailProfileConfigurationUsernameAADSource_UserPrincipalName  IosEasEmailProfileConfigurationUsernameAADSource = "userPrincipalName"
)

func PossibleValuesForIosEasEmailProfileConfigurationUsernameAADSource() []string {
	return []string{
		string(IosEasEmailProfileConfigurationUsernameAADSource_PrimarySmtpAddress),
		string(IosEasEmailProfileConfigurationUsernameAADSource_SamAccountName),
		string(IosEasEmailProfileConfigurationUsernameAADSource_UserPrincipalName),
	}
}

func (s *IosEasEmailProfileConfigurationUsernameAADSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosEasEmailProfileConfigurationUsernameAADSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosEasEmailProfileConfigurationUsernameAADSource(input string) (*IosEasEmailProfileConfigurationUsernameAADSource, error) {
	vals := map[string]IosEasEmailProfileConfigurationUsernameAADSource{
		"primarysmtpaddress": IosEasEmailProfileConfigurationUsernameAADSource_PrimarySmtpAddress,
		"samaccountname":     IosEasEmailProfileConfigurationUsernameAADSource_SamAccountName,
		"userprincipalname":  IosEasEmailProfileConfigurationUsernameAADSource_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosEasEmailProfileConfigurationUsernameAADSource(input)
	return &out, nil
}
