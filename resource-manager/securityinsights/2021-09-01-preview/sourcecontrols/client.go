package sourcecontrols

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SourceControlsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewSourceControlsClientWithBaseURI(endpoint string) SourceControlsClient {
	return SourceControlsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
