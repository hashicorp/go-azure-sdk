package gatewayapi

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiType string

const (
	ApiTypeGraphql   ApiType = "graphql"
	ApiTypeHttp      ApiType = "http"
	ApiTypeSoap      ApiType = "soap"
	ApiTypeWebsocket ApiType = "websocket"
)

func PossibleValuesForApiType() []string {
	return []string{
		string(ApiTypeGraphql),
		string(ApiTypeHttp),
		string(ApiTypeSoap),
		string(ApiTypeWebsocket),
	}
}

func parseApiType(input string) (*ApiType, error) {
	vals := map[string]ApiType{
		"graphql":   ApiTypeGraphql,
		"http":      ApiTypeHttp,
		"soap":      ApiTypeSoap,
		"websocket": ApiTypeWebsocket,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ApiType(input)
	return &out, nil
}

type BearerTokenSendingMethods string

const (
	BearerTokenSendingMethodsAuthorizationHeader BearerTokenSendingMethods = "authorizationHeader"
	BearerTokenSendingMethodsQuery               BearerTokenSendingMethods = "query"
)

func PossibleValuesForBearerTokenSendingMethods() []string {
	return []string{
		string(BearerTokenSendingMethodsAuthorizationHeader),
		string(BearerTokenSendingMethodsQuery),
	}
}

func parseBearerTokenSendingMethods(input string) (*BearerTokenSendingMethods, error) {
	vals := map[string]BearerTokenSendingMethods{
		"authorizationheader": BearerTokenSendingMethodsAuthorizationHeader,
		"query":               BearerTokenSendingMethodsQuery,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BearerTokenSendingMethods(input)
	return &out, nil
}

type Protocol string

const (
	ProtocolHttp  Protocol = "http"
	ProtocolHttps Protocol = "https"
	ProtocolWs    Protocol = "ws"
	ProtocolWss   Protocol = "wss"
)

func PossibleValuesForProtocol() []string {
	return []string{
		string(ProtocolHttp),
		string(ProtocolHttps),
		string(ProtocolWs),
		string(ProtocolWss),
	}
}

func parseProtocol(input string) (*Protocol, error) {
	vals := map[string]Protocol{
		"http":  ProtocolHttp,
		"https": ProtocolHttps,
		"ws":    ProtocolWs,
		"wss":   ProtocolWss,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Protocol(input)
	return &out, nil
}

type ProvisioningState string

const (
	ProvisioningStateCreated ProvisioningState = "created"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCreated),
	}
}

func parseProvisioningState(input string) (*ProvisioningState, error) {
	vals := map[string]ProvisioningState{
		"created": ProvisioningStateCreated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningState(input)
	return &out, nil
}

type VersioningScheme string

const (
	VersioningSchemeHeader  VersioningScheme = "Header"
	VersioningSchemeQuery   VersioningScheme = "Query"
	VersioningSchemeSegment VersioningScheme = "Segment"
)

func PossibleValuesForVersioningScheme() []string {
	return []string{
		string(VersioningSchemeHeader),
		string(VersioningSchemeQuery),
		string(VersioningSchemeSegment),
	}
}

func parseVersioningScheme(input string) (*VersioningScheme, error) {
	vals := map[string]VersioningScheme{
		"header":  VersioningSchemeHeader,
		"query":   VersioningSchemeQuery,
		"segment": VersioningSchemeSegment,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VersioningScheme(input)
	return &out, nil
}
