package advisors

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdvisorsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewAdvisorsClientWithBaseURI(endpoint string) AdvisorsClient {
	return AdvisorsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
