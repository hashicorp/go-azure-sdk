package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkNineWorkEasConfigurationEmailAddressSource string

const (
	AndroidForWorkNineWorkEasConfigurationEmailAddressSource_PrimarySmtpAddress AndroidForWorkNineWorkEasConfigurationEmailAddressSource = "primarySmtpAddress"
	AndroidForWorkNineWorkEasConfigurationEmailAddressSource_UserPrincipalName  AndroidForWorkNineWorkEasConfigurationEmailAddressSource = "userPrincipalName"
)

func PossibleValuesForAndroidForWorkNineWorkEasConfigurationEmailAddressSource() []string {
	return []string{
		string(AndroidForWorkNineWorkEasConfigurationEmailAddressSource_PrimarySmtpAddress),
		string(AndroidForWorkNineWorkEasConfigurationEmailAddressSource_UserPrincipalName),
	}
}

func (s *AndroidForWorkNineWorkEasConfigurationEmailAddressSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkNineWorkEasConfigurationEmailAddressSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkNineWorkEasConfigurationEmailAddressSource(input string) (*AndroidForWorkNineWorkEasConfigurationEmailAddressSource, error) {
	vals := map[string]AndroidForWorkNineWorkEasConfigurationEmailAddressSource{
		"primarysmtpaddress": AndroidForWorkNineWorkEasConfigurationEmailAddressSource_PrimarySmtpAddress,
		"userprincipalname":  AndroidForWorkNineWorkEasConfigurationEmailAddressSource_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkNineWorkEasConfigurationEmailAddressSource(input)
	return &out, nil
}
