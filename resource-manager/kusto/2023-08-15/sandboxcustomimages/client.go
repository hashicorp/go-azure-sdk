package sandboxcustomimages

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SandboxCustomImagesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewSandboxCustomImagesClientWithBaseURI(endpoint string) SandboxCustomImagesClient {
	return SandboxCustomImagesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
