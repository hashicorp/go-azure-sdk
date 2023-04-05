package api

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiType string

const (
	ApiTypeGraphql   ApiType = "graphql"
	ApiTypeHTTP      ApiType = "http"
	ApiTypeSoap      ApiType = "soap"
	ApiTypeWebsocket ApiType = "websocket"
)

func PossibleValuesForApiType() []string {
	return []string{
		string(ApiTypeGraphql),
		string(ApiTypeHTTP),
		string(ApiTypeSoap),
		string(ApiTypeWebsocket),
	}
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

type ContentFormat string

const (
	ContentFormatGraphqlNegativelink             ContentFormat = "graphql-link"
	ContentFormatOpenapi                         ContentFormat = "openapi"
	ContentFormatOpenapiNegativelink             ContentFormat = "openapi-link"
	ContentFormatOpenapiPositivejson             ContentFormat = "openapi+json"
	ContentFormatOpenapiPositivejsonNegativelink ContentFormat = "openapi+json-link"
	ContentFormatSwaggerNegativejson             ContentFormat = "swagger-json"
	ContentFormatSwaggerNegativelinkNegativejson ContentFormat = "swagger-link-json"
	ContentFormatWadlNegativelinkNegativejson    ContentFormat = "wadl-link-json"
	ContentFormatWadlNegativexml                 ContentFormat = "wadl-xml"
	ContentFormatWsdl                            ContentFormat = "wsdl"
	ContentFormatWsdlNegativelink                ContentFormat = "wsdl-link"
)

func PossibleValuesForContentFormat() []string {
	return []string{
		string(ContentFormatGraphqlNegativelink),
		string(ContentFormatOpenapi),
		string(ContentFormatOpenapiNegativelink),
		string(ContentFormatOpenapiPositivejson),
		string(ContentFormatOpenapiPositivejsonNegativelink),
		string(ContentFormatSwaggerNegativejson),
		string(ContentFormatSwaggerNegativelinkNegativejson),
		string(ContentFormatWadlNegativelinkNegativejson),
		string(ContentFormatWadlNegativexml),
		string(ContentFormatWsdl),
		string(ContentFormatWsdlNegativelink),
	}
}

type Protocol string

const (
	ProtocolHTTP  Protocol = "http"
	ProtocolHTTPS Protocol = "https"
	ProtocolWs    Protocol = "ws"
	ProtocolWss   Protocol = "wss"
)

func PossibleValuesForProtocol() []string {
	return []string{
		string(ProtocolHTTP),
		string(ProtocolHTTPS),
		string(ProtocolWs),
		string(ProtocolWss),
	}
}

type SoapApiType string

const (
	SoapApiTypeGraphql   SoapApiType = "graphql"
	SoapApiTypeHTTP      SoapApiType = "http"
	SoapApiTypeSoap      SoapApiType = "soap"
	SoapApiTypeWebsocket SoapApiType = "websocket"
)

func PossibleValuesForSoapApiType() []string {
	return []string{
		string(SoapApiTypeGraphql),
		string(SoapApiTypeHTTP),
		string(SoapApiTypeSoap),
		string(SoapApiTypeWebsocket),
	}
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
