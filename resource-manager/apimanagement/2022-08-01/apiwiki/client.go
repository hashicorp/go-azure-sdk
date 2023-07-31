package apiwiki

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiWikiClient struct {
	Client *resourcemanager.Client
}

func NewApiWikiClientWithBaseURI(api environments.Api) (*ApiWikiClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "apiwiki", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ApiWikiClient: %+v", err)
	}

	return &ApiWikiClient{
		Client: client,
	}, nil
}
