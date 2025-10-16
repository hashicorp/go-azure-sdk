package softdeletedcontainers

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SoftDeletedContainersClient struct {
	Client  autorest.Client
	baseUri string
}

func NewSoftDeletedContainersClientWithBaseURI(endpoint string) SoftDeletedContainersClient {
	return SoftDeletedContainersClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
