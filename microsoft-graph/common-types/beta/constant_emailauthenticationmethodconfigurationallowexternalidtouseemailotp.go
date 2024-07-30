package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmailAuthenticationMethodConfigurationAllowExternalIdToUseEmailOtp string

const (
	EmailAuthenticationMethodConfigurationAllowExternalIdToUseEmailOtp_Default  EmailAuthenticationMethodConfigurationAllowExternalIdToUseEmailOtp = "default"
	EmailAuthenticationMethodConfigurationAllowExternalIdToUseEmailOtp_Disabled EmailAuthenticationMethodConfigurationAllowExternalIdToUseEmailOtp = "disabled"
	EmailAuthenticationMethodConfigurationAllowExternalIdToUseEmailOtp_Enabled  EmailAuthenticationMethodConfigurationAllowExternalIdToUseEmailOtp = "enabled"
)

func PossibleValuesForEmailAuthenticationMethodConfigurationAllowExternalIdToUseEmailOtp() []string {
	return []string{
		string(EmailAuthenticationMethodConfigurationAllowExternalIdToUseEmailOtp_Default),
		string(EmailAuthenticationMethodConfigurationAllowExternalIdToUseEmailOtp_Disabled),
		string(EmailAuthenticationMethodConfigurationAllowExternalIdToUseEmailOtp_Enabled),
	}
}

func (s *EmailAuthenticationMethodConfigurationAllowExternalIdToUseEmailOtp) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEmailAuthenticationMethodConfigurationAllowExternalIdToUseEmailOtp(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEmailAuthenticationMethodConfigurationAllowExternalIdToUseEmailOtp(input string) (*EmailAuthenticationMethodConfigurationAllowExternalIdToUseEmailOtp, error) {
	vals := map[string]EmailAuthenticationMethodConfigurationAllowExternalIdToUseEmailOtp{
		"default":  EmailAuthenticationMethodConfigurationAllowExternalIdToUseEmailOtp_Default,
		"disabled": EmailAuthenticationMethodConfigurationAllowExternalIdToUseEmailOtp_Disabled,
		"enabled":  EmailAuthenticationMethodConfigurationAllowExternalIdToUseEmailOtp_Enabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EmailAuthenticationMethodConfigurationAllowExternalIdToUseEmailOtp(input)
	return &out, nil
}
