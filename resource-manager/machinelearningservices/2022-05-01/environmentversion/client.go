package environmentversion

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnvironmentVersionClient struct {
	Client  autorest.Client
	baseUri string
}

func NewEnvironmentVersionClientWithBaseURI(endpoint string) EnvironmentVersionClient {
	return EnvironmentVersionClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
