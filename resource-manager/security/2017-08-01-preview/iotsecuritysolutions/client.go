package iotsecuritysolutions

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IotSecuritySolutionsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewIotSecuritySolutionsClientWithBaseURI(endpoint string) IotSecuritySolutionsClient {
	return IotSecuritySolutionsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
