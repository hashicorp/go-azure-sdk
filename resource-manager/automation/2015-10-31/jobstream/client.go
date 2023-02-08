package jobstream

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobStreamClient struct {
	Client  autorest.Client
	baseUri string
}

func NewJobStreamClientWithBaseURI(endpoint string) JobStreamClient {
	return JobStreamClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
