package subvolumes

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubVolumesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewSubVolumesClientWithBaseURI(endpoint string) SubVolumesClient {
	return SubVolumesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
