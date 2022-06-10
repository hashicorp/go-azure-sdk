package diskpoolzones

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiskPoolZonesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewDiskPoolZonesClientWithBaseURI(endpoint string) DiskPoolZonesClient {
	return DiskPoolZonesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
