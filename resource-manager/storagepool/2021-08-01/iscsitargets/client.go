package iscsitargets

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IscsiTargetsClient struct {
	Client *resourcemanager.Client
}

func NewIscsiTargetsClientWithBaseURI(api environments.Api) (*IscsiTargetsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "iscsitargets", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IscsiTargetsClient: %+v", err)
	}

	return &IscsiTargetsClient{
		Client: client,
	}, nil
}
