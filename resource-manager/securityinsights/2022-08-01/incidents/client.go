package incidents

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IncidentsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewIncidentsClientWithBaseURI(endpoint string) IncidentsClient {
	return IncidentsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
