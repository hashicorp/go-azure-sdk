package advisorscore

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdvisorScoreClient struct {
	Client  autorest.Client
	baseUri string
}

func NewAdvisorScoreClientWithBaseURI(endpoint string) AdvisorScoreClient {
	return AdvisorScoreClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
