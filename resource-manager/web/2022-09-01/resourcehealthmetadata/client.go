package resourcehealthmetadata

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceHealthMetadataClient struct {
	Client  autorest.Client
	baseUri string
}

func NewResourceHealthMetadataClientWithBaseURI(endpoint string) ResourceHealthMetadataClient {
	return ResourceHealthMetadataClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
