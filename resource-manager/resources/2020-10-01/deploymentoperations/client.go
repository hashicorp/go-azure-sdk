package deploymentoperations

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeploymentOperationsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewDeploymentOperationsClientWithBaseURI(endpoint string) DeploymentOperationsClient {
	return DeploymentOperationsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
