package onboardtod4api

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnboardToD4APIClient struct {
	Client  autorest.Client
	baseUri string
}

func NewOnboardToD4APIClientWithBaseURI(endpoint string) OnboardToD4APIClient {
	return OnboardToD4APIClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
