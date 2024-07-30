package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SmsAuthenticationMethodTargetTargetType string

const (
	SmsAuthenticationMethodTargetTargetType_Group SmsAuthenticationMethodTargetTargetType = "group"
	SmsAuthenticationMethodTargetTargetType_User  SmsAuthenticationMethodTargetTargetType = "user"
)

func PossibleValuesForSmsAuthenticationMethodTargetTargetType() []string {
	return []string{
		string(SmsAuthenticationMethodTargetTargetType_Group),
		string(SmsAuthenticationMethodTargetTargetType_User),
	}
}

func (s *SmsAuthenticationMethodTargetTargetType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSmsAuthenticationMethodTargetTargetType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSmsAuthenticationMethodTargetTargetType(input string) (*SmsAuthenticationMethodTargetTargetType, error) {
	vals := map[string]SmsAuthenticationMethodTargetTargetType{
		"group": SmsAuthenticationMethodTargetTargetType_Group,
		"user":  SmsAuthenticationMethodTargetTargetType_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SmsAuthenticationMethodTargetTargetType(input)
	return &out, nil
}
