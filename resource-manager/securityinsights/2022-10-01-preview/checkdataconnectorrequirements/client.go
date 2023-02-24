package checkdataconnectorrequirements

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CheckDataConnectorRequirementsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewCheckDataConnectorRequirementsClientWithBaseURI(endpoint string) CheckDataConnectorRequirementsClient {
	return CheckDataConnectorRequirementsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
