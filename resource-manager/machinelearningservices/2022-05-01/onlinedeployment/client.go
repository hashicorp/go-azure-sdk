package onlinedeployment

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnlineDeploymentClient struct {
	Client  autorest.Client
	baseUri string
}

func NewOnlineDeploymentClientWithBaseURI(endpoint string) OnlineDeploymentClient {
	return OnlineDeploymentClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
