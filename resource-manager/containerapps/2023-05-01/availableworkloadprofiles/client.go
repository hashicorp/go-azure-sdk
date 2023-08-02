package availableworkloadprofiles

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AvailableWorkloadProfilesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewAvailableWorkloadProfilesClientWithBaseURI(endpoint string) AvailableWorkloadProfilesClient {
	return AvailableWorkloadProfilesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
