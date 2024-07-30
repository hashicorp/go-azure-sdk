package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosEasEmailProfileConfigurationUsernameSource string

const (
	IosEasEmailProfileConfigurationUsernameSource_PrimarySmtpAddress IosEasEmailProfileConfigurationUsernameSource = "primarySmtpAddress"
	IosEasEmailProfileConfigurationUsernameSource_UserPrincipalName  IosEasEmailProfileConfigurationUsernameSource = "userPrincipalName"
)

func PossibleValuesForIosEasEmailProfileConfigurationUsernameSource() []string {
	return []string{
		string(IosEasEmailProfileConfigurationUsernameSource_PrimarySmtpAddress),
		string(IosEasEmailProfileConfigurationUsernameSource_UserPrincipalName),
	}
}

func (s *IosEasEmailProfileConfigurationUsernameSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosEasEmailProfileConfigurationUsernameSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosEasEmailProfileConfigurationUsernameSource(input string) (*IosEasEmailProfileConfigurationUsernameSource, error) {
	vals := map[string]IosEasEmailProfileConfigurationUsernameSource{
		"primarysmtpaddress": IosEasEmailProfileConfigurationUsernameSource_PrimarySmtpAddress,
		"userprincipalname":  IosEasEmailProfileConfigurationUsernameSource_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosEasEmailProfileConfigurationUsernameSource(input)
	return &out, nil
}
