package changedetection

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChangeDetectionClient struct {
	Client  autorest.Client
	baseUri string
}

func NewChangeDetectionClientWithBaseURI(endpoint string) ChangeDetectionClient {
	return ChangeDetectionClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
