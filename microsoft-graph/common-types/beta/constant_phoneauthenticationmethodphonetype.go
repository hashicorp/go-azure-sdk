package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PhoneAuthenticationMethodPhoneType string

const (
	PhoneAuthenticationMethodPhoneType_AlternateMobile PhoneAuthenticationMethodPhoneType = "alternateMobile"
	PhoneAuthenticationMethodPhoneType_Mobile          PhoneAuthenticationMethodPhoneType = "mobile"
	PhoneAuthenticationMethodPhoneType_Office          PhoneAuthenticationMethodPhoneType = "office"
)

func PossibleValuesForPhoneAuthenticationMethodPhoneType() []string {
	return []string{
		string(PhoneAuthenticationMethodPhoneType_AlternateMobile),
		string(PhoneAuthenticationMethodPhoneType_Mobile),
		string(PhoneAuthenticationMethodPhoneType_Office),
	}
}

func (s *PhoneAuthenticationMethodPhoneType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePhoneAuthenticationMethodPhoneType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePhoneAuthenticationMethodPhoneType(input string) (*PhoneAuthenticationMethodPhoneType, error) {
	vals := map[string]PhoneAuthenticationMethodPhoneType{
		"alternatemobile": PhoneAuthenticationMethodPhoneType_AlternateMobile,
		"mobile":          PhoneAuthenticationMethodPhoneType_Mobile,
		"office":          PhoneAuthenticationMethodPhoneType_Office,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PhoneAuthenticationMethodPhoneType(input)
	return &out, nil
}
