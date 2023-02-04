package recoverableservers

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecoverableServersClient struct {
	Client *resourcemanager.Client
}

func NewRecoverableServersClientWithBaseURI(api environments.Api) (*RecoverableServersClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "recoverableservers", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RecoverableServersClient: %+v", err)
	}

	return &RecoverableServersClient{
		Client: client,
	}, nil
}
