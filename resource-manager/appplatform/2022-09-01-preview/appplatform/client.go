package appplatform

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppPlatformClient struct {
	Client  autorest.Client
	baseUri string
}

func NewAppPlatformClientWithBaseURI(endpoint string) AppPlatformClient {
	return AppPlatformClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
