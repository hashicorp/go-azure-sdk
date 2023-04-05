package apioperationsbytag

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

type ProductState string

const (
	ProductStateNotPublished ProductState = "notPublished"
	ProductStatePublished    ProductState = "published"
)

func PossibleValuesForProductState() []string {
	return []string{
		string(ProductStateNotPublished),
		string(ProductStatePublished),
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
