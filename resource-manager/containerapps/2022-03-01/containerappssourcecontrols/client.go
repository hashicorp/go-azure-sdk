package containerappssourcecontrols

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContainerAppsSourceControlsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewContainerAppsSourceControlsClientWithBaseURI(endpoint string) ContainerAppsSourceControlsClient {
	return ContainerAppsSourceControlsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
