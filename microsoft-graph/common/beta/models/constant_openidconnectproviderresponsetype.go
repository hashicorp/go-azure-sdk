package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OpenIdConnectProviderResponseType string

const (
	OpenIdConnectProviderResponseTypecode     OpenIdConnectProviderResponseType = "Code"
	OpenIdConnectProviderResponseTypeid_token OpenIdConnectProviderResponseType = "Idtoken"
	OpenIdConnectProviderResponseTypetoken    OpenIdConnectProviderResponseType = "Token"
)

func PossibleValuesForOpenIdConnectProviderResponseType() []string {
	return []string{
		string(OpenIdConnectProviderResponseTypecode),
		string(OpenIdConnectProviderResponseTypeid_token),
		string(OpenIdConnectProviderResponseTypetoken),
	}
}

func parseOpenIdConnectProviderResponseType(input string) (*OpenIdConnectProviderResponseType, error) {
	vals := map[string]OpenIdConnectProviderResponseType{
		"code":    OpenIdConnectProviderResponseTypecode,
		"idtoken": OpenIdConnectProviderResponseTypeid_token,
		"token":   OpenIdConnectProviderResponseTypetoken,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OpenIdConnectProviderResponseType(input)
	return &out, nil
}
