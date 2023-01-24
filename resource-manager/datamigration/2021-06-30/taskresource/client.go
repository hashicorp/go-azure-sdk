package taskresource

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TaskResourceClient struct {
	Client  autorest.Client
	baseUri string
}

func NewTaskResourceClientWithBaseURI(endpoint string) TaskResourceClient {
	return TaskResourceClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
