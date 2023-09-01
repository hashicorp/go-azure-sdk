package mhsmlistregions

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MHSMListRegionsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewMHSMListRegionsClientWithBaseURI(endpoint string) MHSMListRegionsClient {
	return MHSMListRegionsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
