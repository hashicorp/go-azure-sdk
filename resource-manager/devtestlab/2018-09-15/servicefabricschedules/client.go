package servicefabricschedules

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceFabricSchedulesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewServiceFabricSchedulesClientWithBaseURI(endpoint string) ServiceFabricSchedulesClient {
	return ServiceFabricSchedulesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
