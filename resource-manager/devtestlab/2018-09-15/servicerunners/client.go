package servicerunners

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceRunnersClient struct {
	Client  autorest.Client
	baseUri string
}

func NewServiceRunnersClientWithBaseURI(endpoint string) ServiceRunnersClient {
	return ServiceRunnersClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
