package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SmsAuthenticationMethodConfigurationState string

const (
	SmsAuthenticationMethodConfigurationState_Disabled SmsAuthenticationMethodConfigurationState = "disabled"
	SmsAuthenticationMethodConfigurationState_Enabled  SmsAuthenticationMethodConfigurationState = "enabled"
)

func PossibleValuesForSmsAuthenticationMethodConfigurationState() []string {
	return []string{
		string(SmsAuthenticationMethodConfigurationState_Disabled),
		string(SmsAuthenticationMethodConfigurationState_Enabled),
	}
}

func (s *SmsAuthenticationMethodConfigurationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSmsAuthenticationMethodConfigurationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSmsAuthenticationMethodConfigurationState(input string) (*SmsAuthenticationMethodConfigurationState, error) {
	vals := map[string]SmsAuthenticationMethodConfigurationState{
		"disabled": SmsAuthenticationMethodConfigurationState_Disabled,
		"enabled":  SmsAuthenticationMethodConfigurationState_Enabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SmsAuthenticationMethodConfigurationState(input)
	return &out, nil
}
