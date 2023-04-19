package consumerinvitation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConsumerInvitationClient struct {
	Client *resourcemanager.Client
}

func NewConsumerInvitationClientWithBaseURI(api environments.Api) (*ConsumerInvitationClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "consumerinvitation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ConsumerInvitationClient: %+v", err)
	}

	return &ConsumerInvitationClient{
		Client: client,
	}, nil
}
