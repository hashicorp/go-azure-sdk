package testjob

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TestJobClient struct {
	Client  autorest.Client
	baseUri string
}

func NewTestJobClientWithBaseURI(endpoint string) TestJobClient {
	return TestJobClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
