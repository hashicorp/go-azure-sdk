package authorizationrulesdisasterrecoveryconfigs

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthorizationRulesDisasterRecoveryConfigsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewAuthorizationRulesDisasterRecoveryConfigsClientWithBaseURI(endpoint string) AuthorizationRulesDisasterRecoveryConfigsClient {
	return AuthorizationRulesDisasterRecoveryConfigsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
