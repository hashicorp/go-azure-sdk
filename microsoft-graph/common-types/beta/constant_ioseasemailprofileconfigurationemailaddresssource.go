package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosEasEmailProfileConfigurationEmailAddressSource string

const (
	IosEasEmailProfileConfigurationEmailAddressSource_PrimarySmtpAddress IosEasEmailProfileConfigurationEmailAddressSource = "primarySmtpAddress"
	IosEasEmailProfileConfigurationEmailAddressSource_UserPrincipalName  IosEasEmailProfileConfigurationEmailAddressSource = "userPrincipalName"
)

func PossibleValuesForIosEasEmailProfileConfigurationEmailAddressSource() []string {
	return []string{
		string(IosEasEmailProfileConfigurationEmailAddressSource_PrimarySmtpAddress),
		string(IosEasEmailProfileConfigurationEmailAddressSource_UserPrincipalName),
	}
}

func (s *IosEasEmailProfileConfigurationEmailAddressSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosEasEmailProfileConfigurationEmailAddressSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosEasEmailProfileConfigurationEmailAddressSource(input string) (*IosEasEmailProfileConfigurationEmailAddressSource, error) {
	vals := map[string]IosEasEmailProfileConfigurationEmailAddressSource{
		"primarysmtpaddress": IosEasEmailProfileConfigurationEmailAddressSource_PrimarySmtpAddress,
		"userprincipalname":  IosEasEmailProfileConfigurationEmailAddressSource_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosEasEmailProfileConfigurationEmailAddressSource(input)
	return &out, nil
}
