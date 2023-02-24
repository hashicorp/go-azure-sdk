package pipelineruns

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PipelineRunsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewPipelineRunsClientWithBaseURI(endpoint string) PipelineRunsClient {
	return PipelineRunsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
