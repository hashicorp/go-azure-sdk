package associations

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssociationsClient struct {
	Client *resourcemanager.Client
}

func NewAssociationsClientWithBaseURI(api environments.Api) (*AssociationsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "associations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AssociationsClient: %+v", err)
	}

	return &AssociationsClient{
		Client: client,
	}, nil
}
