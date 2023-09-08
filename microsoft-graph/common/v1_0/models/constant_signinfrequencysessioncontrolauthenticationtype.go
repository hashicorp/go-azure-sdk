package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInFrequencySessionControlAuthenticationType string

const (
	SignInFrequencySessionControlAuthenticationTypeprimaryAndSecondaryAuthentication SignInFrequencySessionControlAuthenticationType = "PrimaryAndSecondaryAuthentication"
	SignInFrequencySessionControlAuthenticationTypesecondaryAuthentication           SignInFrequencySessionControlAuthenticationType = "SecondaryAuthentication"
)

func PossibleValuesForSignInFrequencySessionControlAuthenticationType() []string {
	return []string{
		string(SignInFrequencySessionControlAuthenticationTypeprimaryAndSecondaryAuthentication),
		string(SignInFrequencySessionControlAuthenticationTypesecondaryAuthentication),
	}
}

func parseSignInFrequencySessionControlAuthenticationType(input string) (*SignInFrequencySessionControlAuthenticationType, error) {
	vals := map[string]SignInFrequencySessionControlAuthenticationType{
		"primaryandsecondaryauthentication": SignInFrequencySessionControlAuthenticationTypeprimaryAndSecondaryAuthentication,
		"secondaryauthentication":           SignInFrequencySessionControlAuthenticationTypesecondaryAuthentication,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SignInFrequencySessionControlAuthenticationType(input)
	return &out, nil
}
