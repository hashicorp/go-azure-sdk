package containerappsauthconfigs

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContainerAppsAuthConfigsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewContainerAppsAuthConfigsClientWithBaseURI(endpoint string) ContainerAppsAuthConfigsClient {
	return ContainerAppsAuthConfigsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
