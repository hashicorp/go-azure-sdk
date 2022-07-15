package volumesrelocation

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VolumesRelocationClient struct {
	Client  autorest.Client
	baseUri string
}

func NewVolumesRelocationClientWithBaseURI(endpoint string) VolumesRelocationClient {
	return VolumesRelocationClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
