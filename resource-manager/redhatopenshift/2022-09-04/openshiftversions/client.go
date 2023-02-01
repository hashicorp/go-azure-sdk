package openshiftversions

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OpenShiftVersionsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewOpenShiftVersionsClientWithBaseURI(endpoint string) OpenShiftVersionsClient {
	return OpenShiftVersionsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
