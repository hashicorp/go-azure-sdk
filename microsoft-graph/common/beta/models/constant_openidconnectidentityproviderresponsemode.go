package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OpenIdConnectIdentityProviderResponseMode string

const (
	OpenIdConnectIdentityProviderResponseModeform_post OpenIdConnectIdentityProviderResponseMode = "Formpost"
	OpenIdConnectIdentityProviderResponseModequery     OpenIdConnectIdentityProviderResponseMode = "Query"
)

func PossibleValuesForOpenIdConnectIdentityProviderResponseMode() []string {
	return []string{
		string(OpenIdConnectIdentityProviderResponseModeform_post),
		string(OpenIdConnectIdentityProviderResponseModequery),
	}
}

func parseOpenIdConnectIdentityProviderResponseMode(input string) (*OpenIdConnectIdentityProviderResponseMode, error) {
	vals := map[string]OpenIdConnectIdentityProviderResponseMode{
		"formpost": OpenIdConnectIdentityProviderResponseModeform_post,
		"query":    OpenIdConnectIdentityProviderResponseModequery,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OpenIdConnectIdentityProviderResponseMode(input)
	return &out, nil
}
