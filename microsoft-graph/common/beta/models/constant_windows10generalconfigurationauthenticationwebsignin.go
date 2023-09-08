package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationAuthenticationWebSignIn string

const (
	Windows10GeneralConfigurationAuthenticationWebSignIndisabled      Windows10GeneralConfigurationAuthenticationWebSignIn = "Disabled"
	Windows10GeneralConfigurationAuthenticationWebSignInenabled       Windows10GeneralConfigurationAuthenticationWebSignIn = "Enabled"
	Windows10GeneralConfigurationAuthenticationWebSignInnotConfigured Windows10GeneralConfigurationAuthenticationWebSignIn = "NotConfigured"
)

func PossibleValuesForWindows10GeneralConfigurationAuthenticationWebSignIn() []string {
	return []string{
		string(Windows10GeneralConfigurationAuthenticationWebSignIndisabled),
		string(Windows10GeneralConfigurationAuthenticationWebSignInenabled),
		string(Windows10GeneralConfigurationAuthenticationWebSignInnotConfigured),
	}
}

func parseWindows10GeneralConfigurationAuthenticationWebSignIn(input string) (*Windows10GeneralConfigurationAuthenticationWebSignIn, error) {
	vals := map[string]Windows10GeneralConfigurationAuthenticationWebSignIn{
		"disabled":      Windows10GeneralConfigurationAuthenticationWebSignIndisabled,
		"enabled":       Windows10GeneralConfigurationAuthenticationWebSignInenabled,
		"notconfigured": Windows10GeneralConfigurationAuthenticationWebSignInnotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationAuthenticationWebSignIn(input)
	return &out, nil
}
