package featuresupport

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FeatureSupportClient struct {
	Client  autorest.Client
	baseUri string
}

func NewFeatureSupportClientWithBaseURI(endpoint string) FeatureSupportClient {
	return FeatureSupportClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
