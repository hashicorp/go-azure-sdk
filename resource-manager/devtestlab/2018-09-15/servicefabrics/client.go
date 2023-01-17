package servicefabrics

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceFabricsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewServiceFabricsClientWithBaseURI(endpoint string) ServiceFabricsClient {
	return ServiceFabricsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
