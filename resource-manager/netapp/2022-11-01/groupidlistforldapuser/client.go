package groupidlistforldapuser

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupIdListForLDAPUserClient struct {
	Client  autorest.Client
	baseUri string
}

func NewGroupIdListForLDAPUserClientWithBaseURI(endpoint string) GroupIdListForLDAPUserClient {
	return GroupIdListForLDAPUserClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
