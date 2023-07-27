package deletedservice

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeletedServiceClient struct {
	Client  autorest.Client
	baseUri string
}

func NewDeletedServiceClientWithBaseURI(endpoint string) DeletedServiceClient {
	return DeletedServiceClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
