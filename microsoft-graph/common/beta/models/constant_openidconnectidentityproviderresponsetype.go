package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OpenIdConnectIdentityProviderResponseType string

const (
	OpenIdConnectIdentityProviderResponseTypecode     OpenIdConnectIdentityProviderResponseType = "Code"
	OpenIdConnectIdentityProviderResponseTypeid_token OpenIdConnectIdentityProviderResponseType = "Idtoken"
	OpenIdConnectIdentityProviderResponseTypetoken    OpenIdConnectIdentityProviderResponseType = "Token"
)

func PossibleValuesForOpenIdConnectIdentityProviderResponseType() []string {
	return []string{
		string(OpenIdConnectIdentityProviderResponseTypecode),
		string(OpenIdConnectIdentityProviderResponseTypeid_token),
		string(OpenIdConnectIdentityProviderResponseTypetoken),
	}
}

func parseOpenIdConnectIdentityProviderResponseType(input string) (*OpenIdConnectIdentityProviderResponseType, error) {
	vals := map[string]OpenIdConnectIdentityProviderResponseType{
		"code":    OpenIdConnectIdentityProviderResponseTypecode,
		"idtoken": OpenIdConnectIdentityProviderResponseTypeid_token,
		"token":   OpenIdConnectIdentityProviderResponseTypetoken,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OpenIdConnectIdentityProviderResponseType(input)
	return &out, nil
}
