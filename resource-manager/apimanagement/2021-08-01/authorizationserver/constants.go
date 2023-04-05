package authorizationserver

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
