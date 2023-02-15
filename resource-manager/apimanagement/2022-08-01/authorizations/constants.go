package authorizations

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthorizationType string

const (
	AuthorizationTypeOAuthTwo AuthorizationType = "OAuth2"
)

func PossibleValuesForAuthorizationType() []string {
	return []string{
		string(AuthorizationTypeOAuthTwo),
	}
}

func parseAuthorizationType(input string) (*AuthorizationType, error) {
	vals := map[string]AuthorizationType{
		"oauth2": AuthorizationTypeOAuthTwo,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthorizationType(input)
	return &out, nil
}

type OAuth2GrantType string

const (
	OAuth2GrantTypeAuthorizationCode OAuth2GrantType = "AuthorizationCode"
	OAuth2GrantTypeClientCredentials OAuth2GrantType = "ClientCredentials"
)

func PossibleValuesForOAuth2GrantType() []string {
	return []string{
		string(OAuth2GrantTypeAuthorizationCode),
		string(OAuth2GrantTypeClientCredentials),
	}
}

func parseOAuth2GrantType(input string) (*OAuth2GrantType, error) {
	vals := map[string]OAuth2GrantType{
		"authorizationcode": OAuth2GrantTypeAuthorizationCode,
		"clientcredentials": OAuth2GrantTypeClientCredentials,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OAuth2GrantType(input)
	return &out, nil
}
