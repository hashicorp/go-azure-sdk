package containerappsrevisionreplicas

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContainerAppsRevisionReplicasClient struct {
	Client  autorest.Client
	baseUri string
}

func NewContainerAppsRevisionReplicasClientWithBaseURI(endpoint string) ContainerAppsRevisionReplicasClient {
	return ContainerAppsRevisionReplicasClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
