package privatelinkassociation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateLinkAssociationClient struct {
	Client *resourcemanager.Client
}

func NewPrivateLinkAssociationClientWithBaseURI(api environments.Api) (*PrivateLinkAssociationClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "privatelinkassociation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivateLinkAssociationClient: %+v", err)
	}

	return &PrivateLinkAssociationClient{
		Client: client,
	}, nil
}
