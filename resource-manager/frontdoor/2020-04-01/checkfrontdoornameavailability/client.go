package checkfrontdoornameavailability

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CheckFrontDoorNameAvailabilityClient struct {
	Client  autorest.Client
	baseUri string
}

func NewCheckFrontDoorNameAvailabilityClientWithBaseURI(endpoint string) CheckFrontDoorNameAvailabilityClient {
	return CheckFrontDoorNameAvailabilityClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
