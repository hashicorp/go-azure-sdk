package appservicecertificateorders

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppServiceCertificateOrdersClient struct {
	Client  autorest.Client
	baseUri string
}

func NewAppServiceCertificateOrdersClientWithBaseURI(endpoint string) AppServiceCertificateOrdersClient {
	return AppServiceCertificateOrdersClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
