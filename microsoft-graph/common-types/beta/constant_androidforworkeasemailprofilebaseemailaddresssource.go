package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkEasEmailProfileBaseEmailAddressSource string

const (
	AndroidForWorkEasEmailProfileBaseEmailAddressSource_PrimarySmtpAddress AndroidForWorkEasEmailProfileBaseEmailAddressSource = "primarySmtpAddress"
	AndroidForWorkEasEmailProfileBaseEmailAddressSource_UserPrincipalName  AndroidForWorkEasEmailProfileBaseEmailAddressSource = "userPrincipalName"
)

func PossibleValuesForAndroidForWorkEasEmailProfileBaseEmailAddressSource() []string {
	return []string{
		string(AndroidForWorkEasEmailProfileBaseEmailAddressSource_PrimarySmtpAddress),
		string(AndroidForWorkEasEmailProfileBaseEmailAddressSource_UserPrincipalName),
	}
}

func (s *AndroidForWorkEasEmailProfileBaseEmailAddressSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkEasEmailProfileBaseEmailAddressSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkEasEmailProfileBaseEmailAddressSource(input string) (*AndroidForWorkEasEmailProfileBaseEmailAddressSource, error) {
	vals := map[string]AndroidForWorkEasEmailProfileBaseEmailAddressSource{
		"primarysmtpaddress": AndroidForWorkEasEmailProfileBaseEmailAddressSource_PrimarySmtpAddress,
		"userprincipalname":  AndroidForWorkEasEmailProfileBaseEmailAddressSource_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkEasEmailProfileBaseEmailAddressSource(input)
	return &out, nil
}
