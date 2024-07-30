package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmailAuthenticationMethodConfigurationState string

const (
	EmailAuthenticationMethodConfigurationState_Disabled EmailAuthenticationMethodConfigurationState = "disabled"
	EmailAuthenticationMethodConfigurationState_Enabled  EmailAuthenticationMethodConfigurationState = "enabled"
)

func PossibleValuesForEmailAuthenticationMethodConfigurationState() []string {
	return []string{
		string(EmailAuthenticationMethodConfigurationState_Disabled),
		string(EmailAuthenticationMethodConfigurationState_Enabled),
	}
}

func (s *EmailAuthenticationMethodConfigurationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEmailAuthenticationMethodConfigurationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEmailAuthenticationMethodConfigurationState(input string) (*EmailAuthenticationMethodConfigurationState, error) {
	vals := map[string]EmailAuthenticationMethodConfigurationState{
		"disabled": EmailAuthenticationMethodConfigurationState_Disabled,
		"enabled":  EmailAuthenticationMethodConfigurationState_Enabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EmailAuthenticationMethodConfigurationState(input)
	return &out, nil
}
