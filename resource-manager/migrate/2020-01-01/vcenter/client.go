package vcenter

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VCenterClient struct {
	Client  autorest.Client
	baseUri string
}

func NewVCenterClientWithBaseURI(endpoint string) VCenterClient {
	return VCenterClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
