package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidEasEmailProfileConfigurationEmailAddressSource string

const (
	AndroidEasEmailProfileConfigurationEmailAddressSource_PrimarySmtpAddress AndroidEasEmailProfileConfigurationEmailAddressSource = "primarySmtpAddress"
	AndroidEasEmailProfileConfigurationEmailAddressSource_UserPrincipalName  AndroidEasEmailProfileConfigurationEmailAddressSource = "userPrincipalName"
)

func PossibleValuesForAndroidEasEmailProfileConfigurationEmailAddressSource() []string {
	return []string{
		string(AndroidEasEmailProfileConfigurationEmailAddressSource_PrimarySmtpAddress),
		string(AndroidEasEmailProfileConfigurationEmailAddressSource_UserPrincipalName),
	}
}

func (s *AndroidEasEmailProfileConfigurationEmailAddressSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidEasEmailProfileConfigurationEmailAddressSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidEasEmailProfileConfigurationEmailAddressSource(input string) (*AndroidEasEmailProfileConfigurationEmailAddressSource, error) {
	vals := map[string]AndroidEasEmailProfileConfigurationEmailAddressSource{
		"primarysmtpaddress": AndroidEasEmailProfileConfigurationEmailAddressSource_PrimarySmtpAddress,
		"userprincipalname":  AndroidEasEmailProfileConfigurationEmailAddressSource_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidEasEmailProfileConfigurationEmailAddressSource(input)
	return &out, nil
}
