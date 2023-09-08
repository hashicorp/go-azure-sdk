package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmailAuthenticationMethodConfigurationAllowExternalIdToUseEmailOtp string

const (
	EmailAuthenticationMethodConfigurationAllowExternalIdToUseEmailOtpdefault  EmailAuthenticationMethodConfigurationAllowExternalIdToUseEmailOtp = "Default"
	EmailAuthenticationMethodConfigurationAllowExternalIdToUseEmailOtpdisabled EmailAuthenticationMethodConfigurationAllowExternalIdToUseEmailOtp = "Disabled"
	EmailAuthenticationMethodConfigurationAllowExternalIdToUseEmailOtpenabled  EmailAuthenticationMethodConfigurationAllowExternalIdToUseEmailOtp = "Enabled"
)

func PossibleValuesForEmailAuthenticationMethodConfigurationAllowExternalIdToUseEmailOtp() []string {
	return []string{
		string(EmailAuthenticationMethodConfigurationAllowExternalIdToUseEmailOtpdefault),
		string(EmailAuthenticationMethodConfigurationAllowExternalIdToUseEmailOtpdisabled),
		string(EmailAuthenticationMethodConfigurationAllowExternalIdToUseEmailOtpenabled),
	}
}

func parseEmailAuthenticationMethodConfigurationAllowExternalIdToUseEmailOtp(input string) (*EmailAuthenticationMethodConfigurationAllowExternalIdToUseEmailOtp, error) {
	vals := map[string]EmailAuthenticationMethodConfigurationAllowExternalIdToUseEmailOtp{
		"default":  EmailAuthenticationMethodConfigurationAllowExternalIdToUseEmailOtpdefault,
		"disabled": EmailAuthenticationMethodConfigurationAllowExternalIdToUseEmailOtpdisabled,
		"enabled":  EmailAuthenticationMethodConfigurationAllowExternalIdToUseEmailOtpenabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EmailAuthenticationMethodConfigurationAllowExternalIdToUseEmailOtp(input)
	return &out, nil
}
