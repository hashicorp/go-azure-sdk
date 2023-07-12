package groupidlistforldapuser

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupIdListForLDAPUserClient struct {
	Client *resourcemanager.Client
}

func NewGroupIdListForLDAPUserClientWithBaseURI(api environments.Api) (*GroupIdListForLDAPUserClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "groupidlistforldapuser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupIdListForLDAPUserClient: %+v", err)
	}

	return &GroupIdListForLDAPUserClient{
		Client: client,
	}, nil
}
