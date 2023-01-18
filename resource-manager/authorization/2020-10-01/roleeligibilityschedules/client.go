package roleeligibilityschedules

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleEligibilitySchedulesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewRoleEligibilitySchedulesClientWithBaseURI(endpoint string) RoleEligibilitySchedulesClient {
	return RoleEligibilitySchedulesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
