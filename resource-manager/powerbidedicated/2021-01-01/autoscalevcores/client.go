package autoscalevcores

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutoScaleVCoresClient struct {
	Client  autorest.Client
	baseUri string
}

func NewAutoScaleVCoresClientWithBaseURI(endpoint string) AutoScaleVCoresClient {
	return AutoScaleVCoresClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
