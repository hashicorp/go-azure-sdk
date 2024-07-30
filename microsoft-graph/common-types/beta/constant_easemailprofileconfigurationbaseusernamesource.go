package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EasEmailProfileConfigurationBaseUsernameSource string

const (
	EasEmailProfileConfigurationBaseUsernameSource_PrimarySmtpAddress EasEmailProfileConfigurationBaseUsernameSource = "primarySmtpAddress"
	EasEmailProfileConfigurationBaseUsernameSource_UserPrincipalName  EasEmailProfileConfigurationBaseUsernameSource = "userPrincipalName"
)

func PossibleValuesForEasEmailProfileConfigurationBaseUsernameSource() []string {
	return []string{
		string(EasEmailProfileConfigurationBaseUsernameSource_PrimarySmtpAddress),
		string(EasEmailProfileConfigurationBaseUsernameSource_UserPrincipalName),
	}
}

func (s *EasEmailProfileConfigurationBaseUsernameSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEasEmailProfileConfigurationBaseUsernameSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEasEmailProfileConfigurationBaseUsernameSource(input string) (*EasEmailProfileConfigurationBaseUsernameSource, error) {
	vals := map[string]EasEmailProfileConfigurationBaseUsernameSource{
		"primarysmtpaddress": EasEmailProfileConfigurationBaseUsernameSource_PrimarySmtpAddress,
		"userprincipalname":  EasEmailProfileConfigurationBaseUsernameSource_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EasEmailProfileConfigurationBaseUsernameSource(input)
	return &out, nil
}
