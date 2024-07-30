package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileEasEmailProfileBaseEmailAddressSource string

const (
	AndroidWorkProfileEasEmailProfileBaseEmailAddressSource_PrimarySmtpAddress AndroidWorkProfileEasEmailProfileBaseEmailAddressSource = "primarySmtpAddress"
	AndroidWorkProfileEasEmailProfileBaseEmailAddressSource_UserPrincipalName  AndroidWorkProfileEasEmailProfileBaseEmailAddressSource = "userPrincipalName"
)

func PossibleValuesForAndroidWorkProfileEasEmailProfileBaseEmailAddressSource() []string {
	return []string{
		string(AndroidWorkProfileEasEmailProfileBaseEmailAddressSource_PrimarySmtpAddress),
		string(AndroidWorkProfileEasEmailProfileBaseEmailAddressSource_UserPrincipalName),
	}
}

func (s *AndroidWorkProfileEasEmailProfileBaseEmailAddressSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWorkProfileEasEmailProfileBaseEmailAddressSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWorkProfileEasEmailProfileBaseEmailAddressSource(input string) (*AndroidWorkProfileEasEmailProfileBaseEmailAddressSource, error) {
	vals := map[string]AndroidWorkProfileEasEmailProfileBaseEmailAddressSource{
		"primarysmtpaddress": AndroidWorkProfileEasEmailProfileBaseEmailAddressSource_PrimarySmtpAddress,
		"userprincipalname":  AndroidWorkProfileEasEmailProfileBaseEmailAddressSource_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileEasEmailProfileBaseEmailAddressSource(input)
	return &out, nil
}
