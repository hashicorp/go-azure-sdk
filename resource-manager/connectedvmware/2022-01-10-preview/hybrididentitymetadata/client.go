package hybrididentitymetadata

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HybridIdentityMetadataClient struct {
	Client  autorest.Client
	baseUri string
}

func NewHybridIdentityMetadataClientWithBaseURI(endpoint string) HybridIdentityMetadataClient {
	return HybridIdentityMetadataClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
