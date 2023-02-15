package authorizationserver

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthorizationMethod string

const (
	AuthorizationMethodDELETE  AuthorizationMethod = "DELETE"
	AuthorizationMethodGET     AuthorizationMethod = "GET"
	AuthorizationMethodHEAD    AuthorizationMethod = "HEAD"
	AuthorizationMethodOPTIONS AuthorizationMethod = "OPTIONS"
	AuthorizationMethodPATCH   AuthorizationMethod = "PATCH"
	AuthorizationMethodPOST    AuthorizationMethod = "POST"
	AuthorizationMethodPUT     AuthorizationMethod = "PUT"
	AuthorizationMethodTRACE   AuthorizationMethod = "TRACE"
)

func PossibleValuesForAuthorizationMethod() []string {
	return []string{
		string(AuthorizationMethodDELETE),
		string(AuthorizationMethodGET),
		string(AuthorizationMethodHEAD),
		string(AuthorizationMethodOPTIONS),
		string(AuthorizationMethodPATCH),
		string(AuthorizationMethodPOST),
		string(AuthorizationMethodPUT),
		string(AuthorizationMethodTRACE),
	}
}

func parseAuthorizationMethod(input string) (*AuthorizationMethod, error) {
	vals := map[string]AuthorizationMethod{
		"delete":  AuthorizationMethodDELETE,
		"get":     AuthorizationMethodGET,
		"head":    AuthorizationMethodHEAD,
		"options": AuthorizationMethodOPTIONS,
		"patch":   AuthorizationMethodPATCH,
		"post":    AuthorizationMethodPOST,
		"put":     AuthorizationMethodPUT,
		"trace":   AuthorizationMethodTRACE,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthorizationMethod(input)
	return &out, nil
}

type BearerTokenSendingMethod string

const (
	BearerTokenSendingMethodAuthorizationHeader BearerTokenSendingMethod = "authorizationHeader"
	BearerTokenSendingMethodQuery               BearerTokenSendingMethod = "query"
)

func PossibleValuesForBearerTokenSendingMethod() []string {
	return []string{
		string(BearerTokenSendingMethodAuthorizationHeader),
		string(BearerTokenSendingMethodQuery),
	}
}

func parseBearerTokenSendingMethod(input string) (*BearerTokenSendingMethod, error) {
	vals := map[string]BearerTokenSendingMethod{
		"authorizationheader": BearerTokenSendingMethodAuthorizationHeader,
		"query":               BearerTokenSendingMethodQuery,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BearerTokenSendingMethod(input)
	return &out, nil
}

type ClientAuthenticationMethod string

const (
	ClientAuthenticationMethodBasic ClientAuthenticationMethod = "Basic"
	ClientAuthenticationMethodBody  ClientAuthenticationMethod = "Body"
)

func PossibleValuesForClientAuthenticationMethod() []string {
	return []string{
		string(ClientAuthenticationMethodBasic),
		string(ClientAuthenticationMethodBody),
	}
}

func parseClientAuthenticationMethod(input string) (*ClientAuthenticationMethod, error) {
	vals := map[string]ClientAuthenticationMethod{
		"basic": ClientAuthenticationMethodBasic,
		"body":  ClientAuthenticationMethodBody,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ClientAuthenticationMethod(input)
	return &out, nil
}

type GrantType string

const (
	GrantTypeAuthorizationCode     GrantType = "authorizationCode"
	GrantTypeClientCredentials     GrantType = "clientCredentials"
	GrantTypeImplicit              GrantType = "implicit"
	GrantTypeResourceOwnerPassword GrantType = "resourceOwnerPassword"
)

func PossibleValuesForGrantType() []string {
	return []string{
		string(GrantTypeAuthorizationCode),
		string(GrantTypeClientCredentials),
		string(GrantTypeImplicit),
		string(GrantTypeResourceOwnerPassword),
	}
}

func parseGrantType(input string) (*GrantType, error) {
	vals := map[string]GrantType{
		"authorizationcode":     GrantTypeAuthorizationCode,
		"clientcredentials":     GrantTypeClientCredentials,
		"implicit":              GrantTypeImplicit,
		"resourceownerpassword": GrantTypeResourceOwnerPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GrantType(input)
	return &out, nil
}
