package job

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobClient struct {
	Client  autorest.Client
	baseUri string
}

func NewJobClientWithBaseURI(endpoint string) JobClient {
	return JobClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
