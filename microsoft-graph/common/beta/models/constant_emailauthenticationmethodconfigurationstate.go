package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmailAuthenticationMethodConfigurationState string

const (
	EmailAuthenticationMethodConfigurationStatedisabled EmailAuthenticationMethodConfigurationState = "Disabled"
	EmailAuthenticationMethodConfigurationStateenabled  EmailAuthenticationMethodConfigurationState = "Enabled"
)

func PossibleValuesForEmailAuthenticationMethodConfigurationState() []string {
	return []string{
		string(EmailAuthenticationMethodConfigurationStatedisabled),
		string(EmailAuthenticationMethodConfigurationStateenabled),
	}
}

func parseEmailAuthenticationMethodConfigurationState(input string) (*EmailAuthenticationMethodConfigurationState, error) {
	vals := map[string]EmailAuthenticationMethodConfigurationState{
		"disabled": EmailAuthenticationMethodConfigurationStatedisabled,
		"enabled":  EmailAuthenticationMethodConfigurationStateenabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EmailAuthenticationMethodConfigurationState(input)
	return &out, nil
}
