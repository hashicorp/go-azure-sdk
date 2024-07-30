package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EasEmailProfileConfigurationBaseUsernameAADSource string

const (
	EasEmailProfileConfigurationBaseUsernameAADSource_PrimarySmtpAddress EasEmailProfileConfigurationBaseUsernameAADSource = "primarySmtpAddress"
	EasEmailProfileConfigurationBaseUsernameAADSource_SamAccountName     EasEmailProfileConfigurationBaseUsernameAADSource = "samAccountName"
	EasEmailProfileConfigurationBaseUsernameAADSource_UserPrincipalName  EasEmailProfileConfigurationBaseUsernameAADSource = "userPrincipalName"
)

func PossibleValuesForEasEmailProfileConfigurationBaseUsernameAADSource() []string {
	return []string{
		string(EasEmailProfileConfigurationBaseUsernameAADSource_PrimarySmtpAddress),
		string(EasEmailProfileConfigurationBaseUsernameAADSource_SamAccountName),
		string(EasEmailProfileConfigurationBaseUsernameAADSource_UserPrincipalName),
	}
}

func (s *EasEmailProfileConfigurationBaseUsernameAADSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEasEmailProfileConfigurationBaseUsernameAADSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEasEmailProfileConfigurationBaseUsernameAADSource(input string) (*EasEmailProfileConfigurationBaseUsernameAADSource, error) {
	vals := map[string]EasEmailProfileConfigurationBaseUsernameAADSource{
		"primarysmtpaddress": EasEmailProfileConfigurationBaseUsernameAADSource_PrimarySmtpAddress,
		"samaccountname":     EasEmailProfileConfigurationBaseUsernameAADSource_SamAccountName,
		"userprincipalname":  EasEmailProfileConfigurationBaseUsernameAADSource_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EasEmailProfileConfigurationBaseUsernameAADSource(input)
	return &out, nil
}
